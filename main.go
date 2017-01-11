package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"github.com/jasonlvhit/gocron"
)

var (
	token string
	bot *tgbotapi.BotAPI
	config Config
	d = GetStorage("./db")
)

func task() {
	chats, _ := GetAllChats(d)
	for _, chat := range chats {
		chatID := int64(chat.Id)
		messageText := "Bot initialized message"
		msg := tgbotapi.NewMessage(chatID, messageText)
		bot.Send(msg)
	}
}

func initBot() {
	config = NewConfig("config", "yml")
	token = config.Get("telegram-token").(string)
	log.Printf("%v\n", token)

	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
}

func listenAndReactInUserMessages() {
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

func runCronToSendScheduledMessages() {
	gocron.Every(10).Seconds().Do(task)
	<- gocron.Start()
}

func main() {
	initBot()
	go listenAndReactInUserMessages()
	runCronToSendScheduledMessages()
}