package main

import (
	"fmt"
	"testing"
)

var arangePathField = Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, true, true, true, false, true, true, true, true, false}, {false, true, true, false, false, false, false, true, false, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, false, true, false, false, false, true, false, false}, {false, false, true, true, false, false, true, true, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

func Test_path_T(t *testing.T) {
	piece := Piece{Name: "T", Rotation: 0}
	piece.InitSpace(Cell{X: 3, Y: 19})

	hole0 := Cell{X: 3, Y: 16}
	hole1 := Cell{X: 2, Y: 13}
	hole2 := Cell{X: 3, Y: 11}
	hole3 := Cell{X: 7, Y: 12}
	hole_bad1 := Cell{X: 2, Y: 6}
	hole_bad2 := Cell{X: 8, Y: 11}

	good_positions := arangePathField.FixHoles(piece, []Cell{hole0, hole1, hole2, hole3})
	bad_positions := arangePathField.FixHoles(piece, []Cell{hole_bad1, hole_bad2})

	if len(good_positions) != 12 {
		for _, pos := range good_positions {
			fmt.Println(pos.Moves)
		}
		t.Fail()
	}

	if len(bad_positions) != 0 {
		for _, pos := range bad_positions {
			fmt.Println(pos.Moves)
		}
		t.Fail()
	}
}
