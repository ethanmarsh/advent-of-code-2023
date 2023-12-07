package gameset

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

func New(filename string) GameSet {
	fmt.Println("Parsing GameSet")

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	games := make([]Game, len(lines))
	for _, line := range lines {
		games = append(games, NewGame(line))
	}

	gameset := GameSet{games}
	return gameset
}
