package gameset

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
	fmt.Println("Parsing Reveal from: " + dice)

	diceStrings := strings.Split(dice, ", ")
	diceSlices := make([]Dice, len(diceStrings))
	for _, diceString := range diceStrings {
		diceSlices = append(diceSlices, NewDice(diceString))
	}

	reveal := Reveal{diceSlices}
	return reveal
}
