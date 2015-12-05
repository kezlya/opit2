package main

import (
	"fmt"
	"testing"
)

var testField = testGrid.ToField()

func Test_FindPositions_I(t *testing.T) {
	//arrange
	piece := InitPiece("I", 3, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 15)
}

func Test_FindPositions_J(t *testing.T) {
	//arrange
	piece := InitPiece("J", 3, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 27)
}

func Test_FindPositions_L(t *testing.T) {
	//arrange
	piece := InitPiece("L", 3, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 25)
}

func Test_FindPositions_O(t *testing.T) {
	//arrange
	piece := InitPiece("O", 4, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 8)
}

func Test_FindPositions_S(t *testing.T) {
	//arrange
	piece := InitPiece("S", 3, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 15)
}

func Test_FindPositions_T(t *testing.T) {
	//arrange
	piece := InitPiece("T", 3, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 28)
}

func Test_FindPositions_T2(t *testing.T) {
	//arrange
	grid := Grid{
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, true, true, true, true, false, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{false, true, true, true, false, true, true, true, true, false},
		{false, true, true, false, false, false, false, true, false, false},
		{false, false, true, true, false, false, false, false, false, false},
		{false, false, false, true, false, false, false, true, false, false},
		{false, false, true, true, false, false, true, true, true, false},
		{false, false, true, true, false, false, false, false, true, false},
		{false, false, true, false, false, false, false, false, true, false},
		{false, false, true, true, false, false, false, false, true, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	field := grid.ToField()
	piece := InitPiece("T", 3, 19)

	//act
	positions := field.FindPositions(piece)

	//assert
	assertPositions(t, positions, 39)
}

func Test_FindPositions_Z(t *testing.T) {
	//arrange
	piece := InitPiece("Z", 3, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 13)
}

func Test_After(t *testing.T) {
	//arrange
	/*var arangeT = testField
	space := make(map[string]Cell, 4)
	space["m1"] = Cell{X: 0, Y: 13}
	space["m2"] = Cell{X: 1, Y: 13}
	space["m3"] = Cell{X: 2, Y: 13}
	space["t2"] = Cell{X: 1, Y: 14}
	piece := Piece{Name: "T", Space: space}

	//act

	expected := Field{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, true, true, true, false, true, true, true, true, false}, {false, true, true, false, false, false, false, true, false, false}, {false, false, true, true, false, false, false, true, false, false}, {true, true, true, true, false, false, false, true, false, false}, {false, true, true, true, false, false, true, true, true, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	result := arangeT.AfterHole(piece.Space)

	//assert
	checkResults(t, expected, result)*/
}

/*

func Test_Picks(t *testing.T) {
	arrange := Field{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expect := Picks{12, 13, 14, 14, 14, 17, 14, 14, 13, 14}

	result := arrange.Picks()

	if !PicksIsEqual(expect, result) {
		fmt.Println(result)
		fmt.Println(expect)
		t.Fail()
	}
}
*/

/*
func Test_IsValid(t *testing.T) {
	arrange := Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, true, true, true, false, true, true, true, true, false}, {false, true, true, false, false, false, false, true, false, false}, {false, false, true, true, false, false, false, true, false, false}, {false, false, false, true, false, false, false, true, false, false}, {false, false, true, true, false, false, true, true, true, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	good1 := Cell{X: 2, Y: 6}
	good2 := Cell{X: 5, Y: 17}
	good3 := Cell{X: 0, Y: 17}
	good4 := Cell{X: 1, Y: 0}
	bad1 := Cell{X: 3, Y: 3}
	bad2 := Cell{X: -1, Y: 3}
	bad3 := Cell{X: 3, Y: -1}
	bad4 := Cell{X: 2, Y: 18}

	if !arrange.IsValid(&map[string]Cell{"x": good1}) {
		fmt.Println("should be valid", good1)
		t.Fail()
	}

	if !arrange.IsValid(&map[string]Cell{"x": good2}) {
		fmt.Println("should be valid", good2)
		t.Fail()
	}

	if !arrange.IsValid(&map[string]Cell{"x": good3}) {
		fmt.Println("should be valid", good3)
		t.Fail()
	}

	if !arrange.IsValid(&map[string]Cell{"x": good4}) {
		fmt.Println("should be valid", good4)
		t.Fail()
	}

	if arrange.IsValid(&map[string]Cell{"x": bad1}) {
		fmt.Println("should not be valid", bad1)
		t.Fail()
	}

	if arrange.IsValid(&map[string]Cell{"x": bad2}) {
		fmt.Println("should not be valid", bad2)
		t.Fail()
	}

	if arrange.IsValid(&map[string]Cell{"x": bad3}) {
		fmt.Println("should not be valid", bad3)
		t.Fail()
	}

	if arrange.IsValid(&map[string]Cell{"x": bad4}) {
		fmt.Println("should not be valid", bad4)
		t.Fail()
	}

	if arrange.IsValid(&map[string]Cell{"1": good1, "2": good2, "3": good3, "4": good4, "5": bad1, "6": bad2, "7": bad3, "8": bad4}) {
		fmt.Println("should not be valid")
		t.Fail()
	}
}

*/

func assertPositions(t *testing.T, positions []Piece, expectedCount int) {
	if len(positions) != expectedCount {
		t.Fail()
		fmt.Println(len(positions), "positions found, should be", expectedCount)
	}
	//testField.Grid.visual()
	for _, p := range positions {
		if p.FieldAfter == nil {
			t.Fail()
			fmt.Println("FieldAfter is nil", p)
		}
		//p.FieldAfter.Grid.visual()
		//fmt.Println(p.Key)
	}
}
