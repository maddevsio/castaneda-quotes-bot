package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"github.com/jasonlvhit/gocron"
	"github.com/puzanov/castaneda-quotes-bot/service"
)

var (
	d = service.GetStorage("./db")
	quotesFilePath = "./quotes.txt"
	config = service.NewConfig("config", "yml")
	bot = service.InitBot(config)
)

func main() {
	go service.ListenAndReactInUserMessages(bot, d)
	gocron.Every(10).Seconds().Do(func() {
		chats, _ := service.GetAllChats(d)
		messageText := service.GetRandomQuote(quotesFilePath)
		for _, chat := range chats {
			chatID := int64(chat.Id)
			msg := tgbotapi.NewMessage(chatID, messageText)
			bot.Send(msg)
		}
	})
	<- gocron.Start()
}