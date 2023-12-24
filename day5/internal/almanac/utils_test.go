package almanac

import "testing"

func TestIntSliceIsEqual(t *testing.T) {
	lhs := []int{79, 14, 55, 13}
	rhs := []int{79, 14, 55, 13}

	if !IsEqual(lhs, rhs) {
		t.Error("Int slice IsEqual not working")
	}
}

func TestMinimum(t *testing.T) {
	values := []int{82, 43, 86, 35}
	expectedLowest := 35
	actualLowest := GetMinimum(values)
	if expectedLowest != actualLowest {
		t.Error("Expected lowest int does not match actual lowest int")
	}
}
