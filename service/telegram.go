package service

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"github.com/peterbourgon/diskv"
)

func InitBot(config Config) *tgbotapi.BotAPI {
	token := config.Get("telegram-token").(string)
	log.Printf("%v\n", token)

	var err error
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	return bot
}

func ListenAndReactInUserMessages(bot *tgbotapi.BotAPI, d *diskv.Diskv) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		chat := Chat{update.Message.Chat.ID, update.Message.From.UserName}
		chat.Save(d)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
