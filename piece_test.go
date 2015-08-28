package main

import (
	"fmt"
	"testing"
)

func pieceAssert(t *testing.T, a, b Piece) {
	pass := true
	for i, cell := range a.Space {
		if b.Space[i].X != cell.X || b.Space[i].Y != cell.Y {
			pass = false
			break
		}
	}
	if !pass {
		fmt.Println(a.Space)
		fmt.Println(b.Space)
		t.Fail()
	}

	if a.Name != b.Name {
		fmt.Println(a.Name, b.Name)
		t.Fail()
	}

	if a.Rotation != b.Rotation {
		fmt.Println(a.Rotation, b.Rotation)
		t.Fail()
	}

	if a.CurrentX != b.CurrentX {
		fmt.Println(a.CurrentX, b.CurrentX)
		t.Fail()
	}
}

func Test_Left(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["t3"] = Cell{X: 3, Y: 1}
	space["m3"] = Cell{X: 3, Y: 2}
	space["b3"] = Cell{X: 3, Y: 3}
	space["x3"] = Cell{X: 3, Y: 4}
	arrange := Piece{Space: space, Name: "I", Rotation: 1, CurrentX: 3}

	ex_space := make(map[string]Cell, 4)
	ex_space["t3"] = Cell{X: 2, Y: 1}
	ex_space["m3"] = Cell{X: 2, Y: 2}
	ex_space["b3"] = Cell{X: 2, Y: 3}
	ex_space["x3"] = Cell{X: 2, Y: 4}
	expected := Piece{Space: ex_space, Name: "I", Rotation: 1, CurrentX: 2}

	result := arrange.Left()

	pieceAssert(t, expected, result)
}

func Test_Right(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["t3"] = Cell{X: 3, Y: 1}
	space["m3"] = Cell{X: 3, Y: 2}
	space["b3"] = Cell{X: 3, Y: 3}
	space["x3"] = Cell{X: 3, Y: 4}
	arrange := Piece{Space: space, Name: "I", Rotation: 1, CurrentX: 3}

	ex_space := make(map[string]Cell, 4)
	ex_space["t3"] = Cell{X: 4, Y: 1}
	ex_space["m3"] = Cell{X: 4, Y: 2}
	ex_space["b3"] = Cell{X: 4, Y: 3}
	ex_space["x3"] = Cell{X: 4, Y: 4}
	expected := Piece{Space: ex_space, Name: "I", Rotation: 1, CurrentX: 4}

	result := arrange.Right()

	pieceAssert(t, expected, result)
}

func Test_Down(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["t3"] = Cell{X: 3, Y: 1}
	space["m3"] = Cell{X: 3, Y: 2}
	space["b3"] = Cell{X: 3, Y: 3}
	space["x3"] = Cell{X: 3, Y: 4}
	arrange := Piece{Space: space, Name: "I", Rotation: 1, CurrentX: 3}

	ex_space := make(map[string]Cell, 4)
	ex_space["t3"] = Cell{X: 3, Y: 0}
	ex_space["m3"] = Cell{X: 3, Y: 1}
	ex_space["b3"] = Cell{X: 3, Y: 2}
	ex_space["x3"] = Cell{X: 3, Y: 3}
	expected := Piece{Space: ex_space, Name: "I", Rotation: 1, CurrentX: 3}

	result := arrange.Down()

	pieceAssert(t, expected, result)
}

func Test_TurnRight_I0(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["m1"] = Cell{X: 4, Y: 12}
	space["m2"] = Cell{X: 5, Y: 12}
	space["m3"] = Cell{X: 6, Y: 12}
	space["m4"] = Cell{X: 7, Y: 12}
	arrange := Piece{Space: space, Name: "I", Rotation: 0, CurrentX: 4}

	ex_space := make(map[string]Cell, 4)
	ex_space["t3"] = Cell{X: 6, Y: 13} //
	ex_space["m3"] = Cell{X: 6, Y: 12}
	ex_space["b3"] = Cell{X: 6, Y: 11} //
	ex_space["x3"] = Cell{X: 6, Y: 10} //
	expected := Piece{Space: ex_space, Name: "I", Rotation: 1, CurrentX: 6}

	result := arrange.Turnright()

	pieceAssert(t, expected, result)
}

