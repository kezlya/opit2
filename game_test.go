package main

import (
	"fmt"
	"testing"
)

func Test_55c29f6435ec1d070e2b66e9_40(t *testing.T) {
	game := Game{
		DamageK: 1,
		HoleK:   1,
		PostyK:  1,
		BurnK:   1,
	}
	field := Field{{true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, false, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {false, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, false, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {false, true, true, false, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, false}, {false, true, true, false, false, true, true, true, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	game.MyPlayer = &Player{Field: field}
	game.CurrentPiece = "T"
	game.NextPiece = "O"

	pos := game.calculateMoves()

	if pos.X != 8 || pos.Rotation != 3 {
		t.Fail()
		fmt.Println(pos.X, pos.Rotation)
	}
}

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
