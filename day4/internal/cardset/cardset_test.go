package cardset

import (
	"fmt"
	"testing"
)

func createTestCardset() Cardset {
	testCard1 := Card{1, []int{1, 2}, []int{2, 3}}
	testCard2 := Card{2, []int{3, 4}, []int{5, 4, 6, 3}}
	return Cardset{[]Card{testCard1, testCard2}}
}

func createTestCardsetStrings() []string {
	return []string{"Card 1: 1 2 | 2 3", "Card 2: 3 4 | 5 4 6 3"}
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
	expectedPoints := 3
	actualPoints := testCardset.Points()
	if expectedPoints != actualPoints {
		fmt.Printf("Expected Points: %d\n", expectedPoints)
		fmt.Printf("Actual Points: %d\n", actualPoints)
		t.Error("Expected points doesn't equal actual points")
	}
}
