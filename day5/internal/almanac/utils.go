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
