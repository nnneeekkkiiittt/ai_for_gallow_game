package main

import (
	"math"
	"math/rand"
)

const (
	inputSize  = 48
	hiddenSize = 100
	outputSize = 26
)

type Genome struct {
	W1 [hiddenSize][inputSize]float64
	B1 [hiddenSize]float64

	W2 [outputSize][hiddenSize]float64
	B2 [outputSize]float64

	currentFiness float32
}

func newGenome() *Genome {
	genome := &Genome{}
	const weightScale = 0.1

	for h := range genome.W1 {
		for i := range genome.W1[h] {
			genome.W1[h][i] = (2*rand.Float64() - 1) * weightScale
		}
	}

	for letter := range genome.W2 {
		for h := range genome.W2[letter] {
			genome.W2[letter][h] = (2*rand.Float64() - 1) * weightScale
		}
	}

	return genome
}

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
	Won                int // 0-1
}

func (g *Genome) pickLetter(game GameState) int {
	var input [inputSize]float64

	for i, letter := range game.Mask {
		input[i] = float64(letter+1) / 26.0
	}

	for letter, result := range game.Letters {
		if result != 0 {
			input[len(game.Mask)+letter] = 1
		}
	}

	input[inputSize-2] = float64(game.HP) / float64(baseHP)
	input[inputSize-1] = float64(game.Length) / float64(len(game.Mask))

	var hidden [hiddenSize]float64
	for h, weights := range g.W1 {
		sum := g.B1[h]
		for i, weight := range weights {
			sum += weight * input[i]
		}
		hidden[h] = math.Tanh(sum)
	}

	bestLetter := -1
	bestScore := math.Inf(-1)
	for letter, weights := range g.W2 {
		if game.Letters[letter] != 0 {
			continue
		}

		score := g.B2[letter]
		for h, weight := range weights {
			score += weight * hidden[h]
		}

		if score > bestScore {
			bestLetter = letter
			bestScore = score
		}
	}

	return bestLetter
}
