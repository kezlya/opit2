package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Copy(t *testing.T) {
	//arrange
	a := Field{
		Width:   2,
		Height:  2,
		Empty:   1,
		MaxPick: 1,
		Grid:    [][]bool{{false, true}, {false, false}},
		Picks:   []int{0, 1},
	}

	//act
	b := a.Copy()
	a.Width = 1

	//assert
	if b.Width != 2 || b.Height != 2 || b.Empty != 1 || b.MaxPick != 1 {
		t.Fail()
		fmt.Println("Properties of the Field was not copied")
	}
	ag := reflect.ValueOf(a.Grid).Pointer()
	bg := reflect.ValueOf(b.Grid).Pointer()
	if !a.Grid.isEqual(b.Grid) || ag == bg {
		t.Fail()
		fmt.Println("Grid of the Field was not copied")
		fmt.Println("Grid pointers", ag, "and", bg, "should be different")
	}
	ap := reflect.ValueOf(a.Picks).Pointer()
	bp := reflect.ValueOf(b.Picks).Pointer()
	if !a.Picks.isEqual(b.Picks) || ap == bp {
		t.Fail()
		fmt.Println("Picks of the Field was not copied")
		fmt.Println("Picks pointers", ap, "and", bp, "should be different")
	}
}

func Test_FindPositions(t *testing.T) {
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
	positions, countSearchCalls, bagLen := field.FindPositions(piece)

	//assert
	if len(positions) != 39 {
		fmt.Println("Count of positions is wrong:", len(positions))
		grid.visual()
		t.Fail()
		for _, p := range positions {
			p.FieldAfter.Grid.visual()
			fmt.Println(p.Key)
		}
	}
	if countSearchCalls != 575 || bagLen != 231 {
		fmt.Println("countSearchCalls", countSearchCalls)
		fmt.Println("bagLen", bagLen)
		t.Fail()
	}
}

/*
func Test_Burn(t *testing.T) {
	arange := Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, false, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, false, true, true, false, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expect := Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, false, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, false, true, true, false, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	arange.Burn()

	if !FieldIsEqual(arange, expect) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], arange[y-i])
		}
	}
}

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



func Test_FixHoles_Z(t *testing.T) {
	var arangePathField = Field{{false, false, false, false, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, true, true}, {false, false, false, false, false, false, false, false, true, true}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	piece := Piece{Name: "Z", Rotation: 0}
	piece.InitSpace(Cell{X: 3, Y: 19})
	hole := Cell{X: 2, Y: 0}

	good_positions := arangePathField.FixHoles(piece, []Cell{hole})

	if len(good_positions) != 1 {
		for _, pos := range good_positions {
			fmt.Println("good positions failed")
			fmt.Println(pos.Moves)
			PrintVisual(arangePathField)
		}
		t.Fail()
	}
}

func Test_FixHoles_J(t *testing.T) {
	var arangePathField = Field{{false, false, false, false, true, true, false, false, false, false}, {false, false, false, true, true, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	piece := Piece{Name: "J", Rotation: 0}
	piece.InitSpace(Cell{X: 3, Y: 19})
	hole := Cell{X: 3, Y: 0}

	good_positions := arangePathField.FixHoles(piece, []Cell{hole})

	if len(good_positions) != 1 {
		for _, pos := range good_positions {
			fmt.Println("good positions failed")
			fmt.Println(pos.Moves)
			PrintVisual(arangePathField)
		}
		t.Fail()
	}
}

func Test_ValidatePosition_I(t *testing.T) {
	var arangeField = Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, true, true, true, false, true, true, true, true, false}, {false, true, true, false, false, false, false, true, false, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, false, true, false, false, false, true, false, false}, {false, false, true, true, false, false, true, true, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	piece := Piece{Name: "I", Rotation: 0}
	piece.InitSpace(Cell{X: 3, Y: 19})

	validPieces := arangeField.ValidPosition(piece)

	if len(validPieces) != 13 {
		PrintVisual(arangeField)
		fmt.Println("should be 13 but returned", len(validPieces))
		t.Fail()
	}
}

func Test_ValidatePosition_T(t *testing.T) {
	var arangeField = Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, true, true, true, false, true, true, true, true, false}, {false, true, true, false, false, false, false, true, false, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, false, true, false, false, false, true, false, false}, {false, false, true, true, false, false, true, true, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	piece := Piece{Name: "T", Rotation: 0}
	piece.InitSpace(Cell{X: 3, Y: 19})

	validPieces := arangeField.ValidPosition(piece)

	if len(validPieces) != 23 {
		PrintVisual(arangeField)
		fmt.Println("should be 23 but returned", len(validPieces))
		t.Fail()
	}
}
*/
