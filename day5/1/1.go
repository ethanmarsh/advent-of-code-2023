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
}