func Test_TurnRight_I1(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["t3"] = Cell{X: 6, Y: 13}
	space["m3"] = Cell{X: 6, Y: 12}
	space["b3"] = Cell{X: 6, Y: 11}
	space["x3"] = Cell{X: 6, Y: 10}
	arrange := Piece{Space: space, Name: "I", Rotation: 1, CurrentX: 6}

	ex_space := make(map[string]Cell, 4)
	ex_space["b1"] = Cell{X: 4, Y: 11} //
	ex_space["b2"] = Cell{X: 5, Y: 11} //
	ex_space["b3"] = Cell{X: 6, Y: 11}
	ex_space["b4"] = Cell{X: 7, Y: 11} //
	expected := Piece{Space: ex_space, Name: "I", Rotation: 2, CurrentX: 4}

	result := arrange.Turnright()

	pieceAssert(t, expected, result)
}

func Test_TurnRight_I2(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["b1"] = Cell{X: 4, Y: 11}
	space["b2"] = Cell{X: 5, Y: 11}
	space["b3"] = Cell{X: 6, Y: 11}
	space["b4"] = Cell{X: 7, Y: 11}
	arrange := Piece{Space: space, Name: "I", Rotation: 2, CurrentX: 4}

	ex_space := make(map[string]Cell, 4)
	ex_space["t2"] = Cell{X: 5, Y: 13} //
	ex_space["m2"] = Cell{X: 5, Y: 12} //
	ex_space["b2"] = Cell{X: 5, Y: 11}
	ex_space["x2"] = Cell{X: 5, Y: 10} //
	expected := Piece{Space: ex_space, Name: "I", Rotation: 3, CurrentX: 5}

	result := arrange.Turnright()

	pieceAssert(t, expected, result)
}

func Test_TurnRight_I3(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["t2"] = Cell{X: 5, Y: 13}
	space["m2"] = Cell{X: 5, Y: 12}
	space["b2"] = Cell{X: 5, Y: 11}
	space["x2"] = Cell{X: 5, Y: 10}
	arrange := Piece{Space: space, Name: "I", Rotation: 3, CurrentX: 5}

	ex_space := make(map[string]Cell, 4)
	ex_space["m1"] = Cell{X: 4, Y: 12} //
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["m3"] = Cell{X: 6, Y: 12} //
	ex_space["m4"] = Cell{X: 7, Y: 12} //
	expected := Piece{Space: ex_space, Name: "I", Rotation: 0, CurrentX: 4}

	result := arrange.Turnright()

	pieceAssert(t, expected, result)
}

