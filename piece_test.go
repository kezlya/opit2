package main

import (
	"fmt"
	"testing"
)

var t1 = Cell{X: 1, Y: 13}
var t2 = Cell{X: 2, Y: 13}
var t3 = Cell{X: 3, Y: 13}
var t4 = Cell{X: 4, Y: 13}

var m1 = Cell{X: 1, Y: 12}
var m2 = Cell{X: 2, Y: 12}
var m3 = Cell{X: 3, Y: 12}
var m4 = Cell{X: 4, Y: 12}

var b1 = Cell{X: 1, Y: 11}
var b2 = Cell{X: 2, Y: 11}
var b3 = Cell{X: 3, Y: 11}
var b4 = Cell{X: 4, Y: 11}

var x1 = Cell{X: 1, Y: 10}
var x2 = Cell{X: 2, Y: 10}
var x3 = Cell{X: 3, Y: 10}
var x4 = Cell{X: 4, Y: 10}

var I0 = map[string]Cell{"m1": m1, "m2": m2, "m3": m3, "m4": m4}
var I1 = map[string]Cell{"t3": t3, "m3": m3, "b3": b3, "x3": x3}
var I2 = map[string]Cell{"b1": b1, "b2": b2, "b3": b3, "b4": b4}
var I3 = map[string]Cell{"t2": t2, "m2": m2, "b2": b2, "x2": x2}

var J0 = map[string]Cell{"t1": t1, "m1": m1, "m2": m2, "m3": m3}
var J1 = map[string]Cell{"t3": t3, "t2": t2, "m2": m2, "b2": b2}
var J2 = map[string]Cell{"b3": b3, "m1": m1, "m2": m2, "m3": m3}
var J3 = map[string]Cell{"b1": b1, "t2": t2, "m2": m2, "b2": b2}

var L0 = map[string]Cell{"t3": t3, "m1": m1, "m2": m2, "m3": m3}
var L1 = map[string]Cell{"b3": b3, "t2": t2, "m2": m2, "b2": b2}
var L2 = map[string]Cell{"b1": b1, "m1": m1, "m2": m2, "m3": m3}
var L3 = map[string]Cell{"t1": t1, "t2": t2, "m2": m2, "b2": b2}

var T0 = map[string]Cell{"t2": t2, "m1": m1, "m2": m2, "m3": m3}
var T1 = map[string]Cell{"m3": m3, "t2": t2, "m2": m2, "b2": b2}
var T2 = map[string]Cell{"b2": b2, "m1": m1, "m2": m2, "m3": m3}
var T3 = map[string]Cell{"m1": m1, "t2": t2, "m2": m2, "b2": b2}

var S0 = map[string]Cell{"m1": m1, "m2": m2, "t2": t2, "t3": t3}
var S1 = map[string]Cell{"t2": t2, "m2": m2, "m3": m3, "b3": b3}
var S2 = map[string]Cell{"m2": m2, "m3": m3, "b1": b1, "b2": b2}
var S3 = map[string]Cell{"t1": t1, "m1": m1, "m2": m2, "b2": b2}

var Z0 = map[string]Cell{"t2": t2, "t1": t1, "m2": m2, "m3": m3}
var Z1 = map[string]Cell{"m3": m3, "t3": t3, "m2": m2, "b2": b2}
var Z2 = map[string]Cell{"b2": b2, "m1": m1, "m2": m2, "b3": b3}
var Z3 = map[string]Cell{"m1": m1, "t2": t2, "m2": m2, "b1": b1}

func pieceAssert(t *testing.T, a, b *Piece) {
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

	if a.CurrentY != b.CurrentY {
		fmt.Println(a.CurrentY, b.CurrentY)
		t.Fail()
	}
}

func Test_Left(t *testing.T) {
	arrange := Piece{Space: I1, Name: I, Rotation: 1, CurrentX: 3, CurrentY: 10}
	ex_space := make(map[string]Cell, 4)
	ex_space["t3"] = Cell{X: 2, Y: 13}
	ex_space["m3"] = Cell{X: 2, Y: 12}
	ex_space["b3"] = Cell{X: 2, Y: 11}
	ex_space["x3"] = Cell{X: 2, Y: 10}
	expected := &Piece{Space: ex_space, Name: I, Rotation: 1, CurrentX: 2, CurrentY: 10}
	result := arrange.Left()
	pieceAssert(t, expected, result)
}

