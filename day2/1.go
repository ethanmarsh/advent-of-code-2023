package main

import (
	"fmt"
)

func main() {
	fmt.Println("Begin Day2ls Part 1")
	gameset := NewGameSet("test-input.txt")
	fmt.Println("--------------------")
	fmt.Println(gameset)

	test1 := NewDiceColor("red")
	test2 := NewDiceColor("red")
	if test1 == test2 {
		fmt.Println("DiceColor equality is working")
	} else {
		fmt.Println("DiceColor equality is NOT working")
	}
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Filter[T any](ts []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, t := range ts {
		if fn(t) {
			result = append(result, t)
		}
	}
	return result
}
