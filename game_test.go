package main

import (
	"fmt"
	"testing"
)

func checkScores(t *testing.T, e, a Score) {
	if e.BHoles != a.BHoles ||
		e.Burn != a.Burn ||
		e.CHoles != a.CHoles ||
		e.FHoles != a.FHoles ||
		e.HighY != a.HighY ||
		e.Step != a.Step {
		t.Fail()
		fmt.Printf("expect: %+v\n", e)
		fmt.Printf("actual: %+v\n", a)
	}
}

func Test_applySolidLines(t *testing.T) {
	//arrange
	game := Game{
		Strategy: strategy,
		Round:    19,
	}
	grid := Grid{
		{false, true},
		{true, false},
		{false, false},
		{false, false},
		{false, false},
	}
	field := grid.ToField()
	piece := Piece{FieldAfter: &field}

	//act
	game.applySolidLines(&piece)

	//assert
	field.Grid.assertNotEqualTo(piece.FieldAfter.Grid, t)
}

func Test_calculateMoves(t *testing.T) {
	//arrange
	game := Game{Strategy: strategy}
	grid := Grid{
		{false, false, false, false, true, true, true, true, true, true},
		{false, false, false, false, false, false, true, true, true, true},
		{false, false, false, false, false, false, false, false, true, true},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	game.MyPlayer = &Player{Field: grid.ToField()}
	game.asignUpdates("game", "this_piece_type", I)
	game.asignUpdates("game", "next_piece_type", Z)
	game.initPieces()
	expectedGrid := Grid{
		{false, false, false, false, true, true, true, true, true, true},
		{false, false, true, true, true, true, true, true, true, true},
		{false, false, false, false, false, false, false, false, true, true},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	expectedScore := Score{
		Burn:   0,
		BHoles: 1,
		FHoles: 1,
		CHoles: 0,
		HighY:  1,
		Step:   2,
	}

	//act
	result := game.calculateMoves()

	//assert
	result.FieldAfter.Grid.assertEqualTo(expectedGrid, t)
	checkScores(t, expectedScore, result.Score)
}

func Test_55d7935d35ec1d06d15c9d7e_31_T_Spin_Single(t *testing.T) {
	//arrange
	game := Game{Strategy: strategy}
	grid := Grid{
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	game.MyPlayer = &Player{Field: grid.ToField()}
	game.asignUpdates("game", "this_piece_type", Z)
	game.asignUpdates("game", "next_piece_type", T)
	game.initPieces()
	game.MyPlayer.Combo = 2

	expectedGrid := Grid{
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{false, false, false, false, true, true, false, false, false, false},
		{false, false, false, true, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	expectedScore := Score{
		Burn:   0,
		BHoles: 0,
		FHoles: 1,
		CHoles: 0,
		HighY:  4,
		Step:   4,
	}

	//act
	result := game.calculateMoves()

	//assert
	result.FieldAfter.Grid.assertEqualTo(expectedGrid, t)
	checkScores(t, expectedScore, result.Score)
}

func Test_55edfd6135ec1d06d15dad14_42_T_Spin_Double(t *testing.T) {
	//arrange
	game := Game{Strategy: strategy}
	grid := Grid{
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, true, true, true, true, false, true, true, true, true},
		{true, true, true, true, false, false, false, true, true, true},
		{true, false, true, true, true, false, false, false, true, true},
		{false, false, true, false, false, false, false, false, true, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	game.MyPlayer = &Player{Field: grid.ToField()}
	game.asignUpdates("game", "this_piece_type", T)
	game.asignUpdates("game", "next_piece_type", J)
	game.initPieces()
	expectedGrid := Grid{
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, false, true, true, true, false, false, false, true, true},
		{false, false, true, false, false, false, false, false, true, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	expectedScore := Score{
		Burn:   2,
		BHoles: -1,
		FHoles: -1,
		CHoles: 0,
		HighY:  7,
		Step:   0,
	}

	//act
	result := game.calculateMoves()

	//assert
	result.FieldAfter.Grid.assertEqualTo(expectedGrid, t)
	checkScores(t, expectedScore, result.Score)
}

func Test_55c29f6435ec1d070e2b66e9_40(t *testing.T) {
	//arrange
	game := Game{Strategy: strategy}
	grid := Grid{
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, true, true, true, true, true, false, true, false, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{false, true, true, false, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, false},
		{false, true, true, false, false, true, true, true, false, false},
		{false, true, true, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	game.MyPlayer = &Player{Field: grid.ToField()}
	game.asignUpdates("game", "this_piece_type", T)
	game.asignUpdates("game", "next_piece_type", J)
	game.initPieces()
	expectedGrid := Grid{
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, true, true, true, true, true, false, true, false, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{false, true, true, false, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, false, false, true, true, true, true, true},
		{false, true, true, false, false, false, false, false, false, true},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	expectedScore := Score{
		Burn:   0,
		BHoles: 0,
		FHoles: 0,
		CHoles: 0,
		HighY:  12,
		Step:   0,
	}

	//act
	result := game.calculateMoves()

	//assert
	result.FieldAfter.Grid.assertEqualTo(expectedGrid, t)
	checkScores(t, expectedScore, result.Score)
}

func Test_55dc7ff01c687b0946a742f3_67(t *testing.T) {
	//arrange
	game := Game{Strategy: strategy}
	grid := Grid{
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, false, false, true, true, true},
		{false, true, true, true, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	game.MyPlayer = &Player{Field: grid.ToField()}
	game.asignUpdates("game", "this_piece_type", T)
	game.asignUpdates("game", "next_piece_type", I)
	game.initPieces()
	game.MyPlayer.Combo = 1
	expectedGrid := Grid{
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{false, true, true, true, true, true, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	expectedScore := Score{
		Burn:   0,
		BHoles: 0,
		FHoles: 0,
		CHoles: 0,
		HighY:  8,
		Step:   3,
	}

	//act
	result := game.calculateMoves()

	//assert
	result.FieldAfter.Grid.assertEqualTo(expectedGrid, t)
	checkScores(t, expectedScore, result.Score)
}

func Test_55dc5b501c687b0946a741a2_35(t *testing.T) {
	//arrange
	game := Game{Strategy: strategy}
	grid := Grid{
		{true, false, true, true, true, true, true, true, false, false},
		{true, false, false, false, true, false, false, false, false, false},
		{true, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	game.MyPlayer = &Player{Field: grid.ToField()}
	game.asignUpdates("game", "this_piece_type", S)
	game.asignUpdates("game", "next_piece_type", J)
	game.initPieces()
	game.MyPlayer.Combo = 3
	expectedGrid := Grid{
		{true, false, true, true, true, true, true, true, true, false},
		{true, false, false, false, true, false, false, true, true, false},
		{true, false, false, false, true, false, false, true, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	expectedScore := Score{
		Burn:   0,
		BHoles: 0,
		FHoles: 0,
		CHoles: 0,
		HighY:  0,
		Step:   4,
	}

	//act
	result := game.calculateMoves()

	//assert
	result.FieldAfter.Grid.assertEqualTo(expectedGrid, t)
	checkScores(t, expectedScore, result.Score)
}

func Test_55ededec1c687b0946a7e6c6_08(t *testing.T) {
	//arrange
	game := Game{Strategy: strategy}
	grid := Grid{
		{true, true, true, true, true, true, false, false, false, false},
		{true, true, true, true, false, false, false, false, false, false},
		{true, true, true, true, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	game.MyPlayer = &Player{Field: grid.ToField()}
	game.asignUpdates("game", "this_piece_type", S)
	game.asignUpdates("game", "next_piece_type", S)
	game.initPieces()
	game.MyPlayer.Combo = 0
	expectedGrid := Grid{
		{true, true, true, true, true, true, true, false, false, false},
		{true, true, true, true, false, true, true, false, false, false},
		{true, true, true, true, false, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	expectedScore := Score{
		Burn:   0,
		BHoles: 0,
		FHoles: 0,
		CHoles: 0,
		HighY:  0,
		Step:   4,
	}

	//act
	result := game.calculateMoves()

	//assert
	result.FieldAfter.Grid.assertEqualTo(expectedGrid, t)
	checkScores(t, expectedScore, result.Score)
}

func Test_560b136035ec1d3214e473b8_74(t *testing.T) {
	//arrange
	game := Game{Strategy: strategy}
	grid := Grid{
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, true, true, false, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, false, true, true, true, true, true, true, true, true},
		{true, false, true, true, true, true, true, true, true, true},
		{false, false, true, true, true, true, true, true, false, false},
		{false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	game.MyPlayer = &Player{Field: grid.ToField()}
	game.asignUpdates("game", "this_piece_type", I)
	game.asignUpdates("game", "next_piece_type", L)
	game.initPieces()
	expectedGrid := Grid{
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, true, true, false, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{false, true, true, true, true, true, true, true, false, false},
		{false, true, false, true, true, true, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	expectedScore := Score{
		Burn:   2,
		BHoles: -1,
		FHoles: 0,
		CHoles: 3,
		HighY:  14,
		Step:   3,
	}

	//act
	result := game.calculateMoves()

	//assert
	result.FieldAfter.Grid.assertEqualTo(expectedGrid, t)
	checkScores(t, expectedScore, result.Score)
}

func Test_563bfe8035ec1d521be3ee36_2(t *testing.T) {
	//arrange
	game := Game{Strategy: strategy}
	grid := Grid{
		{false, false, false, false, true, true, false, false, false, false},
		{false, false, false, true, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	game.MyPlayer = &Player{Field: grid.ToField()}
	game.asignUpdates("game", "this_piece_type", J)
	game.asignUpdates("game", "next_piece_type", I)
	game.initPieces()
	expectedGrid := Grid{
		{false, true, true, true, true, true, false, false, false, false},
		{false, true, false, true, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	expectedScore := Score{
		Burn:   0,
		BHoles: 0,
		FHoles: -1,
		CHoles: 0,
		HighY:  0,
		Step:   2,
	}

	//act
	result := game.calculateMoves()

	//assert
	result.FieldAfter.Grid.assertEqualTo(expectedGrid, t)
	checkScores(t, expectedScore, result.Score)
}

func Test_01(t *testing.T) {
	//arrange
	game := Game{Strategy: strategy}
	grid := Grid{
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, false, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, false, true, true},
		{false, true, true, true, true, true, true, false, true, false},
		{false, true, false, true, true, true, false, false, false, false},
		{false, false, false, false, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	game.MyPlayer = &Player{Field: grid.ToField()}
	game.asignUpdates("game", "this_piece_type", L)
	game.asignUpdates("game", "next_piece_type", S)
	game.initPieces()
	expectedGrid := Grid{
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, false, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, false},
		{false, true, false, true, true, true, true, true, false, false},
		{false, false, false, false, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	expectedScore := Score{
		Burn:   0,
		BHoles: 0,
		FHoles: 0,
		CHoles: 0,
		HighY:  8,
		Step:   1,
	}

	//act
	result := game.calculateMoves()

	//assert
	result.FieldAfter.Grid.assertEqualTo(expectedGrid, t)
	checkScores(t, expectedScore, result.Score)
}
