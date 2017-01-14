package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
)

var dTest = GetStorage("../db_test")

func TestChatSaveAngGet(t *testing.T) {
	c1 := Chat{1, "Blah"}
	c1.Save(dTest)

	var c2 Chat
	c2.Id = 1
	c2.Get(dTest)

	assert.Equal(t, c2.UserName, "Blah")

	var c3 Chat
	c3.Id = 2
	c3.Get(dTest)

	assert.Equal(t, c3.UserName, "")
}

func TestGetAllKeys(t *testing.T) {
	c1 := Chat{88, "Test"}
	c1.Save(dTest)
	chats, _ := GetAllChats(dTest)
	for _, chat := range chats {
		log.Printf("CHAT %v\n", chat)
		log.Printf("CHAT %v\n", chat.Id)
		log.Printf("CHAT %v\n", chat.UserName)
	}
}