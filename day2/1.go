package main

import (
	"fmt"
)

func main() {
	fmt.Println("Begin Day2 Part 1")
	gameset := NewGameSet("puzzle-input.txt")
	fmt.Println("--------------------")

	bag := Bag{Dice{12, Red}, Dice{13, Green}, Dice{14, Blue}}
	possibleGames := gameset.PossibleGames(bag)

	sum := 0
	for _, game := range possibleGames {
		sum += game.id
	}

	fmt.Printf("Sum: %d\n", sum)
}
