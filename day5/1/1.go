package main

import (
	"almanac"
	"filereader"
	"fmt"
)

func main() {
	fmt.Println("😵‍💫 Beginning Day5 Part1")

	testLines := filereader.ReadFile("test_input.txt")
	testAlmanac := almanac.NewAlmanac(testLines)
	fmt.Printf("Test Almanac Lowest Location: %d\n", testAlmanac.GetLowestSeedLocation())

	realLines := filereader.ReadFile("puzzle_input.txt")
	realAlmanac := almanac.NewAlmanac(realLines)
	fmt.Printf("Real Almanac Lowest Location: %d\n", realAlmanac.GetLowestSeedLocation())
}
