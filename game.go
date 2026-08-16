package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const fileName = "google-10000-english-usa-no-swears-long.txt"
const baseHP = 5

type GameState struct {
	HP                 int
	Word               string
	Letters            [26]int
	Length             int
	Mask               [20]int
	Genome             *Genome
	CountOfRepetitions int
	LettersRemained    int
	Attempt            int
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

	// letter was picked - score from -1 to 1
	//letter was NOT picked - score = 0
	var letters [26]int
	for i := range 26 {
		letters[i] = 0
	}

	var mask [20]int
	for i := range 20 {
		mask[i] = -1
	}

	return &GameState{
		HP:                 baseHP,
		Word:               word,
		Letters:            letters,
		Length:             len(word),
		Mask:               mask,
		LettersRemained:    len(word),
		CountOfRepetitions: 0,
		Attempt:            0,
	}, nil
}

func game(state *GameState) bool {
	if state.HP == 0 {
		return false
	}
	if state.LettersRemained == 0 {
		return true
	}
	state.Attempt++
	letter := state.Genome.pickLetter(*state)
	runeLetter := rune(letter + 97)
	if state.Letters[letter] != 0 {
		state.HP--
		state.CountOfRepetitions++
	} else {
		if strings.ContainsRune(state.Word, runeLetter) {
			for i := range state.Length {
				if state.Word[i] == byte(runeLetter) {
					state.Mask[i] = letter
					state.Letters[letter]++
					state.LettersRemained--
				}
			}
		} else {
			state.HP--
			state.Letters[letter] = -1
		}
	}
	return game(state)
}
