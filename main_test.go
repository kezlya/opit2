package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	//arrange
	g := &Game{Strategy: strategy}
	g.asignSettings("timebank", "10000")
	g.asignSettings("time_per_move", "500")
	g.asignSettings("player_names", "player1,player2")
	g.asignSettings("your_bot", "player1")
	g.asignSettings("field_width", "10")
	g.asignSettings("field_height", "20")
	g.asignUpdates("game", "round", "2")
	g.asignUpdates("game", "this_piece_type", J)
	g.asignUpdates("game", "next_piece_type", Z)
	g.asignUpdates("game", "this_piece_position", "3,-1")
	g.asignUpdates("player1", "row_points", "0")
	g.asignUpdates("player1", "combo", "0")
	g.asignUpdates("player1", "skips", "0")
	g.asignUpdates("player1", "field", "0,0,0,1,1,1,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,2,2,2,2,0,0,0")
	g.asignUpdates("player2", "row_points", "0")
	g.asignUpdates("player2", "combo", "0")
	g.asignUpdates("player2", "skips", "0")
	g.asignUpdates("player2", "field", "0,0,0,1,1,1,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,2,0,0;0,0,0,0,0,0,0,2,0,0;0,0,0,0,0,0,0,2,0,0;0,0,0,0,0,0,0,2,0,0")

	//act
	position := g.calculateMoves()

	//assert
	if position == nil {
		t.Fail()
		fmt.Println("Positon not found")
	} else if position.Moves == "" {
		t.Fail()
		fmt.Println("No moves computed")
	}
}
