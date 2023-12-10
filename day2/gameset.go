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
