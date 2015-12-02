package main

import (
	"fmt"
	"testing"
)

var testField = Field{
	Grid: [][]bool{
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
	},
}
var tfPicks = testField.Picks()

func Test_availablePositions_I(t *testing.T) {
	piece := Piece{Name: "I"}
	piece.InitSpace(Cell{X: 3, Y: 19})
	expectVp := 15
	expectedFh := 0

	resultVp := testField.ValidPosition(piece)
	_, hFixable := testField.FindHoles()
	resultFh := testField.FixHoles(piece, hFixable)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		PrintVisual(testField)
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		PrintVisual(testField)
	}
}

func Test_availablePositions_J(t *testing.T) {
	piece := Piece{Name: "J"}
	piece.InitSpace(Cell{X: 3, Y: 19})
	expectVp := 24
	expectedFh := 3

	resultVp := testField.ValidPosition(piece)
	_, hFixable := testField.FindHoles()
	resultFh := testField.FixHoles(piece, hFixable)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		PrintVisual(testField)
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		PrintVisual(testField)
	}
}

func Test_availablePositions_L(t *testing.T) {
	piece := Piece{Name: "L"}
	piece.InitSpace(Cell{X: 3, Y: 19})
	expectVp := 23
	expectedFh := 2

	resultVp := testField.ValidPosition(piece)
	_, hFixable := testField.FindHoles()
	resultFh := testField.FixHoles(piece, hFixable)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		PrintVisual(testField)
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		PrintVisual(testField)
	}
}

func Test_availablePositions_O(t *testing.T) {
	piece := Piece{Name: "O"}
	piece.InitSpace(Cell{X: 4, Y: 19})
	expectVp := 7
	expectedFh := 1

	resultVp := testField.ValidPosition(piece)
	_, hFixable := testField.FindHoles()
	resultFh := testField.FixHoles(piece, hFixable)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		PrintVisual(testField)
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		PrintVisual(testField)
	}
}

func Test_availablePositions_S(t *testing.T) {
	piece := Piece{Name: "S"}
	piece.InitSpace(Cell{X: 3, Y: 19})
	expectVp := 12
	expectedFh := 3

	resultVp := testField.ValidPosition(piece)
	_, hFixable := testField.FindHoles()
	resultFh := testField.FixHoles(piece, hFixable)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		PrintVisual(testField)
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		PrintVisual(testField)
	}
}

func Test_availablePositions_T(t *testing.T) {
	piece := Piece{Name: "T"}
	piece.InitSpace(Cell{X: 3, Y: 19})
	expectVp := 24
	expectedFh := 4

	resultVp := testField.ValidPosition(piece)
	_, hFixable := testField.FindHoles()
	resultFh := testField.FixHoles(piece, hFixable)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		PrintVisual(testField)
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		PrintVisual(testField)
	}
}

func Test_availablePositions_Z(t *testing.T) {
	piece := Piece{Name: "Z"}
	piece.InitSpace(Cell{X: 3, Y: 19})
	expectVp := 12
	expectedFh := 1

	resultVp := testField.ValidPosition(piece)
	_, hFixable := testField.FindHoles()
	resultFh := testField.FixHoles(piece, hFixable)

	if len(resultVp) != expectVp {
		t.Fail()
		fmt.Println(expectVp, "!=", len(resultVp))
		PrintVisual(testField)
	}
	if len(resultFh) != expectedFh {
		t.Fail()
		fmt.Println(expectedFh, "!=", len(resultFh))
		PrintVisual(testField)
	}
}
