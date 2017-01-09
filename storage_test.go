package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestChatSaveAngGet(t *testing.T) {
	c1 := Chat{1, "Blah"}
	c1.Save()

	var c2 Chat
	c2.Id = 1
	c2.Get()

	assert.Equal(t, c2.Info, "Blah")

	var c3 Chat
	c3.Id = 2
	c3.Get()

	assert.Equal(t, c3.Info, "")
}