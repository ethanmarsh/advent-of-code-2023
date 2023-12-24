package almanac

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Map struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func (daMap Map) String() string {
	return fmt.Sprintf("Dest: %d, Source: %d, Length: %d\n", daMap.destinationRangeStart, daMap.sourceRangeStart, daMap.rangeLength)
}

func Equals(lhs []Map, rhs []Map) bool {
	if len(lhs) != len(rhs) {
		fmt.Println("Lengths of maps don't match")
		return false
	}

	for index := range lhs {
		if lhs[index] != rhs[index] {
			fmt.Println("LHS map doesn't match RHS map")
			return false
		}
	}

	return true
}

func ParseSeeds(line string) ([]int, error) {
	seedsString, found := strings.CutPrefix(line, "seeds: ")
	if !found {
		return []int{}, errors.New("didn't find seeds string")
	}

	fields := strings.Fields(seedsString)
	var convertedFields []int
	for _, field := range fields {
		convert, err := strconv.Atoi(field)
		CheckError(err)
		convertedFields = append(convertedFields, convert)
	}

	return convertedFields, nil
}

func ParseMap(lines []string, startingIndex int) []Map {
	var daMap []Map
	for i := startingIndex; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}

		fields := strings.Fields(lines[i])
		if len(fields) != 3 {
			panic("error parsing map - should have 3 numbers per line")
		}

		destinationRangeStart, err0 := strconv.Atoi(fields[0])
		CheckError(err0)
		sourceRangeStart, err1 := strconv.Atoi(fields[1])
		CheckError(err1)
		rangeLength, err2 := strconv.Atoi(fields[2])
		CheckError(err2)
		daMap = append(daMap, Map{destinationRangeStart, sourceRangeStart, rangeLength})
	}
	return daMap
}
