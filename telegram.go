package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
	"strconv"
)

type Telegram struct {
	*tgbotapi.BotAPI
	sessions *GameSessions
}

func (t *Telegram) SendAnswer(id int64, answer string) error {
	msg := tgbotapi.NewMessage(id, answer)
	_, err := t.Send(msg)
	return err
}


func (t *Telegram) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := t.GetUpdatesChan(u)
	if err != nil {
		GetConfig().LogError.Fatalf("[Get updates chan] %v", err)
	}

	for update := range updates {
		t.handlingMessage(update)
	}
}


func (t *Telegram) handlingMessage(update tgbotapi.Update) {

	if update.Message == nil {
		return
	}

	GetConfig().LogInfo.Printf("[New message] [%s] %s", getUserInfo(update), update.Message.Text)

	if update.Message.Command() == "new" {
		game := t.sessions.NewGame(update.Message.Chat.ID)
		GetConfig().LogInfo.Printf("[New game] [%s] %s", getUserInfo(update), game)
		answer := "Отгадай число"
		err := t.SendAnswer(update.Message.Chat.ID, answer)
		if err != nil {
			GetConfig().LogError.Printf("[ERROR SEND MESSAGE] msg: %s; err: %s", answer, err)
		}
		return
	}

	if update.Message.Command() != "" || len(update.Message.Text) != 4 {
		answer := "Напишите четыре цифры, а не буквы, без слеша в начале. К примеру '0123'"
		err := t.SendAnswer(update.Message.Chat.ID, answer)
		if err != nil {
			GetConfig().LogError.Printf("[ERROR SEND MESSAGE] msg: %s; err: %s", answer, err)
		}
		return
	}

	if _, err := strconv.Atoi(update.Message.Text); err != nil {
		answer := "Напишите четыре цифры, а не буквы, без слеша в начале. К примеру '0123'"
		err = t.SendAnswer(update.Message.Chat.ID, answer)
		if err != nil {
			GetConfig().LogError.Printf("[ERROR SEND MESSAGE] msg: %s; err: %s", answer, err)
		}
		return
	}

	if game, ok := t.sessions.GetGame(update.Message.Chat.ID); ok {
		if answer, done := t.sessions.Play(update.Message.Chat.ID, update.Message.Text); done {
			str := fmt.Sprint(answer, " Число найдено за ", game.step)
			err := t.SendAnswer(update.Message.Chat.ID, str)
			if err != nil {
				GetConfig().LogError.Printf("[ERROR SEND MESSAGE] msg: %s; err: %s", str, err)
			}
		} else {
			err := t.SendAnswer(update.Message.Chat.ID, answer)
			if err != nil {
				GetConfig().LogError.Printf("[ERROR SEND MESSAGE] msg: %s; err: %s", answer, err)
			}
		}
		return
	}

	answer := "Упс, ты еще не начал играть, напиши /new для начала игры"
	err := t.SendAnswer(update.Message.Chat.ID, "Упс, ты еще не начал играть, напиши /new для начала игры")
	if err != nil {
		GetConfig().LogError.Printf("[ERROR SEND MESSAGE] msg: %s; err: %s", answer, err)
	}
}

func (t *Telegram) connectTelegramBot() error {
	var err error
	t.BotAPI, err = tgbotapi.NewBotAPI(GetConfig().Token)
	if err != nil {
		return err
	}
	return nil
}

func NewTelegramBot(sessions *GameSessions) *Telegram {
	t := new(Telegram)
	err := t.connectTelegramBot()
	if err != nil {
		GetConfig().LogError.Fatalf("[Connect telegram] %s", err)
	}
	GetConfig().LogInfo.Printf("[Authorized on account] %s", t.Self.UserName)
	t.sessions = sessions
	return t
}

func getUserInfo(update tgbotapi.Update) string {
	return update.Message.Chat.UserName + ":" + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName
}
