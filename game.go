package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const fileName = "google-10000-english-usa-no-swears-long.txt"
const baseHP = 5

type GameState struct {
	HP      int
	Word    string
	Letters [26]int
	Length  int
}

func extractWords() ([]string, error) {
	words := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read the file: %v", err)
	}

	return words, nil

}

func generateGame() (*GameState, error) {
	words, err := extractWords()
	if err != nil {
		return nil, fmt.Errorf("failed to extract words: %v", err)
	}
	word := words[rand.Int()%len(words)]

	var letters [26]int
	for i := range 26 {
		letters[i] = 0
	}

	return &GameState{
		HP:      baseHP,
		Word:    word,
		Letters: letters,
		Length:  len(word),
	}, nil
}

func (*GameState) game() {

}
