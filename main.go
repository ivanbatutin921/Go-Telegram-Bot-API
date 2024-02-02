package main

import (
	"log"
	//"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6865656772:AAFBTWReSQgSEVBGUaTGultJ6xsoXfzeKVc") // use an environment variable for the token
	if err != nil {
		log.Printf("Ошибка при запуске обоих : %s", err)
		return // don't panic, gracefully handle the error
	}

	bot.Debug = true

	log.Printf("Авторизован на учетной записи %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // Если мы получим сообщение
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				log.Printf("Сообщение об ошибке отправки: %s", err)
				continue 
			}
		}
	}
}