/*
func Test_TurnLeft_I0(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["t2"] = Cell{X: 5, Y: 13}
	space["m1"] = Cell{X: 4, Y: 12}
	space["m2"] = Cell{X: 5, Y: 12}
	space["m3"] = Cell{X: 6, Y: 12}
	arrange := Piece{Space: space, Name: "T", Rotation: 0, CurrentX: 4}

	ex_space := make(map[string]Cell, 4)
	ex_space["t2"] = Cell{X: 5, Y: 13}
	ex_space["m1"] = Cell{X: 4, Y: 12}
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["b2"] = Cell{X: 5, Y: 11} //
	expected := Piece{Space: ex_space, Name: "T", Rotation: 3, CurrentX: 4}

	result := arrange.Turnleft()

	pieceAssert(t, expected, result)
}

func Test_TurnLeft_I3(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["t2"] = Cell{X: 5, Y: 13}
	space["m1"] = Cell{X: 4, Y: 12}
	space["m2"] = Cell{X: 5, Y: 12}
	space["b2"] = Cell{X: 5, Y: 11}
	arrange := Piece{Space: space, Name: "T", Rotation: 3, CurrentX: 4}

	ex_space := make(map[string]Cell, 4)
	ex_space["m3"] = Cell{X: 6, Y: 12} //
	ex_space["m1"] = Cell{X: 4, Y: 12}
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["b2"] = Cell{X: 5, Y: 11}
	expected := Piece{Space: ex_space, Name: "T", Rotation: 2, CurrentX: 4}

	result := arrange.Turnleft()

	pieceAssert(t, expected, result)
}

func Test_TurnLeft_I2(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["m3"] = Cell{X: 6, Y: 12}
	space["m1"] = Cell{X: 4, Y: 12}
	space["m2"] = Cell{X: 5, Y: 12}
	space["b2"] = Cell{X: 5, Y: 11}
	arrange := Piece{Space: space, Name: "T", Rotation: 2, CurrentX: 4}

	ex_space := make(map[string]Cell, 4)
	ex_space["m3"] = Cell{X: 6, Y: 12}
	ex_space["t2"] = Cell{X: 5, Y: 13} //
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["b2"] = Cell{X: 5, Y: 11}
	expected := Piece{Space: ex_space, Name: "T", Rotation: 1, CurrentX: 5}

	result := arrange.Turnleft()

	pieceAssert(t, expected, result)
}

func Test_TurnLeft_I1(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["m3"] = Cell{X: 6, Y: 12}
	space["t2"] = Cell{X: 5, Y: 13}
	space["m2"] = Cell{X: 5, Y: 12}
	space["b2"] = Cell{X: 5, Y: 11}
	arrange := Piece{Space: space, Name: "T", Rotation: 1, CurrentX: 5}

	ex_space := make(map[string]Cell, 4)
	ex_space["m3"] = Cell{X: 6, Y: 12}
	ex_space["t2"] = Cell{X: 5, Y: 13}
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["m1"] = Cell{X: 4, Y: 12} //
	expected := Piece{Space: ex_space, Name: "T", Rotation: 0, CurrentX: 4}

	result := arrange.Turnleft()

	pieceAssert(t, expected, result)
}
*/
func Test_TurnRight_T0(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["t2"] = Cell{X: 5, Y: 13}
	space["m1"] = Cell{X: 4, Y: 12}
	space["m2"] = Cell{X: 5, Y: 12}
	space["m3"] = Cell{X: 6, Y: 12}
	arrange := Piece{Space: space, Name: "T", Rotation: 0, CurrentX: 4}

	ex_space := make(map[string]Cell, 4)
	ex_space["t2"] = Cell{X: 5, Y: 13}
	ex_space["b2"] = Cell{X: 5, Y: 11} //
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["m3"] = Cell{X: 6, Y: 12}
	expected := Piece{Space: ex_space, Name: "T", Rotation: 1, CurrentX: 5}

	result := arrange.Turnright()

	pieceAssert(t, expected, result)
}

func Test_TurnRight_T1(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["t2"] = Cell{X: 5, Y: 13}
	space["b2"] = Cell{X: 5, Y: 11}
	space["m2"] = Cell{X: 5, Y: 12}
	space["m3"] = Cell{X: 6, Y: 12}
	arrange := Piece{Space: space, Name: "T", Rotation: 1, CurrentX: 5}

	ex_space := make(map[string]Cell, 4)
	ex_space["m1"] = Cell{X: 4, Y: 12} //
	ex_space["b2"] = Cell{X: 5, Y: 11}
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["m3"] = Cell{X: 6, Y: 12}
	expected := Piece{Space: ex_space, Name: "T", Rotation: 2, CurrentX: 4}

	result := arrange.Turnright()

	pieceAssert(t, expected, result)
}

func Test_TurnRight_T2(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["m1"] = Cell{X: 4, Y: 12}
	space["b2"] = Cell{X: 5, Y: 11}
	space["m2"] = Cell{X: 5, Y: 12}
	space["m3"] = Cell{X: 6, Y: 12}
	arrange := Piece{Space: space, Name: "T", Rotation: 2, CurrentX: 4}

	ex_space := make(map[string]Cell, 4)
	ex_space["m1"] = Cell{X: 4, Y: 12}
	ex_space["b2"] = Cell{X: 5, Y: 11}
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["t2"] = Cell{X: 5, Y: 13} //
	expected := Piece{Space: ex_space, Name: "T", Rotation: 3, CurrentX: 4}

	result := arrange.Turnright()

	pieceAssert(t, expected, result)
}

