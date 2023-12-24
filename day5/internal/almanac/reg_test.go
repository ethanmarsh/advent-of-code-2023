package almanac

import (
	"fmt"
	"testing"
)

func TestParseSeeds(t *testing.T) {
	testString := "seeds: 79 14 55 13"
	expectedSeeds := []int{79, 14, 55, 13}

	actualSeeds, err := ParseSeeds(testString)
	if err != nil {
		t.Error("Failed to parse seeds")
	}

	if !IsEqual(actualSeeds, expectedSeeds) {
		t.Error("Expected Seeds do not equal actual parsed seeds")
	}
}

func TestParseMap(t *testing.T) {
	testStrings := []string{"50 98 2", "52 50 48", ""}
	expectedMaps := []Map{{50, 98, 2}, {52, 50, 48}}
	actualMaps := ParseMap(testStrings, 0)

	if len(actualMaps) != 2 {
		t.Error("length of actual maps should be 2")
	}

	if !Equals(expectedMaps, actualMaps) {
		fmt.Printf("Expected Maps: %v\n", expectedMaps)
		fmt.Printf("Actual Maps: %v\n", actualMaps)
		t.Error("expected maps does not equal actual parsed maps")
	}
}
