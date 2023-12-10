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
