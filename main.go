package main

import (
	"log"
)

func main() {
	config := NewConfig("config", "yml")
	token := config.Get("telegram-token")
	log.Printf("%v\n", token)
}