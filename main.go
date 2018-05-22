package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"fmt"
	"strconv"
)

type Bot struct {
	 *tgbotapi.BotAPI
}



func (b *Bot) New() {
	b.BotAPI, _ = tgbotapi.NewBotAPI(GetConfig().Token)
}

func (b *Bot) SendAnswer(id int64, answer string) error {
	msg := tgbotapi.NewMessage(id, answer)
	_, err := b.Send(msg)
	return err
}

func getUserInfo(update tgbotapi.Update) string {
	return update.Message.Chat.UserName+":"+update.Message.Chat.FirstName+" "+update.Message.Chat.LastName
}

func main() {
	bot := new(Bot)
	bot.New()
//	GetConfig().LogInfo.Printf("[Authorized on account] %s", bot.Self.UserName)
	var game IDPlayers
	game = make(map[int64]*Game)
	game.NewGame(10)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		GetConfig().LogError.Fatalf("[Get updates chan] %v", err)
	}

	for update := range updates {

		if update.Message == nil {
			continue
		}

		GetConfig().LogInfo.Printf("[New message] [%s] %s", getUserInfo(update) , update.Message.Text)

		if update.Message.Command() == "new" {
			game.NewGame(update.Message.Chat.ID)
			GetConfig().LogInfo.Printf("[New game] [%s] %s", getUserInfo(update), *game[update.Message.Chat.ID])
			bot.SendAnswer(update.Message.Chat.ID, "Отгадай число")
			continue
		}

		if update.Message.Command() != "" || len(update.Message.Text) != 4 {
			err = bot.SendAnswer(update.Message.Chat.ID, "Напишите четыре цифры, без слеша в начале. К примеру '0123'")
			if err != nil {
				log.Println(err)
			}
			continue
		}

		if _, err := strconv.Atoi(update.Message.Text); err != nil {
			err = bot.SendAnswer(update.Message.Chat.ID, "Напишите четыре цифры, а не буквы, без слеша в начале. К примеру '0123'")
			if err != nil {
				log.Println(err)
			}
			continue
		}

		if progres, ok := game[update.Message.Chat.ID]; ok {
			fmt.Println(progres)
			if answer, done := game.Play(update.Message.Chat.ID, update.Message.Text); done {
				fmt.Println(done)
				str := fmt.Sprint(answer, " Число найдено за ", progres.step)
				fmt.Println(str)
				err = bot.SendAnswer(update.Message.Chat.ID, str)
				if err != nil {
					log.Println(err)
				}
			} else {
				fmt.Println(answer)
				err = bot.SendAnswer(update.Message.Chat.ID, answer)
				if err != nil {
					log.Println(err)
				}
			}
			continue
		}

		err = bot.SendAnswer(update.Message.Chat.ID, "Упс, ты еще не начал играть, напиши /new для начала игры")
		if err != nil {
			log.Println(err)
		}

	}
}
