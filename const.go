package main

type boardPosition string

func (b boardPosition) str() string {
	return string(b)
}

func (b boardPosition) parse() (int, int) {
	x, y := 0, 0
	switch string(b[0]) {
	case "a":
		x = 0
	case "b":
		x = 1
	case "c":
		x = 2
	}

	switch string(b[1]) {
	case "1":
		y = 0
	case "2":
		y = 1
	case "3":
		y = 2
	}

	return x, y
}

const (
	boardPositionA1 boardPosition = "a1"
	boardPositionA2 boardPosition = "a2"
	boardPositionA3 boardPosition = "a3"
	boardPositionB1 boardPosition = "b1"
	boardPositionB2 boardPosition = "b2"
	boardPositionB3 boardPosition = "b3"
	boardPositionC1 boardPosition = "c1"
	boardPositionC2 boardPosition = "c2"
	boardPositionC3 boardPosition = "c3"
)

var validPositions = map[boardPosition]bool{
	boardPositionA1: true,
	boardPositionA2: true,
	boardPositionA3: true,
	boardPositionB1: true,
	boardPositionB2: true,
	boardPositionB3: true,
	boardPositionC1: true,
	boardPositionC2: true,
	boardPositionC3: true,
}
