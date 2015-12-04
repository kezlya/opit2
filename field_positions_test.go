package main

import (
	"fmt"
	"testing"
)

var testGrid = Grid{
	{true, false, true, true, true, true, true, true, true, true},
	{true, true, false, true, true, false, true, true, true, true},
	{true, true, true, true, true, true, true, false, true, true},
	{true, true, true, true, true, true, true, false, true, true},
	{true, true, true, true, true, true, true, true, true, false},
	{true, true, true, true, false, true, true, false, true, true},
	{false, true, true, true, true, true, true, true, true, true},
	{true, true, true, true, false, true, true, true, true, true},
	{true, true, true, true, true, true, true, true, true, false},
	{true, true, true, true, true, true, true, false, true, true},
	{true, true, false, false, false, false, true, true, true, true},
	{true, true, true, false, false, false, true, true, true, true},
	{false, true, true, false, false, true, true, true, true, true},
	{false, true, true, false, false, false, true, true, false, true},
	{false, false, true, false, false, false, true, false, false, false},
	{false, false, true, false, false, false, true, false, false, false},
	{false, false, true, false, false, false, true, false, false, false},
	{false, false, true, false, false, false, true, false, false, false},
	{false, false, true, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
}
var testField = testGrid.ToField()

var tfPicks = testField.Picks

func Test_availablePositions_I(t *testing.T) {
	piece := InitPiece("I", 3, 19)
	expectVp := 15
	expectedFh := 0

	resultVp := testField.ValidPosition(piece)
	resultFh := testField.FixHoles(piece)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		testField.Grid.visual()
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		testField.Grid.visual()
	}
}

func Test_availablePositions_J(t *testing.T) {
	piece := InitPiece("J", 3, 19)
	expectVp := 24
	expectedFh := 3

	resultVp := testField.ValidPosition(piece)
	resultFh := testField.FixHoles(piece)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		testField.Grid.visual()
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		testField.Grid.visual()
	}
}

func Test_availablePositions_L(t *testing.T) {
	piece := InitPiece("L", 3, 19)
	expectVp := 23
	expectedFh := 2

	resultVp := testField.ValidPosition(piece)
	resultFh := testField.FixHoles(piece)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		testField.Grid.visual()
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		testField.Grid.visual()
	}
}

func Test_availablePositions_O(t *testing.T) {
	piece := InitPiece("O", 4, 19)
	expectVp := 7
	expectedFh := 1

	resultVp := testField.ValidPosition(piece)
	resultFh := testField.FixHoles(piece)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		testField.Grid.visual()
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		testField.Grid.visual()
	}
}

func Test_availablePositions_S(t *testing.T) {
	piece := InitPiece("S", 3, 19)
	expectVp := 12
	expectedFh := 3

	resultVp := testField.ValidPosition(piece)
	resultFh := testField.FixHoles(piece)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		testField.Grid.visual()
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		testField.Grid.visual()
	}
}

func Test_availablePositions_T(t *testing.T) {
	piece := InitPiece("T", 3, 19)
	expectVp := 24
	expectedFh := 4

	resultVp := testField.ValidPosition(piece)
	resultFh := testField.FixHoles(piece)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		testField.Grid.visual()
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		testField.Grid.visual()
	}
}

func Test_availablePositions_Z(t *testing.T) {
	piece := InitPiece("Z", 3, 19)
	expectVp := 12
	expectedFh := 1

	resultVp := testField.ValidPosition(piece)
	resultFh := testField.FixHoles(piece)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		testField.Grid.visual()
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		testField.Grid.visual()
	}
}
