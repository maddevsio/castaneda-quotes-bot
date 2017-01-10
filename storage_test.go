package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/gommon/log"
)

var d = GetStorage("./db")

func TestChatSaveAngGet(t *testing.T) {
	c1 := Chat{1, "Blah"}
	c1.Save(d)

	var c2 Chat
	c2.Id = 1
	c2.Get(d)

	assert.Equal(t, c2.Info, "Blah")

	var c3 Chat
	c3.Id = 2
	c3.Get(d)

	assert.Equal(t, c3.Info, "")
}

func TestGetAllKeys(t *testing.T) {
	c1 := Chat{88, "Test"}
	c1.Save(d)
	chats, _ := GetAllChats(d)
	for _, chat := range chats {
		log.Printf("CHAT %v\n", chat)
		log.Printf("CHAT %v\n", chat.Id)
		log.Printf("CHAT %v\n", chat.Info)
	}
}