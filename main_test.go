package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConfigGet(t *testing.T) {
	config := Config{"config.test", "yml"}
	config.Init()
	token := config.Get("telegram-token")
	assert.Equal(t, token, "blah")
}