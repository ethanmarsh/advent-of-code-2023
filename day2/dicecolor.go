package main

type DiceColor int

const (
	Red DiceColor = iota
	Green
	Blue
)

func (color DiceColor) String() string {
	return [...]string{"red", "green", "blue"}[color]
}

func (color DiceColor) EnumIndex() int {
	return int(color)
}

func NewDiceColor(str string) DiceColor {
	switch str {
	case "red":
		return Red
	case "green":
		return Green
	case "blue":
		return Blue
	default:
		panic("Error creating DiceColor")
	}
}
