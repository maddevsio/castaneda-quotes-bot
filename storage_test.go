package main

import (
	"testing"
//	"github.com/stretchr/testify/assert"
)

func TestChatSaveAngGet(t *testing.T) {
	c1 := Chat{1, "blah"}
	c1.Save()

	c2 := Chat{1, ""}
	c2.Get()

	//assert.Equal(t, c2.info, "Blah")
}