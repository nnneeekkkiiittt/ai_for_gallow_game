package main

type Genome struct {
	W1 [][]float64
	B1 []float64

	W2 [][]float64
	B2 []float64

	currentFiness float32
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

// temporary pad
func (*Genome) pickLetter(game GameState) int {
	return 0
}
