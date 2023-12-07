package main

import (
	"day2/gameset"
	"fmt"
)

func main() {
	fmt.Println("Begin Day2ls Part 1")
	gameset := gameset.New("test-input.txt")
	fmt.Println("--------------------")
	fmt.Println(gameset)
}
