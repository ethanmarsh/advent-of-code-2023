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
	fmt.Println("Parsing Game from: " + line)
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
	fmt.Println("Parsed Game: " + game.String())
	return game
}