func Test_Right(t *testing.T) {
	arrange := Piece{Space: I1, Name: I, Rotation: 1, CurrentX: 3, CurrentY: 1}
	ex_space := make(map[string]Cell, 4)
	ex_space["t3"] = Cell{X: 4, Y: 13}
	ex_space["m3"] = Cell{X: 4, Y: 12}
	ex_space["b3"] = Cell{X: 4, Y: 11}
	ex_space["x3"] = Cell{X: 4, Y: 10}
	expected := &Piece{Space: ex_space, Name: I, Rotation: 1, CurrentX: 4, CurrentY: 1}
	result := arrange.Right()
	pieceAssert(t, expected, result)
}

func Test_Down(t *testing.T) {
	arrange := Piece{Space: I1, Name: I, Rotation: 1, CurrentX: 3, CurrentY: 10}
	ex_space := make(map[string]Cell, 4)
	ex_space["t3"] = Cell{X: 3, Y: 12}
	ex_space["m3"] = Cell{X: 3, Y: 11}
	ex_space["b3"] = Cell{X: 3, Y: 10}
	ex_space["x3"] = Cell{X: 3, Y: 9}
	expected := &Piece{Space: ex_space, Name: I, Rotation: 1, CurrentX: 3, CurrentY: 9}
	result := arrange.Down()
	pieceAssert(t, expected, result)
}

func Test_Drop(t *testing.T) {
	space := make(map[string]Cell, 4)
	space["t2"] = Cell{X: 2, Y: 20}
	space["m1"] = Cell{X: 1, Y: 19}
	space["m2"] = Cell{X: 2, Y: 19}
	space["m3"] = Cell{X: 3, Y: 19}
	arrange := Piece{Space: space, Name: T, Rotation: 0, CurrentX: 1, CurrentY: 19}

	ex_space := make(map[string]Cell, 4)
	ex_space["t2"] = Cell{X: 2, Y: 8}
	ex_space["m1"] = Cell{X: 1, Y: 7}
	ex_space["m2"] = Cell{X: 2, Y: 7}
	ex_space["m3"] = Cell{X: 3, Y: 7}
	expected := &Piece{Space: ex_space, Name: T, Rotation: 0, CurrentX: 1, CurrentY: 7}
	result := arrange.Drop(12)
	pieceAssert(t, expected, result)
}

