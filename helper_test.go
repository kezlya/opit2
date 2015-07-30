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

func Test_fieldAfter(t *testing.T) {
	Width = 10
	arange := [][]bool{{false, false, false, false}, {false, false, false, true}}
	arange[0][0] = true
	arange[0][1] = true

	fa, burned := _fieldAfter(arange, 1, 1, "I")

	fmt.Println(arange)
	fmt.Println(fa)
	fmt.Println(burned)

}

func Test_hh(t *testing.T) {
	arange = "0,0,0,1,1,1,0,0,0,0;" +
		"0,0,0,0,0,0,0,0,0,0;" +
		"0,0,0,0,0,0,0,0,0,0;" +
		"0,0,0,0,0,0,2,0,0,0;" +
		"0,0,0,0,0,0,2,0,0,0;" +
		"0,0,0,0,0,0,2,2,0,0;" +
		"0,0,2,2,0,0,2,2,0,0;" +
		"0,2,2,2,0,0,2,2,0,2;" +
		"2,2,2,2,2,0,2,2,2,2;" +
		"2,2,0,2,2,2,2,2,2,2;" +
		"2,2,2,2,2,2,2,2,2,0;" +
		"2,2,2,2,2,2,2,2,2,0;" +
		"2,2,2,2,2,2,2,2,2,0;" +
		"2,2,2,2,2,2,2,2,2,0;" +
		"2,2,2,2,2,2,2,2,2,0;" +
		"2,2,2,2,2,2,2,2,2,0;" +
		"2,2,2,2,2,2,2,2,2,0;" +
		"2,2,2,2,2,2,2,2,2,0;" +
		"3,3,3,3,3,3,3,3,3,3;" +
		"3,3,3,3,3,3,3,3,3,3"

}
