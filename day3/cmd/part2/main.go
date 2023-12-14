package main

import (
	"engine"
	"filereader"
	"fmt"
)

func main() {
	fmt.Println("😎 Begin Day3 Part2")
	schematic := filereader.ReadFile("puzzle_input.txt")
	sum := engine.GetGearSum(schematic)
	fmt.Printf("Found part 2 sum: %d\n", sum)
	// answer is 84584891
}
