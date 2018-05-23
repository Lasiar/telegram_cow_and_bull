package main

func main() {
	sessions := NewGameSessions()
	tgBotGame := NewTelegramBot(sessions)
	tgBotGame.Run()
}
