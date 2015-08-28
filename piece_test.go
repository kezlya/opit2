package main

import (
	"fmt"
	"testing"
)

func checkSpaceEqual(t *testing.T, a, b Piece) {
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

	checkSpaceEqual(t, expected, result)

	if expected.Name != result.Name {
		fmt.Println(expected.Name, result.Name)
		t.Fail()
	}

	if expected.Rotation != result.Rotation {
		fmt.Println(expected.Rotation, result.Rotation)
		t.Fail()
	}

	if expected.CurrentX != result.CurrentX {
		fmt.Println(expected.CurrentX, result.CurrentX)
		t.Fail()
	}
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

	checkSpaceEqual(t, expected, result)

	if expected.Name != result.Name {
		fmt.Println(expected.Name, result.Name)
		t.Fail()
	}

	if expected.Rotation != result.Rotation {
		fmt.Println(expected.Rotation, result.Rotation)
		t.Fail()
	}

	if expected.CurrentX != result.CurrentX {
		fmt.Println(expected.CurrentX, result.CurrentX)
		t.Fail()
	}
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

	checkSpaceEqual(t, expected, result)

	if expected.Name != result.Name {
		fmt.Println(expected.Name, result.Name)
		t.Fail()
	}

	if expected.Rotation != result.Rotation {
		fmt.Println(expected.Rotation, result.Rotation)
		t.Fail()
	}

	if expected.CurrentX != result.CurrentX {
		fmt.Println(expected.CurrentX, result.CurrentX)
		t.Fail()
	}
}

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

	checkSpaceEqual(t, expected, result)

	if expected.Name != result.Name {
		fmt.Println(expected.Name, result.Name)
		t.Fail()
	}

	if expected.Rotation != result.Rotation {
		fmt.Println(expected.Rotation, result.Rotation)
		t.Fail()
	}

	if expected.CurrentX != result.CurrentX {
		fmt.Println(expected.CurrentX, result.CurrentX)
		t.Fail()
	}
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

	checkSpaceEqual(t, expected, result)

	if expected.Name != result.Name {
		fmt.Println(expected.Name, result.Name)
		t.Fail()
	}

	if expected.Rotation != result.Rotation {
		fmt.Println(expected.Rotation, result.Rotation)
		t.Fail()
	}

	if expected.CurrentX != result.CurrentX {
		fmt.Println(expected.CurrentX, result.CurrentX)
		t.Fail()
	}
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

	checkSpaceEqual(t, expected, result)

	if expected.Name != result.Name {
		fmt.Println(expected.Name, result.Name)
		t.Fail()
	}

	if expected.Rotation != result.Rotation {
		fmt.Println(expected.Rotation, result.Rotation)
		t.Fail()
	}

	if expected.CurrentX != result.CurrentX {
		fmt.Println(expected.CurrentX, result.CurrentX)
		t.Fail()
	}
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

	checkSpaceEqual(t, expected, result)

	if expected.Name != result.Name {
		fmt.Println(expected.Name, result.Name)
		t.Fail()
	}

	if expected.Rotation != result.Rotation {
		fmt.Println(expected.Rotation, result.Rotation)
		t.Fail()
	}

	if expected.CurrentX != result.CurrentX {
		fmt.Println(expected.CurrentX, result.CurrentX)
		t.Fail()
	}
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

	checkSpaceEqual(t, expected, result)

	if expected.Name != result.Name {
		fmt.Println(expected.Name, result.Name)
		t.Fail()
	}

	if expected.Rotation != result.Rotation {
		fmt.Println(expected.Rotation, result.Rotation)
		t.Fail()
	}

	if expected.CurrentX != result.CurrentX {
		fmt.Println(expected.CurrentX, result.CurrentX)
		t.Fail()
	}
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

	checkSpaceEqual(t, expected, result)

	if expected.Name != result.Name {
		fmt.Println(expected.Name, result.Name)
		t.Fail()
	}

	if expected.Rotation != result.Rotation {
		fmt.Println(expected.Rotation, result.Rotation)
		t.Fail()
	}

	if expected.CurrentX != result.CurrentX {
		fmt.Println(expected.CurrentX, result.CurrentX)
		t.Fail()
	}
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

	checkSpaceEqual(t, expected, result)

	if expected.Name != result.Name {
		fmt.Println(expected.Name, result.Name)
		t.Fail()
	}

	if expected.Rotation != result.Rotation {
		fmt.Println(expected.Rotation, result.Rotation)
		t.Fail()
	}

	if expected.CurrentX != result.CurrentX {
		fmt.Println(expected.CurrentX, result.CurrentX)
		t.Fail()
	}
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

	checkSpaceEqual(t, expected, result)

	if expected.Name != result.Name {
		fmt.Println(expected.Name, result.Name)
		t.Fail()
	}

	if expected.Rotation != result.Rotation {
		fmt.Println(expected.Rotation, result.Rotation)
		t.Fail()
	}

	if expected.CurrentX != result.CurrentX {
		fmt.Println(expected.CurrentX, result.CurrentX)
		t.Fail()
	}
}
