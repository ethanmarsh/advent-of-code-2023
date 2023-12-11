package main

import (
	"fmt"
	"math"
)

type Bag struct {
	red, green, blue Dice
}

func (bag Bag) String() string {
	return fmt.Sprintf("Bag - Red:%s, Green:%s, Blue: %s", bag.red, bag.green, bag.blue)
}

func EmptyBag() Bag {
	return Bag{Dice{0, Red}, Dice{0, Green}, Dice{0, Blue}}
}

func (lhs Bag) CombineMax(rhs Bag) Bag {
	maxRed := math.Max(float64(lhs.red.number), float64(rhs.red.number))
	maxGreen := math.Max(float64(lhs.green.number), float64(rhs.green.number))
	maxBlue := math.Max(float64(lhs.blue.number), float64(rhs.blue.number))

	return Bag{Dice{int(maxRed), Red}, Dice{int(maxGreen), Green}, Dice{int(maxBlue), Blue}}
}

func (bag Bag) Power() int {
	return bag.red.number * bag.green.number * bag.blue.number
}
