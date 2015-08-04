package main

import (
	"fmt"
	"testing"
	//"sort"
)

func _setup() {
	_asignSettings("field_width", "10")
	_asignSettings("field_height", "20")
	_asignSettings("player_names", "player1,player2")
	_asignSettings("your_bot", "player1")
	_asignUpdates("game", "this_piece_position", "3,1")
}

func Test_convertField(t *testing.T) {
	_setup()
	cleanInput := "0,0,0,1,1,1,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,2,0,0,0;0,0,0,0,0,0,2,0,0,0;0,0,0,0,0,0,2,2,0,0;0,0,2,2,0,0,2,2,0,0;0,2,2,2,0,0,2,2,0,2;2,2,2,2,2,0,2,2,2,2;2,2,0,2,2,2,2,2,2,2;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0"
	expect := Field{{true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, false, false, true, true, false, true}, {false, false, true, true, false, false, true, true, false, false}, {false, false, false, false, false, false, true, true, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	result := _convertField(cleanInput)

	if !expect.Equal(result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

/*
func Test_isHoleYes(t *testing.T) {
	cols := []int{1, 1, 1, 1, 1, 1, 1, 2, 1, 4}
	yes := _isHole(cols)
	if !yes {
		t.Fail()
	}
}

func Test_isHoleNo(t *testing.T) {
	cols := []int{2, 0, 2, 4, 3, 4, 2, 3, 4, 4}
	yes := _isHole(cols)
	if yes {
		t.Fail()
	}
}
*/
func Test_isBurn(t *testing.T) {
	arange := Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, false, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, false, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expect := 5

	result := _isBurn(arange)
	if expect != result {
		t.Fail()
		fmt.Println(expect, "!=", result)
	}
}

func Test_getPicks(t *testing.T) {
	arrange := Field{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expect := Picks{12, 13, 14, 14, 14, 17, 14, 14, 13, 14}

	result := arrange.Picks()

	if !expect.Equal(result) {
		fmt.Println(result)
		fmt.Println(expect)
		t.Fail()
	}
}

func Test_Sort(t *testing.T) {

	pos0 := Position{Damage: 4, LowY: 7, HighY: 5}
	pos1 := Position{Damage: 4, LowY: 7, HighY: 10}
	pos2 := Position{Damage: 4, LowY: 13, HighY: 3}
	pos3 := Position{Damage: 5, LowY: 1, HighY: 3}
	pos4 := Position{Damage: 5, LowY: 5, HighY: 1}
	pos5 := Position{Damage: 5, LowY: 11, HighY: 4}
	pos6 := Position{Damage: 6, LowY: 1, HighY: 2}

	positions := []Position{pos4, pos0, pos5, pos2, pos6, pos1, pos3}

	OrderedBy(DAMAGE, LOWY, HIGHY).Sort(positions)

	if !(positions[0].HighY == 5 && positions[1].HighY == 10 && positions[2].HighY == 3 && positions[3].HighY == 3 && positions[4].HighY == 1 && positions[5].HighY == 4) {
		t.Fail()
		for _, pos := range positions {
			fmt.Println("damadge:", pos.Damage, "LowY:", pos.LowY, "HighY:", pos.HighY)
		}
	}

}
