package main

import (
	"fmt"
	"testing"
)

func testField() [][]bool{
	return [][]bool {{true, false, true, true, true, true, true, true, true, true}, {true, true, false, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, false, true, true, false, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
}

func Test_availablePositions_I(t *testing.T) {
	piece := "I"
	expectPositions := 15

	result := _availablePositions(piece, testField())

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}
/*
func Test_availablePositions_J(t *testing.T) {
	piece = "J"
	expectPositions := 24

	result := _availablePositions(piece, testField)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}

func Test_availablePositions_L(t *testing.T) {
	_setup2()
	CurrentPiece = "L"
	expectPositions := 24

	result := _availablePositions(CurrentPiece, MyPlayer.Field)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}

func Test_availablePositions_O(t *testing.T) {
	_setup2()
	CurrentPiece = "O"
	expectPositions := 7

	result := _availablePositions(CurrentPiece, MyPlayer.Field)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}

func Test_availablePositions_S(t *testing.T) {
	_setup2()
	CurrentPiece = "S"
	expectPositions := 12

	result := _availablePositions(CurrentPiece, MyPlayer.Field)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}

func Test_availablePositions_T(t *testing.T) {
	_setup2()
	CurrentPiece = "T"
	expectPositions := 25

	result := _availablePositions(CurrentPiece, MyPlayer.Field)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}

func Test_availablePositions_Z(t *testing.T) {
	_setup2()
	CurrentPiece = "Z"
	expectPositions := 12

	result := _availablePositions(CurrentPiece, MyPlayer.Field)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}
*/