func Test_TurnRight_I0(t *testing.T) {
	arrange := Piece{Space: I0, Name: I, Rotation: 0, CurrentX: 1, CurrentY: 12}
	expected := &Piece{Space: I1, Name: I, Rotation: 1, CurrentX: 3, CurrentY: 10}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_I1(t *testing.T) {
	arrange := Piece{Space: I1, Name: I, Rotation: 1, CurrentX: 3, CurrentY: 10}
	expected := &Piece{Space: I2, Name: I, Rotation: 2, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_I2(t *testing.T) {
	arrange := Piece{Space: I2, Name: I, Rotation: 2, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: I3, Name: I, Rotation: 3, CurrentX: 2, CurrentY: 10}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_I3(t *testing.T) {
	arrange := Piece{Space: I3, Name: I, Rotation: 3, CurrentX: 2, CurrentY: 10}
	expected := &Piece{Space: I0, Name: I, Rotation: 0, CurrentX: 1, CurrentY: 12}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_I0(t *testing.T) {
	arrange := Piece{Space: I0, Name: I, Rotation: 0, CurrentX: 1, CurrentY: 12}
	expected := &Piece{Space: I3, Name: I, Rotation: 3, CurrentX: 2, CurrentY: 10}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_I3(t *testing.T) {
	arrange := Piece{Space: I3, Name: I, Rotation: 3, CurrentX: 2, CurrentY: 10}
	expected := &Piece{Space: I2, Name: I, Rotation: 2, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_I2(t *testing.T) {
	arrange := Piece{Space: I2, Name: I, Rotation: 2, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: I1, Name: I, Rotation: 1, CurrentX: 3, CurrentY: 10}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_I1(t *testing.T) {
	arrange := Piece{Space: I1, Name: I, Rotation: 1, CurrentX: 3, CurrentY: 10}
	expected := &Piece{Space: I0, Name: I, Rotation: 0, CurrentX: 1, CurrentY: 12}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_J0(t *testing.T) {
	arrange := Piece{Space: J0, Name: J, Rotation: 0, CurrentX: 1, CurrentY: 12}
	expected := &Piece{Space: J1, Name: J, Rotation: 1, CurrentX: 2, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_J1(t *testing.T) {
	arrange := Piece{Space: J1, Name: J, Rotation: 1, CurrentX: 2, CurrentY: 11}
	expected := &Piece{Space: J2, Name: J, Rotation: 2, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_J2(t *testing.T) {
	arrange := Piece{Space: J2, Name: J, Rotation: 2, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: J3, Name: J, Rotation: 3, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_J3(t *testing.T) {
	arrange := Piece{Space: J3, Name: J, Rotation: 3, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: J0, Name: J, Rotation: 0, CurrentX: 1, CurrentY: 12}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_J0(t *testing.T) {
	arrange := Piece{Space: J0, Name: J, Rotation: 0, CurrentX: 1, CurrentY: 12}
	expected := &Piece{Space: J3, Name: J, Rotation: 3, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_J3(t *testing.T) {
	arrange := Piece{Space: J3, Name: J, Rotation: 3, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: J2, Name: J, Rotation: 2, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_J2(t *testing.T) {
	arrange := Piece{Space: J2, Name: J, Rotation: 2, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: J1, Name: J, Rotation: 1, CurrentX: 2, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_J1(t *testing.T) {
	arrange := Piece{Space: J1, Name: J, Rotation: 1, CurrentX: 2, CurrentY: 11}
	expected := &Piece{Space: J0, Name: J, Rotation: 0, CurrentX: 1, CurrentY: 12}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_L0(t *testing.T) {
	arrange := Piece{Space: L0, Name: L, Rotation: 0, CurrentX: 1, CurrentY: 12}
	expected := &Piece{Space: L1, Name: L, Rotation: 1, CurrentX: 2, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_L1(t *testing.T) {
	arrange := Piece{Space: L1, Name: L, Rotation: 1, CurrentX: 2, CurrentY: 11}
	expected := &Piece{Space: L2, Name: L, Rotation: 2, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_L2(t *testing.T) {
	arrange := Piece{Space: L2, Name: L, Rotation: 2, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: L3, Name: L, Rotation: 3, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_L3(t *testing.T) {
	arrange := Piece{Space: L3, Name: L, Rotation: 3, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: L0, Name: L, Rotation: 0, CurrentX: 1, CurrentY: 12}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_L0(t *testing.T) {
	arrange := Piece{Space: L0, Name: L, Rotation: 0, CurrentX: 1, CurrentY: 12}
	expected := &Piece{Space: L3, Name: L, Rotation: 3, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_L3(t *testing.T) {
	arrange := Piece{Space: L3, Name: L, Rotation: 3, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: L2, Name: L, Rotation: 2, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_L2(t *testing.T) {
	arrange := Piece{Space: L2, Name: L, Rotation: 2, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: L1, Name: L, Rotation: 1, CurrentX: 2, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_L1(t *testing.T) {
	arrange := Piece{Space: L1, Name: L, Rotation: 1, CurrentX: 2, CurrentY: 11}
	expected := &Piece{Space: L0, Name: L, Rotation: 0, CurrentX: 1, CurrentY: 12}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_T0(t *testing.T) {
	arrange := Piece{Space: T0, Name: T, Rotation: 0, CurrentX: 1, CurrentY: 12}
	expected := &Piece{Space: T1, Name: T, Rotation: 1, CurrentX: 2, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_T1(t *testing.T) {
	arrange := Piece{Space: T1, Name: T, Rotation: 1, CurrentX: 2, CurrentY: 11}
	expected := &Piece{Space: T2, Name: T, Rotation: 2, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_T2(t *testing.T) {
	arrange := Piece{Space: T2, Name: T, Rotation: 2, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: T3, Name: T, Rotation: 3, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_T3(t *testing.T) {
	arrange := Piece{Space: T3, Name: T, Rotation: 3, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: T0, Name: T, Rotation: 0, CurrentX: 1, CurrentY: 12}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_T0(t *testing.T) {
	arrange := Piece{Space: T0, Name: T, Rotation: 0, CurrentX: 1, CurrentY: 12}
	expected := &Piece{Space: T3, Name: T, Rotation: 3, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_T3(t *testing.T) {
	arrange := Piece{Space: T3, Name: T, Rotation: 3, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: T2, Name: T, Rotation: 2, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_T2(t *testing.T) {
	arrange := Piece{Space: T2, Name: T, Rotation: 2, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: T1, Name: T, Rotation: 1, CurrentX: 2, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_T1(t *testing.T) {
	arrange := Piece{Space: T1, Name: T, Rotation: 1, CurrentX: 2, CurrentY: 11}
	expected := &Piece{Space: T0, Name: T, Rotation: 0, CurrentX: 1, CurrentY: 12}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_S0(t *testing.T) {
	arrange := Piece{Space: S0, Name: S, Rotation: 0, CurrentX: 1, CurrentY: 12}
	expected := &Piece{Space: S1, Name: S, Rotation: 1, CurrentX: 2, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_S1(t *testing.T) {
	arrange := Piece{Space: S1, Name: S, Rotation: 1, CurrentX: 2, CurrentY: 11}
	expected := &Piece{Space: S2, Name: S, Rotation: 2, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_S2(t *testing.T) {
	arrange := Piece{Space: S2, Name: S, Rotation: 2, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: S3, Name: S, Rotation: 3, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_S3(t *testing.T) {
	arrange := Piece{Space: S3, Name: S, Rotation: 3, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: S0, Name: S, Rotation: 0, CurrentX: 1, CurrentY: 12}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_S0(t *testing.T) {
	arrange := Piece{Space: S0, Name: S, Rotation: 0, CurrentX: 1, CurrentY: 12}
	expected := &Piece{Space: S3, Name: S, Rotation: 3, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_S3(t *testing.T) {
	arrange := Piece{Space: S3, Name: S, Rotation: 3, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: S2, Name: S, Rotation: 2, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_S2(t *testing.T) {
	arrange := Piece{Space: S2, Name: S, Rotation: 2, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: S1, Name: S, Rotation: 1, CurrentX: 2, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_S1(t *testing.T) {
	arrange := Piece{Space: S1, Name: S, Rotation: 1, CurrentX: 2, CurrentY: 11}
	expected := &Piece{Space: S0, Name: S, Rotation: 0, CurrentX: 1, CurrentY: 12}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_Z0(t *testing.T) {
	arrange := Piece{Space: Z0, Name: Z, Rotation: 0, CurrentX: 1, CurrentY: 12}
	expected := &Piece{Space: Z1, Name: Z, Rotation: 1, CurrentX: 2, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_Z1(t *testing.T) {
	arrange := Piece{Space: Z1, Name: Z, Rotation: 1, CurrentX: 2, CurrentY: 11}
	expected := &Piece{Space: Z2, Name: Z, Rotation: 2, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_Z2(t *testing.T) {
	arrange := Piece{Space: Z2, Name: Z, Rotation: 2, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: Z3, Name: Z, Rotation: 3, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnRight_Z3(t *testing.T) {
	arrange := Piece{Space: Z3, Name: Z, Rotation: 3, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: Z0, Name: Z, Rotation: 0, CurrentX: 1, CurrentY: 12}
	result := arrange.Turnright()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_Z0(t *testing.T) {
	arrange := Piece{Space: Z0, Name: Z, Rotation: 0, CurrentX: 1, CurrentY: 12}
	expected := &Piece{Space: Z3, Name: Z, Rotation: 3, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_Z3(t *testing.T) {
	arrange := Piece{Space: Z3, Name: Z, Rotation: 3, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: Z2, Name: Z, Rotation: 2, CurrentX: 1, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_Z2(t *testing.T) {
	arrange := Piece{Space: Z2, Name: Z, Rotation: 2, CurrentX: 1, CurrentY: 11}
	expected := &Piece{Space: Z1, Name: Z, Rotation: 1, CurrentX: 2, CurrentY: 11}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_TurnLeft_Z1(t *testing.T) {
	arrange := Piece{Space: Z1, Name: Z, Rotation: 1, CurrentX: 2, CurrentY: 11}
	expected := &Piece{Space: Z0, Name: Z, Rotation: 0, CurrentX: 1, CurrentY: 12}
	result := arrange.Turnleft()
	pieceAssert(t, expected, result)
}

func Test_isPerfectClear(t *testing.T) {
	//arrange
	clearField := EmptyGrig10x20.Copy().ToField()
	notclearField := EmptyGrig10x20.Copy().ToField()
	notclearField.Grid[7][7] = true
	p1 := Piece{FieldAfter: &notclearField}
	p2 := Piece{FieldAfter: &clearField}

	//act
	good := p1.isPerfectClear()
	notgood := p2.isPerfectClear()

	//assert
	if good {
		t.Fail()
		fmt.Println("should not be perfect clear")
	}

	if !notgood {
		t.Fail()
		fmt.Println("should be perfect clear")
	}
}

func Test_TSpins_Single(t *testing.T) {
	//arrange
	badGrid := Grid{
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, false, true, true, true, false, false, false, true, true},
		{false, false, true, true, false, false, false, false, true, false},
		{false, false, false, true, false, false, false, false, false, false},
		{false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	goodGrid := Grid{
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, false, true, true, true, false, false, false, true, true},
		{false, false, true, true, true, true, false, false, true, false},
		{false, false, false, true, false, false, false, false, false, false},
		{false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	p := Piece{
		Name:     T,
		Space:    T2,
		Rotation: 2,
		Score:    &Score{},
	}
	p.Space["m1"] = Cell{X: 5, Y: 7}
	p.Space["m2"] = Cell{X: 6, Y: 7}
	p.Space["b2"] = Cell{X: 6, Y: 6}
	p.Space["m3"] = Cell{X: 7, Y: 7}
	badField := badGrid.ApplyPiece(p.Space).ToField()
	goodField := goodGrid.ApplyPiece(p.Space).ToField()

	//act
	p.assignField(&badField)
	badSpin := p.Tspin
	p.assignField(&goodField)
	goodSpin := p.Tspin

	//assert
	if badSpin {
		t.Fail()
		badField.Grid.visual()
		fmt.Println("should not be single T spin")
	}
	if !goodSpin {
		t.Fail()
		goodField.Grid.visual()
		fmt.Println("should be single T spin")
	}
}

func Test_TSpins_Double(t *testing.T) {
	//arrange
	badGrid := Grid{
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, true, true, true, true, false, false, false, true, true},
		{false, false, true, true, true, false, false, false, true, false},
		{false, false, false, true, false, false, false, false, false, false},
		{false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	goodGrid := Grid{
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, true, false},
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, false, true, true, true},
		{true, true, true, true, true, false, false, false, true, true},
		{false, false, true, true, true, true, false, false, true, false},
		{false, false, false, true, false, false, false, false, false, false},
		{false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	p := Piece{
		Name:     T,
		Space:    T2,
		Rotation: 2,
		Score:    &Score{},
	}
	p.Space["m1"] = Cell{X: 5, Y: 7}
	p.Space["m2"] = Cell{X: 6, Y: 7}
	p.Space["b2"] = Cell{X: 6, Y: 6}
	p.Space["m3"] = Cell{X: 7, Y: 7}
	badField := badGrid.ApplyPiece(p.Space).ToField()
	goodField := goodGrid.ApplyPiece(p.Space).ToField()

	//act
	p.assignField(&badField)
	badSpin := p.Tspin2
	p.assignField(&goodField)
	goodSpin := p.Tspin2

	//assert
	if badSpin {
		t.Fail()
		badField.Grid.visual()
		fmt.Println("should not be double T spin")
	}
	if !goodSpin {
		t.Fail()
		goodField.Grid.visual()
		fmt.Println("should be double T spin")
	}
}

/*
func Test_setStep(t *testing.T) {
	empty := Field{{false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	fieldAfter := Field{{false, false, true, true, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	L1["b3"] = Cell{X: 3, Y: 0}
	L1["t2"] = Cell{X: 2, Y: 2}
	L1["m2"] = Cell{X: 2, Y: 1}
	L1["b2"] = Cell{X: 2, Y: 0}
	arrangePiece := Piece{Name: L, CurrentX: 2, CurrentY: 0, Rotation: 1, Space: L1}
	arrangePiece.FieldAfter = fieldAfter
	arrangePiece.setStep(empty.Picks())

	if arrangePiece.Score.Step != 4 {
		t.Fail()
		fmt.Println("score should be: 4 but actual ", arrangePiece.Score.Step)
	}

}
*/
