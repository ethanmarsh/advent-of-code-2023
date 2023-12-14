package cardset

import (
	"fmt"
	"testing"
)

func TestCreateCard(t *testing.T) {
	expectedCard := Card{1, []int{1, 2}, []int{2, 3}}
	test_string := "Card 1: 1 2 | 2 3"
	actualCard := NewCard(test_string)

	if !expectedCard.Equal(actualCard) {
		fmt.Println("Expected Card: " + expectedCard.String())
		fmt.Println("Actual Card: " + actualCard.String())
		t.Error("Expected Card does not equal actual card")
	}
}
