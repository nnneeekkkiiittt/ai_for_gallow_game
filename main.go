package main

import "fmt"

const (
	populationSize = 200
	eliteCount     = 100
	gamesPerGen    = 50
	generations    = 100
)

func main() {
	generation := make([]*Genome, populationSize)
	for i := range generation {
		generation[i] = newGenome()
	}

	for round := 1; round <= generations; round++ {
		runGeneration(generation, gamesPerGen)

		best := generation[0]
		for _, genome := range generation[1:] {
			if genome.currentFiness > best.currentFiness {
				best = genome
			}
		}
		fmt.Printf("Поколение %d, лучший результат %.2f\n", round, best.currentFiness)

		if round < generations {
			generation = nextGeneration(generation, eliteCount)
		}
	}
}
