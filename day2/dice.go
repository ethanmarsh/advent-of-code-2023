package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Dice struct {
	number int
	color  DiceColor
}

func (dice Dice) String() string {
	return fmt.Sprintf("%d %s", dice.number, dice.color.String())
}

func NewDice(diceString string) Dice {
	// fmt.Println("      Parsing Dice roll from: " + diceString)

	numberString, colorString, found := strings.Cut(diceString, " ")
	if !found {
		panic("Error parsing dice roll")
	}

	number, err := strconv.Atoi(numberString)
	if err != nil {
		panic("Error parsing dice number")
	}

	color := NewDiceColor(colorString)

	// fmt.Println("        Parsed DiceColor: " + color.String())
	dice := Dice{number, color}
	// fmt.Println("      Parsed Dice: " + dice.String())
	return dice
}
