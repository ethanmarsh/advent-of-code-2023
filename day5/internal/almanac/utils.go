package almanac

func IsEqual(lhs []int, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for index := range lhs {
		if lhs[index] != rhs[index] {
			return false
		}
	}
	return true
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetMinimum(values []int) int {
	var minimum = values[0]
	for _, value := range values {
		if value < minimum {
			minimum = value
		}
	}
	return minimum
}
