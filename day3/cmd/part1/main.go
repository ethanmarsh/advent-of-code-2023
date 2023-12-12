package main

import (
	"engine"
	"filereader"
	"fmt"
)

func main() {
	fmt.Println("😎 Begin Day3 Part1")

	schematic := filereader.ReadFile("test_input.txt")
	sum := engine.NewEngine(schematic)
	fmt.Printf("Found part 1 sum: %d", sum)
}
