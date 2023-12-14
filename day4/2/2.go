package main

import (
	"cardset"
	"filereader"
	"fmt"
)

func main() {
	fmt.Println("👻 Running Day4 Part2")
	realfile := filereader.ReadFile("puzzle_input.txt")
	realcardset := cardset.NewCardSet(realfile)
	fmt.Printf("Total num cards: %d\n", realcardset.TotalCards())
	// answer is 14427616
}
