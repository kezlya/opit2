package main

import (
	"fmt"
	"testing"
)

var gameSt = Strategy{DamageK: 9, PostyK: 3, StepK: 1, BurnK: 8}

/*
func Test_55c29f6435ec1d070e2b66e9_40(t *testing.T) {
	game := Game{Strategy: gameSt}
	field := Field{{true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, false, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {false, true, true, false, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, false}, {false, true, true, false, false, true, true, true, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field}
	game.CurrentPiece = Piece{Name: "T", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "J", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})

	pos := game.calculateMoves()

	if pos.X != 8 || pos.Rotation != 3 {
		t.Fail()
		fmt.Println(pos.Score, pos.X, pos.Rotation, pos.Hole)
		PrintVisual(field)
	}
}
*/
func Test_55d7935d35ec1d06d15c9d7e_31(t *testing.T) {
	game := Game{Strategy: gameSt}
	field := Field{{true, true, true, false, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field}
	game.CurrentPiece = Piece{Name: "Z", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "T", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 3

	pos := game.calculateMoves()

	if pos.X != 2 || pos.Rotation != 1 {
		t.Fail()
		fmt.Println(pos.Score, pos.X, pos.Rotation, pos.Hole)
		PrintVisual(field)
	}
}

/*
func Test_55dc7ff01c687b0946a742f3_67(t *testing.T) {
	game := Game{Strategy: gameSt}
	field := Field{{true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, false, false, true, true, true}, {false, true, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field}
	game.CurrentPiece = Piece{Name: "T", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "I", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 1

	pos := game.calculateMoves()

	if pos.X != 4 || pos.Rotation != 3 {
		t.Fail()
		fmt.Println(pos.Score, pos.X, pos.Rotation, pos.Hole)
		fmt.Println(pos.Damage, pos.HighY, pos.Burn, pos.Step)
		PrintVisual(field)
	}
}

func Test_55dc5b501c687b0946a741a2_35(t *testing.T) {
	game := Game{Strategy: gameSt}
	field := Field{{true, false, true, true, true, true, true, true, false, false}, {true, false, false, false, true, false, false, false, false, false}, {true, false, false, false, true, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field}
	game.CurrentPiece = Piece{Name: "S", Rotation: 0}
	game.CurrentPiece.InitSpace(Cell{3, field.Height()})
	game.NextPiece = Piece{Name: "J", Rotation: 0}
	game.NextPiece.InitSpace(Cell{3, field.Height()})
	game.MyPlayer.Combo = 3

	pos := game.calculateMoves()

	if pos.X != 7 || pos.Rotation != 1 {
		t.Fail()
		fmt.Println(pos.Score, pos.X, pos.Rotation, pos.Hole)
		PrintVisual(field)
	}
}
*/
/*
func Test_55c2d43635ec1d070e2b697c_63(t *testing.T) {
	game := Game{
		DamageK: 1,
		HoleK:   1,
		PostyK:  1,
		BurnK:   1,
	}
	field := Field{{false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, false, true, true, true, true}, {true, true, true, true, true, false, true, true, true, false}, {true, false, false, false, false, false, false, true, true, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field}
	game.CurrentPiece = "I"
	game.NextPiece = "L"

	pos := game.calculateMoves()

	if pos.X != 1 || pos.Rotation != 0 {
		t.Fail()
		fmt.Println(pos.X, pos.Rotation)
	}
}
*/
