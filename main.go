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
	//gocron.Every(10).Seconds().Do(func() {
	gocron.Every(1).Day().At(config.GetString("send-time")).Do(func() {
		log.Printf("Scheduled start in %v", config.GetString("send-time"))
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