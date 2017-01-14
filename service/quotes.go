package service

import (
	"os"
	"bufio"
	"log"
	"math/rand"
	"time"
)

func GetRandomQuote(filePath string) string {
	quotes, err := readLines(filePath)
	if err != nil {
		log.Fatalf("Error reading quotes file: %s \n", err)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	r := rand.Intn(len(quotes))
	return quotes[r]
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}