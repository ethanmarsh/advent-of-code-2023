package almanac

type Almanac struct {
	seeds                  []int
	seedToSoil             MapSet
	soilToFertilizers      MapSet
	fertilizerToWaters     MapSet
	waterToLights          MapSet
	lightToTemperatures    MapSet
	temperatureToHumiditys MapSet
	humidityToLocations    MapSet
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

func (almanac Almanac) GetSeedLocations() []int {
	var locations []int
	for _, seed := range almanac.seeds {
		soil := almanac.seedToSoil.GetMappedValue(seed)
		fertilizer := almanac.soilToFertilizers.GetMappedValue(soil)
		water := almanac.fertilizerToWaters.GetMappedValue(fertilizer)
		light := almanac.waterToLights.GetMappedValue(water)
		temp := almanac.lightToTemperatures.GetMappedValue(light)
		humidity := almanac.temperatureToHumiditys.GetMappedValue(temp)
		location := almanac.humidityToLocations.GetMappedValue(humidity)
		locations = append(locations, location)
	}
	return locations
}

func (almanac Almanac) GetLowestSeedLocation() int {
	return GetMinimum(almanac.GetSeedLocations())
}
