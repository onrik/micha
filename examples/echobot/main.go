package main

import (
	"github.com/onrik/micha"
	"log"
)

func main() {
	bot, err := micha.NewBot("<token>")
	if err != nil {
		log.Fatal(err)
	}

	go bot.Start()

	for update := range bot.Updates() {
		if update.Message != nil {
			bot.SendMessage(update.Message.Chat.Id, update.Message.Text, nil)
		}
	}
}
