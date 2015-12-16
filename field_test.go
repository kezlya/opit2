package main

import (
	"fmt"
	"testing"
)

var testField = testGrid.ToField()

func Test_FindPositions_I(t *testing.T) {
	//arrange
	piece := InitPiece(I, 3, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 15)
}

func Test_FindPositions_J(t *testing.T) {
	//arrange
	piece := InitPiece(J, 3, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 27)
}

func Test_FindPositions_L(t *testing.T) {
	//arrange
	piece := InitPiece(L, 3, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 25)
}

func Test_FindPositions_O(t *testing.T) {
	//arrange
	piece := InitPiece(O, 4, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 8)
}

func Test_FindPositions_S(t *testing.T) {
	//arrange
	piece := InitPiece(S, 3, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 15)
}

func Test_FindPositions_T(t *testing.T) {
	//arrange
	piece := InitPiece(T, 3, 19)

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
	piece := InitPiece(T, 3, 19)

	//act
	positions := field.FindPositions(piece)

	//assert
	assertPositions(t, positions, 39)
}

func Test_FindPositions_Z(t *testing.T) {
	//arrange
	piece := InitPiece(Z, 3, 19)

	//act
	positions := testField.FindPositions(piece)

	//assert
	assertPositions(t, positions, 13)
}

func assertPositions(t *testing.T, positions []*Piece, exCount int) {
	if len(positions) != exCount {
		t.Fail()
		fmt.Println(len(positions), "positions found, should be", exCount)
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
