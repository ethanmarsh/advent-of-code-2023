package engine

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
)

type NumberIndex struct {
	number string
	coord  Coordinate
}

func (numberIndex NumberIndex) String() string {
	return fmt.Sprintf("%s at %s", numberIndex.number, numberIndex.coord.String())
}

type schematic = [][]string

type Coordinate struct {
	x int
	y int
}

func (coord Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", coord.x, coord.y)
}

func charAt(sch schematic, coord Coordinate) (string, error) {
	if coord.x < 0 || coord.x >= len(sch) || coord.y < 0 || coord.y >= len(sch[coord.x]) {
		return "", errors.New("coord out of bounds")
	}

	return sch[coord.x][coord.y], nil
}

func (numberIndex NumberIndex) combine(char string) NumberIndex {
	numberIndex.number = numberIndex.number + char
	return numberIndex
}

func combine(char string, numberIndex NumberIndex) NumberIndex {
	numberIndex.number = char + numberIndex.number
	return numberIndex
}

func findNumberBeginning(sch schematic, coord Coordinate) NumberIndex {
	// fmt.Printf("Find Number Beginning - Coord:%d,%d\n", coord.x, coord.y)
	char, err := charAt(sch, coord)
	if err != nil || !IsDigit(char) {
		return NumberIndex{"", Coordinate{coord.x, coord.y + 1}}
	}

	result := findNumberBeginning(sch, Coordinate{coord.x, coord.y - 1}).combine(char)
	// fmt.Println("Find Number Beginning - " + result)
	return result
}

func findNumberEnd(sch schematic, coord Coordinate) NumberIndex {
	// fmt.Printf("Find Number End - Coord:%d,%d\n", coord.x, coord.y)
	char, err := charAt(sch, coord)
	if err != nil || !IsDigit(char) {
		return NumberIndex{"", Coordinate{coord.x, coord.y - 1}}
	}

	result := combine(char, findNumberEnd(sch, Coordinate{coord.x, coord.y + 1}))
	// fmt.Println("Find Number End - " + result)
	return result
}

func findFullNumber(sch schematic, coord Coordinate) NumberIndex {
	// fmt.Printf("Find Full Number - Coord:%d,%d\n", coord.x, coord.y)
	char, err := charAt(sch, coord)
	if err != nil || !IsDigit(char) {
		return NumberIndex{"", Coordinate{}}
	}

	beginning := findNumberBeginning(sch, Coordinate{coord.x, coord.y - 1})
	end := findNumberEnd(sch, Coordinate{coord.x, coord.y + 1})
	combined := beginning.number + char + end.number
	combinedNumberIndex := NumberIndex{combined, beginning.coord}
	// fmt.Println("Combined: " + combined)
	return combinedNumberIndex
}

func findAdjacentDigit(sch schematic, coord Coordinate) []NumberIndex {
	// fmt.Println("Find Adjacent Digit - Coord: " + coord.String())

	var adjacentNumbers []NumberIndex
	start := coord.x - 1
	end := coord.x + 2
	for x := start; x < end; x++ {
		for y := coord.y - 1; y < coord.y+2; y++ {
			newCoord := Coordinate{x, y}
			// fmt.Printf("Checking Coord:%d,%d\n", newCoord.x, newCoord.y)
			char, err := charAt(sch, newCoord)
			if err != nil {
				break
			}

			if IsDigit(char) {
				// find full number
				fullNumberStr := findFullNumber(sch, newCoord)
				fmt.Println("Found full number " + fullNumberStr.String())
				adjacentNumbers = append(adjacentNumbers, fullNumberStr)
			}
		}
	}
	return adjacentNumbers
}

func removeDuplicates(array []NumberIndex) []NumberIndex {
	// fmt.Printf("STARTING ARRAY: %v\n", array)
	var result []NumberIndex
	for _, numberIndex := range array {
		if !slices.Contains(result, numberIndex) {
			result = append(result, numberIndex)
		}
	}
	// fmt.Printf("DEDUPED ARRAY: %v\n", result)
	return result
}

func sum(array []NumberIndex) int {
	var sum int
	for _, num := range array {
		convertedNum, error := strconv.Atoi(num.number)
		if error != nil {
			panic("bad conversion")
		}
		sum += convertedNum
	}
	return sum
}

func NewEngine(sch schematic) int {
	var allAdjacentNumbers []NumberIndex
	for i, line := range sch {
		for j, character := range line {
			if IsSymbol(character) {
				// fmt.Printf("Found symbol %s\n", character)
				numbers := findAdjacentDigit(sch, Coordinate{i, j})
				allAdjacentNumbers = append(allAdjacentNumbers, numbers...)
			}
		}
	}

	return sum(removeDuplicates(allAdjacentNumbers))
}
