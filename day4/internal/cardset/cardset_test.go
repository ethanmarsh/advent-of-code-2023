package cardset

import (
	"fmt"
	"testing"
)

func createTestCardset() Cardset {
	testCard1 := Card{1, []int{1, 2}, []int{2, 3}, 1}                    // 1 match, 1 card
	testCard2 := Card{2, []int{3, 4}, []int{5, 4, 6, 3}, 1}              // 2 matches, 2 cards
	testCard3 := Card{3, []int{7, 8, 9, 10, 11}, []int{11, 8, 10, 7}, 1} // 4 matches, 3 cards
	return Cardset{[]Card{testCard1, testCard2, testCard3}}
}

func createTestCardsetStrings() []string {
	return []string{"Card 1: 1 2 | 2 3", "Card 2: 3 4 | 5 4 6 3", "Card 3: 7 8 9 10 11 | 11 8 10 7"}
}

func createFullTestCardsetStrings() []string {
	return []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", // 4 matches, 1 card   // total 1
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", // 2 matches, 2 cards  // total 3
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", // 2 matches, 4 cards  // total 7
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", // 1 match,   8 cards  // total 15
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", // 0 matches, 14 cards // total 29
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", // 0 matches, 1 card   // total 30
	}
}

func TestCreateCardset(t *testing.T) {
	expectedCardset := createTestCardset()
	test_string := createTestCardsetStrings()
	actualCardset := NewCardSet(test_string)

	if !expectedCardset.Equals(actualCardset) {
		fmt.Println("Expected Cardset: " + expectedCardset.String())
		fmt.Println("Actual Cardset: " + actualCardset.String())
		t.Error("Expected card set does not equal actual card set")
	}
}

func TestCardsetPoints(t *testing.T) {
	testCardset := createTestCardset()
	expectedPoints := 11
	actualPoints := testCardset.Points()
	if expectedPoints != actualPoints {
		fmt.Printf("Expected Points: %d\n", expectedPoints)
		fmt.Printf("Actual Points: %d\n", actualPoints)
		t.Error("Expected points doesn't equal actual points")
	}
}

func TestCardsetCopiesSimple(t *testing.T) {
	testCardset := createTestCardset()
	expectedTotalCards := 6
	actualTotalCards := testCardset.TotalCards()
	if expectedTotalCards != actualTotalCards {
		fmt.Printf("Expected total cards: %d\n", expectedTotalCards)
		fmt.Printf("Actual total cards: %d\n", actualTotalCards)
		t.Error("Simple: Expected total cards don't equal actual total cards")
	}
}

func TestCardsetCopiesFull(t *testing.T) {
	strings := createFullTestCardsetStrings()
	testCardset := NewCardSet(strings)
	expectedTotalCards := 30
	actualTotalCards := testCardset.TotalCards()
	if expectedTotalCards != actualTotalCards {
		fmt.Printf("Expected total cards: %d\n", expectedTotalCards)
		fmt.Printf("Actual total cards: %d\n", actualTotalCards)
		t.Error("Full: Expected total cards don't equal actual total cards")
	}
}
