package main

type Genome struct {
	W1 [][]float64
	B1 []float64

	W2 [][]float64
	B2 []float64

	currentFiness float32
}

// temporary pad
func (*Genome) pickLetter(game GameState) int {
	return 0
}
