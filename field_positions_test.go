package main

import (
	"fmt"
	"testing"
)

var testField = Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, false, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, false, true, true, false, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, false, false, true, true, true, true, true}, {true, true, true, false, false, false, true, true, true, true}, {false, true, true, false, false, true, true, true, true, true}, {false, true, true, false, false, false, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
var strategy = Strategy{1, 1, 1, 1}

func Test_availablePositions_I(t *testing.T) {
	piece := Piece{Name: "I"}
	piece.InitSpace(Cell{X: 3, Y: 19})
	expectPositions := 15

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
		PrintVisual(testField)
	}
}

func Test_availablePositions_J(t *testing.T) {
	piece := Piece{Name: "J"}
	piece.InitSpace(Cell{X: 3, Y: 19})
	expectPositions := 24

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
		PrintVisual(testField)
	}
}

func Test_availablePositions_L(t *testing.T) {
	piece := Piece{Name: "L"}
	piece.InitSpace(Cell{X: 3, Y: 19})
	expectPositions := 24

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
		PrintVisual(testField)
	}
}

func Test_availablePositions_O(t *testing.T) {
	piece := Piece{Name: "O"}
	piece.InitSpace(Cell{X: 4, Y: 19})
	expectPositions := 7

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
		PrintVisual(testField)
	}
}

func Test_availablePositions_S(t *testing.T) {
	piece := Piece{Name: "S"}
	piece.InitSpace(Cell{X: 3, Y: 19})
	expectPositions := 12

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
		PrintVisual(testField)
	}
}

func Test_availablePositions_T(t *testing.T) {
	piece := Piece{Name: "T"}
	piece.InitSpace(Cell{X: 3, Y: 19})
	expectPositions := 27

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
		PrintVisual(testField)
	}
}

func Test_availablePositions_Z(t *testing.T) {
	piece := Piece{Name: "Z"}
	piece.InitSpace(Cell{X: 3, Y: 19})
	expectPositions := 12

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
		PrintVisual(testField)
	}
}

func PrintVisual(f Field) {
	y := len(f) - 1
	for i := range f {
		for _, c := range f[y-i] {
			if c {
				fmt.Print("⬛ ")
			} else {
				fmt.Print("⬜ ")
			}
		}
		fmt.Println()
	}
}
