package main

import (
	"fmt"
	"math/rand"
	"sort"
)

const mutationScale = 0.02

func inheritParameter(first, second float64) float64 {
	value := first
	if rand.Intn(2) == 0 {
		value = second
	}
	return value + (2*rand.Float64()-1)*mutationScale
}

func crossover(weaker, stronger *Genome) *Genome {
	child := &Genome{}

	for h := range child.W1 {
		for i := range child.W1[h] {
			child.W1[h][i] = inheritParameter(weaker.W1[h][i], stronger.W1[h][i])
		}
		child.B1[h] = inheritParameter(weaker.B1[h], stronger.B1[h])
	}

	for letter := range child.W2 {
		for h := range child.W2[letter] {
			child.W2[letter][h] = inheritParameter(
				weaker.W2[letter][h], stronger.W2[letter][h],
			)
		}
		child.B2[letter] = inheritParameter(weaker.B2[letter], stronger.B2[letter])
	}

	return child
}

func nextGeneration(generation []*Genome, eliteCount int) []*Genome {
	ranked := append([]*Genome(nil), generation...)
	sort.SliceStable(ranked, func(i, j int) bool {
		return ranked[i].currentFiness > ranked[j].currentFiness
	})

	fmt.Println("Лучший воркер угадал ", ranked[0].guessedWords, " из ", gamesPerGen)

	next := make([]*Genome, len(ranked))

	for i, parent := range ranked[:eliteCount] {
		elite := *parent
		elite.currentFiness = 0
		elite.guessedWords = 0
		next[i] = &elite
	}

	for i := eliteCount; i < len(ranked); i++ {
		partner := ranked[rand.Intn(eliteCount)]
		next[i] = crossover(ranked[i], partner)
	}

	return next
}
