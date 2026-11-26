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

func extractWords() []string {
	words := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return words

}

func generateGame(words []string) GameState {
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

	return GameState{
		HP:                 baseHP,
		Word:               word,
		Letters:            letters,
		Length:             len(word),
		Mask:               mask,
		LettersRemained:    len(word),
		CountOfRepetitions: 0,
		Attempt:            0,
	}
}

func assessGenome(genome *Genome, state GameState) {
	progress := 1 - float32(state.LettersRemained)/float32(state.Length)
	genome.currentFiness += 100*float32(state.Won) +
		50*progress*progress +
		5*float32(state.HP) -
		20*float32(state.CountOfRepetitions)
}

func runGame(genome *Genome, state GameState) GameState {
	for {
		if state.HP == 0 {
			state.Won = 0
			assessGenome(genome, state)
			return state
		}
		if state.LettersRemained == 0 {
			state.Won = 1
			assessGenome(genome, state)
			return state
		}

		letter := genome.pickLetter(state)
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

	}
}

func runGeneration(generation []*Genome, circles int) {
	words := extractWords()
	for i := range circles {
		fmt.Println("Начинаем круг ", i+1)
		guessed := 0
		state := generateGame(words)
		fmt.Println("Загадано слово: ", strings.ToUpper(state.Word))
		for k, genome := range generation {
			resState := runGame(genome, state)
			var bonChar string
			if resState.Won == 1 {
				guessed++
				bonChar = " +++!!!"
			}
			fmt.Println("Воркер номер ", k+1, ": ", formatMask(resState.Mask, resState.Length), bonChar)
		}
		fmt.Println("Слово ", strings.ToUpper(state.Word), " угадало ", guessed, " воркеров из ", len(generation))
	}
}

func formatMask(mask [20]int, length int) string {
	var s strings.Builder
	alphabet := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := range length {
		c := mask[i]
		if c == -1 {
			s.WriteString("_")
		} else {
			s.WriteString(string(alphabet[mask[i]]))
		}
	}
	return s.String()
}
