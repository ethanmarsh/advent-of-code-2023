package main

import (
	"testing"
)

func NewTestDice() Dice {
	return Dice{1, Red}
}

func NewTestReveal() Reveal {
	dice := []Dice{
		{9, Red},
		{1, Green},
		{5, Blue},
	}
	return Reveal{dice}
}

func NewTestGame() Game {
	// Game 1
	reveal11 := Reveal{[]Dice{
		{3, Blue},
		{4, Red},
	}}
	reveal12 := Reveal{[]Dice{
		{1, Red},
		{2, Green},
		{6, Blue},
	}}
	reveal13 := Reveal{[]Dice{
		{2, Green},
	}}
	reveals := []Reveal{reveal11, reveal12, reveal13}
	return Game{1, reveals}
}

func NewTestGameSet() GameSet {
	// Game 1
	reveal11 := Reveal{[]Dice{
		{3, Blue},
		{4, Red},
	}}
	reveal12 := Reveal{[]Dice{
		{1, Red},
		{2, Green},
		{6, Blue},
	}}
	reveal13 := Reveal{[]Dice{
		{2, Green},
	}}
	reveals1 := []Reveal{reveal11, reveal12, reveal13}
	game1 := Game{1, reveals1}

	// Game 2
	reveal21 := Reveal{[]Dice{
		{1, Blue},
		{2, Green},
	}}
	reveal22 := Reveal{[]Dice{
		{3, Green},
		{4, Blue},
		{1, Red},
	}}
	reveal23 := Reveal{[]Dice{
		{1, Green},
		{1, Blue},
	}}
	reveals2 := []Reveal{reveal21, reveal22, reveal23}
	game2 := Game{2, reveals2}

	// Game 3
	reveal31 := Reveal{[]Dice{
		{8, Green},
		{6, Blue},
		{20, Red},
	}}
	reveal32 := Reveal{[]Dice{
		{5, Blue},
		{4, Red},
		{13, Green},
	}}
	reveal33 := Reveal{[]Dice{
		{5, Green},
		{1, Red},
	}}
	reveals3 := []Reveal{reveal31, reveal32, reveal33}
	game3 := Game{3, reveals3}

	// Game 4
	reveal41 := Reveal{[]Dice{
		{1, Green},
		{3, Red},
		{6, Blue},
	}}
	reveal42 := Reveal{[]Dice{
		{3, Green},
		{6, Red},
	}}
	reveal43 := Reveal{[]Dice{
		{3, Green},
		{15, Blue},
		{14, Red},
	}}
	reveals4 := []Reveal{reveal41, reveal42, reveal43}
	game4 := Game{4, reveals4}

	// Game 5
	// Game 5: 2 blue, 1 red, 2 green
	reveal51 := Reveal{[]Dice{
		{6, Red},
		{1, Blue},
		{3, Green},
	}}
	reveal52 := Reveal{[]Dice{
		{2, Blue},
		{1, Red},
		{2, Green},
	}}
	reveals5 := []Reveal{reveal51, reveal52}
	game5 := Game{5, reveals5}

	return GameSet{[]Game{game1, game2, game3, game4, game5}}
}

func TestDiceColorEquality(t *testing.T) {
	test1 := Red
	test2 := Red
	if test1 != test2 {
		t.Error("DiceColor equality is NOT working")
	}
}

func TestDiceEquality(t *testing.T) {
	test1 := NewTestDice()
	test2 := NewTestDice()
	if test1 != test2 {
		t.Error("Dice equality is NOT working")
	}
}

func TestRevealEquality(t *testing.T) {
	test1 := NewTestReveal()
	test2 := NewTestReveal()
	if !test1.Equal(test2) {
		t.Error("Reveal equality is NOT working")
	}
}

func TestGameEquality(t *testing.T) {
	test1 := NewTestGame()
	test2 := NewTestGame()
	if !test1.Equal(test2) {
		t.Error("Game equality is NOT working")
	}
}

func TestGameSetEquality(t *testing.T) {
	test1 := NewTestGameSet()
	test2 := NewTestGameSet()
	if !test1.Equal(test2) {
		t.Error("GameSet equality is NOT working")
	}
}

func TestPossibleGames(t *testing.T) {
	test1 := NewTestGameSet()
	bag := Bag{Dice{12, Red}, Dice{13, Green}, Dice{14, Blue}}
	possibleGames := test1.PossibleGames(bag)

	sum := 0
	for _, game := range possibleGames {
		sum += game.id
	}

	if sum != 8 {
		t.Error("Possible game calculation is wrong")
	}
}

func TestMinimumDice(t *testing.T) {
	testGame := NewTestGame()
	expectedBag := Bag{Dice{4, Red}, Dice{2, Green}, Dice{6, Blue}}
	if testGame.MinimumDice() != expectedBag {
		t.Error("Minimum bag calculation is wrong")
	}
}

func TestPower(t *testing.T) {
	testGame := NewTestGame()
	expectedPower := 48
	if testGame.MinimumDice().Power() != expectedPower {
		t.Error("Minimum bag power calculation is wrong")
	}
}

func TestPowerSum(t *testing.T) {
	testGameSet := NewTestGameSet()
	expectedPowerSum := 2286
	if testGameSet.SumOfMinimumBagPower() != expectedPowerSum {
		t.Error("Sum of minimum bag power calculation is wrong")
	}
}
