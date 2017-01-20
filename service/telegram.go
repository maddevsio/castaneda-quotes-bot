package service

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"github.com/peterbourgon/diskv"
	"github.com/maddevsio/simple-config"
)

func InitBot(config simple_config.SimpleConfig) *tgbotapi.BotAPI {
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

func ListenAndReactInUserMessages(bot *tgbotapi.BotAPI, d *diskv.Diskv, quotesFilePath string) {
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

		var rows [][]tgbotapi.KeyboardButton
		rows = append(rows, []tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton("Еще..."),
		})
		keyboard := tgbotapi.NewReplyKeyboard(rows...)
		keyboard.OneTimeKeyboard = true

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetRandomQuote(quotesFilePath))
		msg.ReplyMarkup = keyboard

		bot.Send(msg)
	}
}