func Test_TurnRight_T3(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["m1"] = Cell{X: 4, Y: 12}
	space["b2"] = Cell{X: 5, Y: 11}
	space["m2"] = Cell{X: 5, Y: 12}
	space["t2"] = Cell{X: 5, Y: 13}
	arrange := Piece{Space: space, Name: "T", Rotation: 3, CurrentX: 4}

	ex_space := make(map[string]Cell, 4)
	ex_space["m1"] = Cell{X: 4, Y: 12}
	ex_space["m3"] = Cell{X: 6, Y: 12} //
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["t2"] = Cell{X: 5, Y: 13}
	expected := Piece{Space: ex_space, Name: "T", Rotation: 0, CurrentX: 4}

	result := arrange.Turnright()

	pieceAssert(t, expected, result)
}

func Test_TurnLeft_T0(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["t2"] = Cell{X: 5, Y: 13}
	space["m1"] = Cell{X: 4, Y: 12}
	space["m2"] = Cell{X: 5, Y: 12}
	space["m3"] = Cell{X: 6, Y: 12}
	arrange := Piece{Space: space, Name: "T", Rotation: 0, CurrentX: 4}

	ex_space := make(map[string]Cell, 4)
	ex_space["t2"] = Cell{X: 5, Y: 13}
	ex_space["m1"] = Cell{X: 4, Y: 12}
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["b2"] = Cell{X: 5, Y: 11} //
	expected := Piece{Space: ex_space, Name: "T", Rotation: 3, CurrentX: 4}

	result := arrange.Turnleft()

	pieceAssert(t, expected, result)
}

func Test_TurnLeft_T3(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["t2"] = Cell{X: 5, Y: 13}
	space["m1"] = Cell{X: 4, Y: 12}
	space["m2"] = Cell{X: 5, Y: 12}
	space["b2"] = Cell{X: 5, Y: 11}
	arrange := Piece{Space: space, Name: "T", Rotation: 3, CurrentX: 4}

	ex_space := make(map[string]Cell, 4)
	ex_space["m3"] = Cell{X: 6, Y: 12} //
	ex_space["m1"] = Cell{X: 4, Y: 12}
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["b2"] = Cell{X: 5, Y: 11}
	expected := Piece{Space: ex_space, Name: "T", Rotation: 2, CurrentX: 4}

	result := arrange.Turnleft()

	pieceAssert(t, expected, result)
}

func Test_TurnLeft_T2(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["m3"] = Cell{X: 6, Y: 12}
	space["m1"] = Cell{X: 4, Y: 12}
	space["m2"] = Cell{X: 5, Y: 12}
	space["b2"] = Cell{X: 5, Y: 11}
	arrange := Piece{Space: space, Name: "T", Rotation: 2, CurrentX: 4}

	ex_space := make(map[string]Cell, 4)
	ex_space["m3"] = Cell{X: 6, Y: 12}
	ex_space["t2"] = Cell{X: 5, Y: 13} //
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["b2"] = Cell{X: 5, Y: 11}
	expected := Piece{Space: ex_space, Name: "T", Rotation: 1, CurrentX: 5}

	result := arrange.Turnleft()

	pieceAssert(t, expected, result)
}

func Test_TurnLeft_T1(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["m3"] = Cell{X: 6, Y: 12}
	space["t2"] = Cell{X: 5, Y: 13}
	space["m2"] = Cell{X: 5, Y: 12}
	space["b2"] = Cell{X: 5, Y: 11}
	arrange := Piece{Space: space, Name: "T", Rotation: 1, CurrentX: 5}

	ex_space := make(map[string]Cell, 4)
	ex_space["m3"] = Cell{X: 6, Y: 12}
	ex_space["t2"] = Cell{X: 5, Y: 13}
	ex_space["m2"] = Cell{X: 5, Y: 12}
	ex_space["m1"] = Cell{X: 4, Y: 12} //
	expected := Piece{Space: ex_space, Name: "T", Rotation: 0, CurrentX: 4}

	result := arrange.Turnleft()

	pieceAssert(t, expected, result)
}
