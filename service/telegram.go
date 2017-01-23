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
	bot.Debug = false
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
		log.Printf("Saving... %v", chat)
		err = chat.Save(d)
		if err != nil {
			log.Printf("Error saving data: %v", err)
		}
		SendMessageWithKeyboard(bot, GetRandomQuote(quotesFilePath), update.Message.Chat.ID)
	}
}

func SendMessageWithKeyboard(bot *tgbotapi.BotAPI, message string, chatID int64) {
	var rows [][]tgbotapi.KeyboardButton
	rows = append(rows, []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Еще..."),
	})
	keyboard := tgbotapi.NewReplyKeyboard(rows...)
	keyboard.OneTimeKeyboard = true

	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyMarkup = keyboard

	log.Printf("Sending this msg: %v", msg)
	_, err := bot.Send(msg)
	if (err != nil) {
		log.Printf("Error sending message: %v", err)
		return
	}
	log.Print("Success")
}
