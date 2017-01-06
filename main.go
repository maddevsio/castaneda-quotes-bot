package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	config := NewConfig("config", "yml")
	token := config.Get("telegram-token").(string)
	log.Printf("%v\n", token)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	chatID := int64(7110815)
	messageText := "Bot initialized message"

	msg := tgbotapi.NewMessage(chatID, messageText)
	bot.Send(msg)

	//updates, err := bot.GetUpdatesChan(u)
	//for update := range updates {
	//	if update.Message == nil {
	//		continue
	//	}
	//
	//	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//
	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//	msg.ReplyToMessageID = update.Message.MessageID
	//
	//	bot.Send(msg)
	//}
}