package main

import (
	"fmt"
	"testing"
)

var testField = Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, false, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, false, true, true, false, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
var strategy = Strategy{1, 1, 1, 1}

func Test_availablePositions_I(t *testing.T) {
	piece := "I"
	expectPositions := 15

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}

func Test_availablePositions_J(t *testing.T) {
	piece := "J"
	expectPositions := 24

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}

func Test_availablePositions_L(t *testing.T) {
	piece := "L"
	expectPositions := 24

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}

func Test_availablePositions_O(t *testing.T) {
	piece := "O"
	expectPositions := 7

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}

func Test_availablePositions_S(t *testing.T) {
	piece := "S"
	expectPositions := 12

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}

func Test_availablePositions_T(t *testing.T) {
	piece := "T"
	expectPositions := 24

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}

func Test_availablePositions_Z(t *testing.T) {
	piece := "Z"
	expectPositions := 12

	result := testField.Positions(piece, strategy)

	if len(result) != expectPositions {
		t.Fail()
		fmt.Println(expectPositions, "!=", len(result))
	}
}
