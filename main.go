package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"github.com/jasonlvhit/gocron"
	"github.com/puzanov/castaneda-quotes-bot/service"
	"time"
	"log"
)

var (
	config = service.NewConfig("config", "yml")
	d = service.GetStorage(config.Get("db").(string))
	quotesFilePath = "./quotes.txt"
	bot = service.InitBot(config)
)

func main() {
	go service.ListenAndReactInUserMessages(bot, d)
	changeGocronTimezone()
	//gocron.Every(10).Seconds().Do(func() { // this is for testing
	gocron.Every(1).Day().At(config.Get("send-time").(string)).Do(func() {
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

func changeGocronTimezone() {
	location, err := time.LoadLocation(config.Get("timezone").(string))
	if err != nil {
		log.Fatalf("Unfortunately can't load a location: %v", err)

	} else {
		gocron.ChangeLoc(location)
	}
}