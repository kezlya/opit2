package main

import (
	"fmt"
	"testing"
)

var gameSt = Strategy{
	Burn:   4,
	Step:   1,
	BHoles: 4,
	FHoles: 4,
	LowY:   2,
	HighY:  1,
}

func Test_55c29f6435ec1d070e2b66e9_40(t *testing.T) {
	game := Game{Strategy: gameSt}
	field := Field{{true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, false, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {false, true, true, false, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, false}, {false, true, true, false, false, true, true, true, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "T", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "J", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	expectedField := Field{{true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, false, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {false, true, true, false, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, false, false, true, true, true, true, true}, {false, true, true, false, false, false, false, false, false, true}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	pos := game.calculateMoves()

	checkResults(t, expectedField, pos.FieldAfter)
}

func Test_55d7935d35ec1d06d15c9d7e_31(t *testing.T) {
	game := Game{Strategy: gameSt}
	field := Field{{true, true, true, false, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "Z", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "T", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 3
	expectedField := Field{{true, true, true, false, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, false, true, true, false, false, false, false, false, false}, {false, false, false, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expectedField.Burn()

	pos := game.calculateMoves()

	checkResults(t, expectedField, pos.FieldAfter)
}

func Test_55dc7ff01c687b0946a742f3_67(t *testing.T) {
	game := Game{Strategy: gameSt}
	field := Field{{true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, false, false, true, true, true}, {false, true, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "T", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "I", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 1
	expectedField := Field{{true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, false, true, true, true}, {false, true, true, true, true, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	pos := game.calculateMoves()

	checkResults(t, expectedField, pos.FieldAfter)
}

func Test_55dc5b501c687b0946a741a2_35(t *testing.T) {
	game := Game{Strategy: gameSt}
	field := Field{{true, false, true, true, true, true, true, true, false, false}, {true, false, false, false, true, false, false, false, false, false}, {true, false, false, false, true, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	//PrintVisual(field)
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "S", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "J", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 3
	expectedField := Field{{true, false, true, true, true, true, true, true, true, false}, {true, false, false, false, true, false, false, true, true, false}, {true, false, false, false, true, false, false, true, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	pos := game.calculateMoves()

	//fmt.Println(pos.Damage, pos.HighY, pos.LowY)

	checkResults(t, expectedField, pos.FieldAfter)
}

func Test_01(t *testing.T) {
	game := Game{Strategy: gameSt}
	field := Field{{true, true, true, true, true, true, true, true, false, true}, {true, true, false, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, false, true, true}, {false, true, true, true, true, true, true, false, true, false}, {false, true, false, true, true, true, false, false, false, false}, {false, false, false, false, true, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	//PrintVisual(field)
	game.MyPlayer = &Player{Field: field, Picks: field.Picks()}
	game.CurrentPiece = Piece{Name: "L", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "S", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 0
	expectedField := Field{{true, true, true, true, true, true, true, true, false, true}, {true, true, false, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, false}, {false, true, false, true, true, true, true, true, false, false}, {false, false, false, false, true, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	pos := game.calculateMoves()

	//	fmt.Println(pos.Score, pos.Damage, pos.HighY, pos.LowY)
	fmt.Println(pos.Moves)
	checkResults(t, expectedField, pos.FieldAfter)
}
