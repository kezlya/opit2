package main

import (
	//"fmt"
	"testing"
)

var arangePathField = Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, true, true, true, false, true, true, true, true, false}, {false, true, true, false, false, false, false, true, false, false}, {false, false, true, true, false, false, false, true, false, false}, {false, false, false, true, false, false, false, true, false, false}, {false, false, true, true, false, false, true, true, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

func Test_pathTopPos_I(t *testing.T) {
	piece := Piece{Name: "I", Rotation: 0}
	piece.InitSpace(Cell{X: 3, Y: 19})
	//pos1 := Position{X: 0, Rotation: 1}
	/*pos2 := Position{X: 3, Rotation: 1}
	pos3 := Position{X: 4, Rotation: 0}
	pos4 := Position{X: 9, Rotation: 1}
	pos5 := Position{X: 3, Rotation: 0}
	pos6 := Position{X: 7, Rotation: 1}*/

	//path1 := arangePathField.CalculatePath(pos1, piece)
	/*path2 := arangePathField.CalculatePath(pos2, piece)
	path3 := arangePathField.CalculatePath(pos3, piece)
	path4 := arangePathField.CalculatePath(pos4, piece)
	path5 := arangePathField.CalculatePath(pos5, piece)
	path6 := arangePathField.CalculatePath(pos6, piece)*/

	/*if path1 != "left,left,left,turnleft,left,drop" {
		fmt.Println("bad path1:", path1)
		t.Fail()
	}*/
	/*if path2 != "left,turnleft,drop" {
		fmt.Println("bad path2:", path2)
		t.Fail()
	}
	if path3 != "right,drop" {
		fmt.Println("bad path3:", path3)
		t.Fail()
	}
	if path4 != "" {
		fmt.Println("bad path4:", path4)
		t.Fail()
	}
	if path5 != "drop" {
		fmt.Println("bad path5:", path5)
		t.Fail()
	}
	if path6 != "right,right,turnright,drop" {
		fmt.Println("bad path6:", path6)
		t.Fail()
	}*/
}
