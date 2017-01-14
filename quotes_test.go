package main

import (
	"testing"
	"log"
)

func TestGetRandomQuote(t *testing.T) {
	quote := GetRandomQuote("./quotes.txt")
	log.Printf("%v \n", quote)
}
