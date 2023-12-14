package cardset

import (
	"fmt"
	"testing"
)

func createTestCard() Card {
	return Card{1, []int{1, 2}, []int{2, 3}}
}

func createTestCard2() Card {
	return Card{2, []int{3, 4}, []int{5, 4, 6, 3}}
}

func createTestCard3() Card {
	return Card{3, []int{7, 8, 9, 10, 11}, []int{11, 8, 10, 7}}
}

func createTestCardString() string {
	return "Card 1: 1 2 | 2 3"
}

func TestCreateCard(t *testing.T) {
	expectedCard := createTestCard()
	test_string := createTestCardString()
	actualCard := NewCard(test_string)

	if !expectedCard.Equals(actualCard) {
		fmt.Println("Expected Card: " + expectedCard.String())
		fmt.Println("Actual Card: " + actualCard.String())
		t.Error("Expected Card does not equal actual card")
	}
}

func TestCreateCardWithExtraSpace(t *testing.T) {
	expectedCard := createTestCard()
	test_string := "Card   1: 1 2 | 2 3"
	actualCard := NewCard(test_string)

	if !expectedCard.Equals(actualCard) {
		fmt.Println("Expected Card: " + expectedCard.String())
		fmt.Println("Actual Card: " + actualCard.String())
		t.Error("Expected Card does not equal actual card")
	}
}

func TestFindMatches(t *testing.T) {
	testCard := createTestCard2()
	expectedMatches := []int{4, 3}
	actualMatches := testCard.GetMatches()
	if !equalSlices(expectedMatches, actualMatches) {
		fmt.Printf("Expected matches: %v\n", expectedMatches)
		fmt.Printf("Actual matches: %v\n", actualMatches)
		t.Error("Expected matching numbers did not equal actual matching numbers")
	}
}

func TestPoints(t *testing.T) {
	testCard := createTestCard3()
	expectedPoints := 8
	actualPoints := testCard.Points()
	if expectedPoints != actualPoints {
		fmt.Printf("Expected points: %d\n", expectedPoints)
		fmt.Printf("Actual points: %d\n", actualPoints)
		t.Error("Expected points do not match actual points")
	}
}

func TestPointsNoMatches(t *testing.T) {
	testCard := Card{1, []int{1, 2}, []int{3, 4}}
	expectedPoints := 0
	actualPoints := testCard.Points()
	if expectedPoints != actualPoints {
		fmt.Printf("Expected points: %d\n", expectedPoints)
		fmt.Printf("Actual points: %d\n", actualPoints)
		t.Error("Expected points do not match actual points")
	}
}

func equalSlices(lhs []int, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for x, _ := range lhs {
		if lhs[x] != rhs[x] {
			return false
		}
	}

	return true
}
