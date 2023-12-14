package engine

import "strconv"

func IsSymbol(char string) bool {
	switch char {
	case "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "+", "-", "=", "/":
		return true
	default:
		return false
	}
}

func IsGear(char string) bool {
	switch char {
	case "*":
		return true
	default:
		return false
	}
}

func IsDigit(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}
