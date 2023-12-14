package main

import (
	"cardset"
	"filereader"
	"fmt"
)

func main() {
	fmt.Println("👻 Running Day4 Part1")

	testfile := filereader.ReadFile("test_input.txt")
	testcardset := cardset.NewCardSet(testfile)
	fmt.Printf("Test Card Set Points: %d\n", testcardset.Points())

	realfile := filereader.ReadFile("puzzle_input.txt")
	realcardset := cardset.NewCardSet(realfile)
	fmt.Printf("Real Card Set Points: %d\n", realcardset.Points())
	// Answer: 25004
}
