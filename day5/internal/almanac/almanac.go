package almanac

type Almanac struct {
	seeds                  []int
	seedToSoil             []Map
	soilToFertilizers      []Map
	fertilizerToWaters     []Map
	waterToLights          []Map
	lightToTemperatures    []Map
	temperatureToHumiditys []Map
	humidityToLocations    []Map
}

func (lhs Almanac) Equals(rhs Almanac) bool {
	if !IsEqual(lhs.seeds, rhs.seeds) {
		return false
	}

	if !Equals(lhs.seedToSoil, rhs.seedToSoil) {
		return false
	}

	if !Equals(lhs.soilToFertilizers, rhs.soilToFertilizers) {
		return false
	}

	if !Equals(lhs.fertilizerToWaters, rhs.fertilizerToWaters) {
		return false
	}

	if !Equals(lhs.waterToLights, rhs.waterToLights) {
		return false
	}

	if !Equals(lhs.lightToTemperatures, rhs.lightToTemperatures) {
		return false
	}

	if !Equals(lhs.temperatureToHumiditys, rhs.temperatureToHumiditys) {
		return false
	}

	if !Equals(lhs.humidityToLocations, rhs.humidityToLocations) {
		return false
	}

	return true
}

func NewAlmanac(lines []string) Almanac {
	seeds, err := ParseSeeds(lines[0])
	CheckError(err)

	var almanac Almanac
	almanac.seeds = seeds

	for index, line := range lines {
		if line == "seed-to-soil map:" {
			almanac.seedToSoil = ParseMap(lines, index+1)
		} else if line == "soil-to-fertilizer map:" {
			almanac.soilToFertilizers = ParseMap(lines, index+1)
		} else if line == "fertilizer-to-water map:" {
			almanac.fertilizerToWaters = ParseMap(lines, index+1)
		} else if line == "water-to-light map:" {
			almanac.waterToLights = ParseMap(lines, index+1)
		} else if line == "light-to-temperature map:" {
			almanac.lightToTemperatures = ParseMap(lines, index+1)
		} else if line == "temperature-to-humidity map:" {
			almanac.temperatureToHumiditys = ParseMap(lines, index+1)
		} else if line == "humidity-to-location map:" {
			almanac.humidityToLocations = ParseMap(lines, index+1)
		}
	}

	return almanac
}
