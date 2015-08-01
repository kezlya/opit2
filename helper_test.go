package main

import (
	"fmt"
	"testing"
)

func _setup() {
	_asignSettings("field_width", "10")
	_asignSettings("field_height", "20")
	_asignSettings("player_names", "player1,player2")
	_asignSettings("your_bot", "player1")
	_asignUpdates("game", "this_piece_position", "3,1")
}

func Test_getMoves(t *testing.T) {
	_setup() //round 32 http://theaigames.com/competitions/ai-block-battle/games/55b3f2831c687b544a2b9291
	_asignUpdates("game", "this_piece_type", "L")
	_asignUpdates("player1", "field", "0,0,0,1,1,1,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,2,0,0,0;0,0,0,0,0,0,2,0,0,0;0,0,0,0,0,0,2,2,0,0;0,0,2,2,0,0,2,2,0,0;0,2,2,2,0,0,2,2,0,2;2,2,2,2,2,0,2,2,2,2;2,2,0,2,2,2,2,2,2,2;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;3,3,3,3,3,3,3,3,3,3;3,3,3,3,3,3,3,3,3,3")
	result := _calculateMoves(500)
	if result.Rotation != 3 || result.X != 4 {
		fmt.Println("rotation", 3, "!=", result.Rotation, "X", 4, "!=", result.X)
		t.Fail()
	}
}

func Test_convertField(t *testing.T) {
	_setup()
	cleanInput := "0,0,0,1,1,1,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,2,0,0,0;0,0,0,0,0,0,2,0,0,0;0,0,0,0,0,0,2,2,0,0;0,0,2,2,0,0,2,2,0,0;0,2,2,2,0,0,2,2,0,2;2,2,2,2,2,0,2,2,2,2;2,2,0,2,2,2,2,2,2,2;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0"
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, false, false, true, true, false, true}, {false, false, true, true, false, false, true, true, false, false}, {false, false, false, false, false, false, true, true, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	result, _ := _convertField(cleanInput)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_getAllPossiblePositionsI(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "I")
	posiblePositions := 17
	MyPlayer.Columns = []int{2, 0, 3, 1, 3, 4, 0, 3, 3, 3}
	result, bestScore := _getAllPossiblePositions()
	if len(result) != posiblePositions {
		fmt.Println(posiblePositions, "!=", len(result))
		t.Fail()
	}
	expectedScore := 8
	if bestScore != expectedScore {
		fmt.Println(bestScore, "!=", expectedScore)
		t.Fail()
	}
}

func Test_getAllPossiblePositionsJ(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "J")
	posiblePositions := 34
	MyPlayer.Columns = []int{2, 0, 3, 1, 3, 4, 0, 3, 3, 3}
	result, bestScore := _getAllPossiblePositions()
	if len(result) != posiblePositions {
		fmt.Println(posiblePositions, "!=", len(result))
		t.Fail()
	}
	expectedScore := 8
	if bestScore != expectedScore {
		fmt.Println(bestScore, "!=", expectedScore)
		t.Fail()
	}
}

func Test_getAllPossiblePositionsL(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "L")
	posiblePositions := 34
	MyPlayer.Columns = []int{2, 0, 3, 1, 3, 4, 0, 3, 3, 3}
	result, bestScore := _getAllPossiblePositions()
	if len(result) != posiblePositions {
		fmt.Println(posiblePositions, "!=", len(result))
		t.Fail()
	}
	expectedScore := 7
	if bestScore != expectedScore {
		fmt.Println(bestScore, "!=", expectedScore)
		t.Fail()
	}
}

func Test_getAllPossiblePositionsO(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "O")
	posiblePositions := 9
	MyPlayer.Columns = []int{2, 0, 3, 1, 3, 4, 0, 3, 3, 3}
	result, bestScore := _getAllPossiblePositions()
	if len(result) != posiblePositions {
		fmt.Println(posiblePositions, "!=", len(result))
		t.Fail()
	}
	expectedScore := 9
	if bestScore != expectedScore {
		fmt.Println(bestScore, "!=", expectedScore)
		t.Fail()
	}
}

func Test_getAllPossiblePositionsS(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "S")
	posiblePositions := 17
	MyPlayer.Columns = []int{2, 0, 3, 1, 3, 4, 0, 3, 3, 3}
	result, bestScore := _getAllPossiblePositions()
	if len(result) != posiblePositions {
		fmt.Println(posiblePositions, "!=", len(result))
		t.Fail()
	}
	expectedScore := 9
	if bestScore != expectedScore {
		fmt.Println(bestScore, "!=", expectedScore)
		t.Fail()
	}
}

func Test_getAllPossiblePositionsT(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "T")
	posiblePositions := 34
	MyPlayer.Columns = []int{2, 0, 3, 1, 3, 4, 0, 3, 3, 3}
	result, bestScore := _getAllPossiblePositions()
	if len(result) != posiblePositions {
		fmt.Println(posiblePositions, "!=", len(result))
		t.Fail()
	}
	expectedScore := 9
	if bestScore != expectedScore {
		fmt.Println(bestScore, "!=", expectedScore)
		t.Fail()
	}
}

