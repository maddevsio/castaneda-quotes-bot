package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/puzanov/castaneda-quotes-bot/service"
	"github.com/maddevsio/simple-config"
	"time"
	"log"
)

var (
	config = simple_config.NewSimpleConfig("config", "yml")
	d = service.GetStorage(config.Get("db").(string))
	quotesFilePath = "./quotes.txt"
	bot = service.InitBot(config)
)

func main() {
	go service.ListenAndReactInUserMessages(bot, d, quotesFilePath)
	changeGocronTimezone()
	gocron.Every(10).Seconds().Do(func() { // this is for testing
	//gocron.Every(1).Day().At(config.Get("send-time").(string)).Do(func() {
		chats, _ := service.GetAllChats(d)
		messageText := service.GetRandomQuote(quotesFilePath)
		for _, chat := range chats {
			service.SendMessageWithKeyboard(bot, messageText, int64(chat.Id))
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