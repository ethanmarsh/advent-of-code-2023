package almanac

import "testing"

func TestIntSliceIsEqual(t *testing.T) {
	lhs := []int{79, 14, 55, 13}
	rhs := []int{79, 14, 55, 13}

	if !IsEqual(lhs, rhs) {
		t.Error("Int slice IsEqual not working")
	}
}
