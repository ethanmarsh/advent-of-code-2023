package engine

import (
	"errors"
	"fmt"
)

type NumberIndex struct {
	number     string
	startIndex int
	endIndex   int
}

type schematic = [][]string

type Coordinate struct {
	x int
	y int
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
		return NumberIndex{"", coord.x + 1, 0}
	}

	result := findNumberBeginning(sch, Coordinate{coord.x, coord.y - 1}).combine(char)
	// fmt.Println("Find Number Beginning - " + result)
	return result
}

func findNumberEnd(sch schematic, coord Coordinate) NumberIndex {
	// fmt.Printf("Find Number End - Coord:%d,%d\n", coord.x, coord.y)
	char, err := charAt(sch, coord)
	if err != nil || !IsDigit(char) {
		return NumberIndex{"", 0, coord.x - 1}
	}

	result := combine(char, findNumberEnd(sch, Coordinate{coord.x, coord.y + 1}))
	// fmt.Println("Find Number End - " + result)
	return result
}

func findFullNumber(sch schematic, coord Coordinate) NumberIndex {
	// fmt.Printf("Find Full Number - Coord:%d,%d\n", coord.x, coord.y)
	beginning := findNumberBeginning(sch, Coordinate{coord.x, coord.y - 1})
	end := findNumberEnd(sch, Coordinate{coord.x, coord.y + 1})
	combined := beginning.number + char + end.number
	combinedNumberIndex := NumberIndex{combined, beginning.startIndex, end.endIndex}
	// fmt.Println("Combined: " + combined)
	return combinedNumberIndex
}

func findAdjacentDigit(sch schematic, coord Coordinate) []int {
	fmt.Printf("Find Adjacent Digit - Coord:%d,%d\n", coord.x, coord.y)

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
				fmt.Printf("Found full number %s\n", fullNumberStr)
				adjacentNumbers = append(adjacentNumbers, fullNumberStr)
			}
		}
	}
	return adjacentNumbers
}

func sum(array []int) int {
	var sum int
	for _, num := range array {
		sum += num
	}
	return sum
}

func NewEngine(sch schematic) int {
	var alladjacentNumbers []int
	for i, line := range sch {
		for j, character := range line {
			if IsSymbol(character) {
				fmt.Printf("Found symbol %s\n", character)
				numbers := findAdjacentDigit(sch, Coordinate{i, j})
				alladjacentNumbers = append(alladjacentNumbers, numbers...)
			}
		}
	}
	return sum(alladjacentNumbers)
}
