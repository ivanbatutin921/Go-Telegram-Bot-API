package main

import (
	"log"
	//"os"
	//"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := ("6865656772:AAFBTWReSQgSEVBGUaTGultJ6xsoXfzeKVc")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Printf("Error starting bot: %s", err)
		return // don't panic, gracefully handle the error
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	targetUsername := ("@morteit")
	if targetUsername == "" {
		log.Fatal("TELEGRAM_TARGET_USERNAME environment variable not set")
	}

	var targetUserID int64

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we receive a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// Finding userID by username (assuming the bot has seen a message from the user)
			if update.Message.From.UserName == targetUsername {
				targetUserID = update.Message.From.ID
			}

			// Отправить сообщение указанному пользователю по имени пользователя
			if targetUserID != 0 {
				msg := tgbotapi.NewMessage(targetUserID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID

				if _, err := bot.Send(msg); err != nil {
					log.Printf("Error sending message: %s", err)
					continue
				}
			} else {
				log.Printf("Target user with username %s has not been seen by the bot yet.", targetUsername)
			}
		}
	}
}