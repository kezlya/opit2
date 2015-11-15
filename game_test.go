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

func Test_calculateMoves(t *testing.T) {
	game := Game{Strategy: strategy}
	field := Field{{false, false, false, false, true, true, true, true, true, true}, {false, false, false, false, false, false, true, true, true, true}, {false, false, false, false, false, false, false, false, true, true}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "I", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "Z", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Empty = 18
	expectedField := Field{{false, false, false, false, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, true, true}, {false, false, false, false, false, false, false, false, true, true}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expectedScore := Score{Burn: 0, BHoles: 1, FHoles: 1, CHoles: 0, HighY: 1, Step: 2}

	pos := game.calculateMoves()

	checkResults(t, expectedField, pos.FieldAfter)
	checkScores(t, expectedScore, pos.Score)
}

func Test_55c29f6435ec1d070e2b66e9_40(t *testing.T) {
	game := Game{Strategy: strategy}
	field := Field{{true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, false, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {false, true, true, false, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, false}, {false, true, true, false, false, true, true, true, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "T", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "J", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Empty = 5
	expectedField := Field{{true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, false, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {false, true, true, false, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, false, false, true, true, true, true, true}, {false, true, true, false, false, false, false, false, false, true}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expectedScore := Score{Burn: 0, BHoles: 0, FHoles: 0, CHoles: 0, HighY: 12, Step: 0}

	pos := game.calculateMoves()

	checkResults(t, expectedField, pos.FieldAfter)
	checkScores(t, expectedScore, pos.Score)
}

func Test_55d7935d35ec1d06d15c9d7e_31_T_Spin_Single(t *testing.T) {
	game := Game{Strategy: strategy}
	field := Field{{true, true, true, false, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "Z", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "T", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 2
	game.MyPlayer.Empty = 15

	expectedField := Field{{true, true, true, false, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {false, false, false, false, true, true, false, false, false, false}, {false, false, false, true, true, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expectedField.Burn()
	expectedScore := Score{Burn: 0, BHoles: 0, FHoles: 1, CHoles: 0, HighY: 4, Step: 4}

	pos := game.calculateMoves()

	checkResults(t, expectedField, pos.FieldAfter)
	checkScores(t, expectedScore, pos.Score)
}

func Test_55dc7ff01c687b0946a742f3_67(t *testing.T) {
	game := Game{Strategy: strategy}
	field := Field{{true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, false, false, true, true, true}, {false, true, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "T", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "I", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 1
	expectedField := Field{{true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, false, true, true, true}, {false, true, true, true, true, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expectedScore := Score{Burn: 0, BHoles: 0, FHoles: 0, CHoles: 0, HighY: 8, Step: 3}

	pos := game.calculateMoves()

	checkResults(t, expectedField, pos.FieldAfter)
	checkScores(t, expectedScore, pos.Score)
}

func Test_55dc5b501c687b0946a741a2_35(t *testing.T) {
	game := Game{Strategy: strategy}
	field := Field{{true, false, true, true, true, true, true, true, false, false}, {true, false, false, false, true, false, false, false, false, false}, {true, false, false, false, true, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "S", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "J", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 3
	game.MyPlayer.Empty = 12
	expectedField := Field{{true, false, true, true, true, true, true, true, true, false}, {true, false, false, false, true, false, false, true, true, false}, {true, false, false, false, true, false, false, true, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expectedScore := Score{Burn: 0, BHoles: 0, FHoles: 0, CHoles: 0, HighY: 0, Step: 4}

	pos := game.calculateMoves()

	checkResults(t, expectedField, pos.FieldAfter)
	checkScores(t, expectedScore, pos.Score)
}

func Test_55ededec1c687b0946a7e6c6_08(t *testing.T) {
	game := Game{Strategy: strategy}
	field := Field{{true, true, true, true, true, true, false, false, false, false}, {true, true, true, true, false, false, false, false, false, false}, {true, true, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "S", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "S", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 0
	game.MyPlayer.Empty = 12
	expectedField := Field{{true, true, true, true, true, true, true, false, false, false}, {true, true, true, true, false, true, true, false, false, false}, {true, true, true, true, false, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expectedScore := Score{Burn: 0, BHoles: 0, FHoles: 0, CHoles: 0, HighY: 0, Step: 4}

	pos := game.calculateMoves()

	checkResults(t, expectedField, pos.FieldAfter)
	checkScores(t, expectedScore, pos.Score)
}

func Test_55edfd6135ec1d06d15dad14_42_T_Spin_Double(t *testing.T) {
	game := Game{Strategy: strategy}
	field := Field{{true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {true, true, true, true, false, false, false, true, true, true}, {true, false, true, true, true, false, false, false, true, true}, {false, false, true, false, false, false, false, false, true, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "T", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "J", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 0
	expectedField := Field{{true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, false, true, true, true, false, false, false, true, true}, {false, false, true, false, false, false, false, false, true, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expectedScore := Score{Burn: 2, BHoles: -1, FHoles: -1, CHoles: 0, HighY: 7, Step: 0}

	pos := game.calculateMoves()
	pos.FieldAfter.Burn()

	//PrintVisual(field)
	checkResults(t, expectedField, pos.FieldAfter)
	checkScores(t, expectedScore, pos.Score)
}

func Test_560b136035ec1d3214e473b8_74(t *testing.T) {
	game := Game{Strategy: strategy}
	field := Field{{true, true, false, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, true, true, false, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, false, true, true, true, true, true, true, true, true}, {true, false, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, false}, {false, false, false, true, true, true, false, false, false, false}, {false, false, false, false, true, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "I", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "L", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 0
	game.MyPlayer.Empty = 1
	expectedField := Field{{true, true, false, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, true, true, false, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, false, true}, {false, true, true, true, true, true, true, true, false, false}, {false, true, false, true, true, true, false, false, false, false}, {false, false, false, false, true, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expectedScore := Score{Burn: 2, BHoles: -1, FHoles: 0, CHoles: 3, HighY: 14, Step: 3}

	pos := game.calculateMoves()

	checkResults(t, expectedField, pos.FieldAfter)
	checkScores(t, expectedScore, pos.Score)
}

func Test_563bfe8035ec1d521be3ee36_2(t *testing.T) {
	game := Game{Strategy: strategy}
	field := Field{{false, false, false, false, true, true, false, false, false, false}, {false, false, false, true, true, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "J", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "I", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 0
	game.MyPlayer.Empty = 18
	expectedField := Field{{false, true, true, true, true, true, false, false, false, false}, {false, true, false, true, true, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expectedScore := Score{Burn: 0, BHoles: 0, FHoles: -1, CHoles: 0, HighY: 0, Step: 2}

	pos := game.calculateMoves()

	checkResults(t, expectedField, pos.FieldAfter)
	checkScores(t, expectedScore, pos.Score)
}

func Test_01(t *testing.T) {
	game := Game{Strategy: strategy}
	field := Field{{true, true, true, true, true, true, true, true, false, true}, {true, true, false, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, false, true, true}, {false, true, true, true, true, true, true, false, true, false}, {false, true, false, true, true, true, false, false, false, false}, {false, false, false, false, true, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "L", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "S", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 0
	expectedField := Field{{true, true, true, true, true, true, true, true, false, true}, {true, true, false, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, false}, {false, true, false, true, true, true, true, true, false, false}, {false, false, false, false, true, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expectedScore := Score{Burn: 0, BHoles: 0, FHoles: 0, CHoles: 0, HighY: 8, Step: 1}

	pos := game.calculateMoves()

	//	fmt.Println(pos.CurrentX, pos.Rotation)
	//	PrintVisual(field)

	checkResults(t, expectedField, pos.FieldAfter)
	checkScores(t, expectedScore, pos.Score)
}
