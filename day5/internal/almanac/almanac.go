package almanac

import (
	"regexp"
)

type Map struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

type Almanac struct {
	seeds                  []int
	seedToSoilMaps         []Map
	soilToFertilizers      []Map
	fertilizerToWaters     []Map
	waterToLights          []Map
	lightToTemperatures    []Map
	temperatureToHumiditys []Map
	humidityToLocations    []Map
}

func NewAlmanac(lines []string) Almanac {
	seedsRegex, _ = regexp.Compile("")

	return Almanac{}
}
