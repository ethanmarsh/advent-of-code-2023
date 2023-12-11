package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	id      int
	reveals []Reveal
}

func (game Game) String() string {
	return fmt.Sprintf("Game %d: %s", game.id, game.reveals)
}

func NewGame(line string) Game {
	// fmt.Println("Parsing Game from: " + line)
	gameline, revealline, found := strings.Cut(line, ": ")
	if !found {
		panic("error parsing line: " + line)
	}

	id, foundId := strings.CutPrefix(gameline, "Game ")
	gameid, err := strconv.Atoi(id)
	if !foundId || err != nil {
		panic("error parsing game id")
	}

	revealSlices := strings.Split(revealline, "; ")

	reveals := make([]Reveal, len(revealSlices))
	for i, revealStr := range revealSlices {
		reveals[i] = NewReveal(revealStr)
	}

	game := Game{gameid, reveals}
	// fmt.Println("Parsed Game: " + game.String())
	return game
}

func (lhs Game) Equal(rhs Game) bool {
	if lhs.id != rhs.id {
		return false
	}
	if len(lhs.reveals) != len(rhs.reveals) {
		return false
	}
	for i, lReveal := range lhs.reveals {
		rReveal := rhs.reveals[i]
		if !lReveal.Equal(rReveal) {
			return false
		}
	}

	return true
}

func (game Game) IsPossible(bag Bag) bool {
	return All(game.reveals, func(r Reveal) bool {
		return r.IsPossible(bag)
	})
}

func (game Game) MinimumDice() Bag {
	minimumAll := EmptyBag()
	for _, reveal := range game.reveals {
		minimumAll = minimumAll.CombineMax(reveal.MinimumDice())
	}

	return minimumAll
}
