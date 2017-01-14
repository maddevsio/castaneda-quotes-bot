package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetRandomQuote(t *testing.T) {
	// проверяем не выходит ли наш ранд за пределы списка
	for i := 0; i < 10000; i++ {
		quote := GetRandomQuote("./quotes.txt")
		notEmpty := len(quote) > 0
		assert.True(t, notEmpty)
	}
}
