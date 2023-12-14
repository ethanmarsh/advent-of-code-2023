package cardset

import (
	"fmt"
)

type Cardset struct {
	cards []Card
}

func NewCardSet(lines []string) Cardset {
	var cards []Card
	for _, line := range lines {
		cards = append(cards, NewCard(line))
	}
	return Cardset{cards}
}

func (cardset Cardset) String() string {
	string := "Cardset: \n"
	for _, card := range cardset.cards {
		string = fmt.Sprintf("%s\t%s\n", string, card.String())
	}
	return string
}

func (lhs Cardset) Equals(rhs Cardset) bool {
	if len(lhs.cards) != len(rhs.cards) {
		return false
	}

	for x, _ := range lhs.cards {
		if !lhs.cards[x].Equals(rhs.cards[x]) {
			return false
		}
	}
	return true
}

func (cardset Cardset) Points() int {
	var sum = 0
	for _, card := range cardset.cards {
		sum += card.Points()
	}
	return sum
}
