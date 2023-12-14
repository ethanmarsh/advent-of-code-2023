package cardset

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers []int
	ownedNumbers   []int
}

func NewCard(line string) Card {
	prefix, suffix, found := strings.Cut(line, ":")
	if !found {
		panic("Can't parse card id")
	}

	id, foundId := strings.CutPrefix(prefix, "Card ")
	cardid, err := strconv.Atoi(id)
	if !foundId || err != nil {
		panic("error parsing card id")
	}

	winningNumbersStr, ownedNumbersStr, foundSplit := strings.Cut(suffix, "|")
	if !foundSplit {
		panic("error parsing winning and owned numbers")
	}

	winningNumbersSlices := strings.Fields(winningNumbersStr)
	ownedNumbersSlices := strings.Fields(ownedNumbersStr)

	return Card{cardid, stringSliceToIntSlice(winningNumbersSlices), stringSliceToIntSlice(ownedNumbersSlices)}
}

func (card Card) String() string {
	return fmt.Sprintf("Card %d: %v | %v", card.id, card.winningNumbers, card.ownedNumbers)
}

func (lhs Card) Equal(rhs Card) bool {
	if lhs.id != rhs.id {
		return false
	}

	if len(lhs.winningNumbers) != len(rhs.winningNumbers) {
		return false
	}

	if len(lhs.ownedNumbers) != len(rhs.ownedNumbers) {
		return false
	}

	for x, _ := range lhs.winningNumbers {
		if lhs.winningNumbers[x] != rhs.winningNumbers[x] {
			return false
		}
	}

	for x, _ := range lhs.ownedNumbers {
		if lhs.ownedNumbers[x] != rhs.ownedNumbers[x] {
			return false
		}
	}

	return true
}

func stringSliceToIntSlice(stringSlice []string) []int {
	var intSlice []int
	for _, string := range stringSlice {
		int, err := strconv.Atoi(string)
		check(err)
		intSlice = append(intSlice, int)
	}
	return intSlice
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
