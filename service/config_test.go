package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConfigGet(t *testing.T) {
	config := NewConfig("../config.test", "yml")
	token := config.Get("telegram-token")
	assert.Equal(t, token, "blah")
}