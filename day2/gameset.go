package main

import (
	"bufio"
	"fmt"
	"os"
)

type GameSet struct {
	games []Game
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (gameset GameSet) String() string {
	return fmt.Sprintf("%s", gameset.games)
}

func NewGameSet(filename string) GameSet {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	games := make([]Game, len(lines))
	for i, line := range lines {
		games[i] = NewGame(line)
	}

	gameset := GameSet{games}
	return gameset
}

func (lhs GameSet) Equal(rhs GameSet) bool {
	if len(lhs.games) != len(rhs.games) {
		return false
	}
	for i, lGame := range lhs.games {
		rGame := rhs.games[i]
		if !lGame.Equal(rGame) {
			return false
		}
	}

	return true
}

func (gameset GameSet) PossibleGames(bag Bag) []Game {
	return Filter(gameset.games, func(g Game) bool {
		return g.IsPossible(bag)
	})
}

func (gameset GameSet) SumOfMinimumBagPower() int {
	sum := 0
	for _, game := range gameset.games {
		sum += game.MinimumDice().Power()
	}
	return sum
}
