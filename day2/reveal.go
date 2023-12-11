package main

import (
	"fmt"
	"strings"
)

type Reveal struct {
	diceSet []Dice
}

func (reveal Reveal) String() string {
	return fmt.Sprintf("%s", reveal.diceSet)
}

func NewReveal(dice string) Reveal {
	// fmt.Println("    Parsing Reveal from: " + dice)

	diceStrings := strings.Split(dice, ", ")
	diceSlices := make([]Dice, len(diceStrings))
	for i, diceString := range diceStrings {
		diceSlices[i] = NewDice(diceString)
	}

	reveal := Reveal{diceSlices}
	// fmt.Println("    Parsed Reveal: " + reveal.String())
	return reveal
}

func (lhs Reveal) Equal(rhs Reveal) bool {
	if len(lhs.diceSet) != len(rhs.diceSet) {
		return false
	}
	for i, lDice := range lhs.diceSet {
		rDice := rhs.diceSet[i]
		if lDice != rDice {
			return false
		}
	}

	return true
}

func (reveal Reveal) IsPossible(bag Bag) bool {
	for _, dice := range reveal.diceSet {
		switch dice.color {
		case Red:
			if dice.number > bag.red.number {
				return false
			}
		case Green:
			if dice.number > bag.green.number {
				return false
			}
		case Blue:
			if dice.number > bag.blue.number {
				return false
			}
		}
	}

	return true
}

func (reveal Reveal) MinimumDice() Bag {
	red := 0
	green := 0
	blue := 0
	for _, dice := range reveal.diceSet {
		switch dice.color {
		case Red:
			if dice.number > red {
				red = dice.number
			}
		case Green:
			if dice.number > green {
				green = dice.number
			}
		case Blue:
			if dice.number > blue {
				blue = dice.number
			}
		}
	}

	return Bag{Dice{red, Red}, Dice{green, Green}, Dice{blue, Blue}}
}
