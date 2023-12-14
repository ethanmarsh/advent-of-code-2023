package cardset

import (
	"fmt"
	"math"
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

	prefixSlice := strings.Fields(prefix)
	if len(prefixSlice) != 2 {
		panic("Expected prefix slices to 'Card' and an int")
	}
	cardid, err := strconv.Atoi(prefixSlice[1])
	if err != nil {
		fmt.Println("Bad Card Id: " + prefix)
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

func (lhs Card) Equals(rhs Card) bool {
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

func (card Card) Points() int {
	numMatches := len(card.GetMatches())
	if numMatches == 0 {
		return 0
	}
	return int(math.Pow(2, float64(numMatches-1)))
}

func (card Card) GetMatches() []int {
	var matches []int
	for _, owned := range card.ownedNumbers {
		if intSliceContains(card.winningNumbers, owned) {
			matches = append(matches, owned)
		}
	}
	return matches
}

func intSliceContains(slice []int, num int) bool {
	for _, x := range slice {
		if x == num {
			return true
		}
	}

	return false
}