func Test_getAllPossiblePositionsZ(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "Z")
	posiblePositions := 17
	MyPlayer.Columns = []int{2, 0, 3, 1, 3, 4, 0, 3, 3, 3}
	result, bestScore := _getAllPossiblePositions()
	if len(result) != posiblePositions {
		fmt.Println(posiblePositions, "!=", len(result))
		t.Fail()
	}
	expectedScore := 10
	if bestScore != expectedScore {
		fmt.Println(bestScore, "!=", expectedScore)
		t.Fail()
	}
}

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

func Test_getPicks(t *testing.T) {
	arrange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expect := []int{12, 13, 14, 14, 14, 17, 14, 14, 13, 14}

	result := _getPicks(arrange)

	if !_eq1(expect, result) {
		fmt.Println(result)
		fmt.Println(expect)
		t.Fail()
	}
}

func Test_I1(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "I"
	rotation := 0
	x := 4
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, true, true, true, true, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_I2(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "I"
	rotation := 0
	x := 2
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, true, true, true, true, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_I3(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "I"
	rotation := 1
	x := 6
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_I_FailTall(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "I"
	rotation := 1
	x := 6
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_J1(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "J"
	rotation := 0
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, true, true, true, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_J2_1(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "J"
	rotation := 1
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, true, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_J2_2(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "J"
	rotation := 1
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_J3_1(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "J"
	rotation := 2
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, true, false, false, true, false, false, false}, {false, true, true, true, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_J3_2(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "J"
	rotation := 2
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, true, false, false, true, false, false, false}, {false, true, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_J4(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "J"
	rotation := 3
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, true, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_L1(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "L"
	rotation := 0
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, true, true, true, false, false, true, false, false, false}, {false, false, false, true, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_L2(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "L"
	rotation := 1
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, true, true, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_L3_1(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "L"
	rotation := 2
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_L3_2(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "L"
	rotation := 2
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, true, true, false, false, false, true, false, false, false}, {false, true, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_L4_1(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "L"
	rotation := 3
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_L4_2(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "L"
	rotation := 3
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_O(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "O"
	rotation := 0
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, true, true, false, false, false, true, false, false, false}, {false, true, true, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_S1_1(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "S"
	rotation := 0
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, true, true, false, false, false, true, false, false, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_S1_2(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, true, false, false, true, false, false, false}, {false, false, false, true, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "S"
	rotation := 0
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, true, false, false, true, false, false, false}, {false, true, true, true, false, false, true, false, false, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_S2_1(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "S"
	rotation := 1
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_S2_2(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "S"
	rotation := 1
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_T1(t *testing.T) {
	arange :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "T"
	rotation :=0
	x :=1
	expect :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, true, true, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result,_ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect,result) {
		t.Fail()
		y:= len(expect)-1
		for i := range expect {
			fmt.Println(expect[y-i],result[y-i])
		}
	}
}

func Test_T2_1(t *testing.T) {
	arange :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "T"
	rotation :=1
	x :=1
	expect :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result,_ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect,result) {
		t.Fail()
		y:= len(expect)-1
		for i := range expect {
			fmt.Println(expect[y-i],result[y-i])
		}
	}
}

func Test_T2_2(t *testing.T) {
	arange :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "T"
	rotation :=1
	x :=1
	expect :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result,_ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect,result) {
		t.Fail()
		y:= len(expect)-1
		for i := range expect {
			fmt.Println(expect[y-i],result[y-i])
		}
	}
}

func Test_T3_1(t *testing.T) {
	arange :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "T"
	rotation :=2
	x :=1
	expect :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, true, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result,_ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect,result) {
		t.Fail()
		y:= len(expect)-1
		for i := range expect {
			fmt.Println(expect[y-i],result[y-i])
		}
	}
}

func Test_T3_2(t *testing.T) {
	arange :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "T"
	rotation :=2
	x :=1
	expect :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result,_ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect,result) {
		t.Fail()
		y:= len(expect)-1
		for i := range expect {
			fmt.Println(expect[y-i],result[y-i])
		}
	}
}

func Test_T4_1(t *testing.T) {
	arange :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "T"
	rotation :=3
	x :=1
	expect :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result,_ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect,result) {
		t.Fail()
		y:= len(expect)-1
		for i := range expect {
			fmt.Println(expect[y-i],result[y-i])
		}
	}
}

func Test_T4_2(t *testing.T) {
	arange :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "T"
	rotation :=3
	x :=1
	expect :=[][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, false, false, false, false, false, false, false, false}}

	result,_ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect,result) {
		t.Fail()
		y:= len(expect)-1
		for i := range expect {
			fmt.Println(expect[y-i],result[y-i])
		}
	}
}

func Test_Z1_1(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "Z"
	rotation := 0
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, true, true, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_Z1_2(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "Z"
	rotation := 0
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, true, true, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_Z2_1(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "Z"
	rotation := 1
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_Z2_2(t *testing.T) {
	arange := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := "Z"
	rotation := 1
	x := 1
	expect := [][]bool{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, true, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, true, false, false, false, false, false, false, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}}

	result, _ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect, result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

/*
func Test_(t *testing.T) {
	arange :=
	MyPlayer.Columns = _getPicks(arange)
	Height = len(arange)
	piece := ""
	rotation :=
	x :=
	expect :=

	result,_ := _fieldAfter(arange, x, rotation, piece)

	if !_eq2(expect,result) {
		t.Fail()
		y:= len(expect)-1
		for i := range expect {
			fmt.Println(expect[y-i],result[y-i])
		}
	}
}
*/

func _eq2(a, b [][]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func _eq1(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
