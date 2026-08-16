package main

import (
	"fmt"
)

func main() {
	game, err := generateGame()
	if err != nil {
		fmt.Printf("could not generate game: %v", err)
		return
	}
	fmt.Print(game.Word)
}
