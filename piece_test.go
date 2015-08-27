package main

import (
	"fmt"
	"testing"
)

func isSpaceEqual(a, b Piece) bool {
	for i, cell := range a.Space {
		if b.Space[i].X != cell.X || b.Space[i].Y != cell.Y {
			return false
		}
	}
	return true
}

func Test_Left(t *testing.T) {
	arrange := Piece{Space: [4]Cell{
		Cell{X: 3, Y: 1},
		Cell{X: 3, Y: 2},
		Cell{X: 3, Y: 3},
		Cell{X: 3, Y: 4}}}
	expected := Piece{Space: [4]Cell{
		Cell{X: 2, Y: 1},
		Cell{X: 2, Y: 2},
		Cell{X: 2, Y: 3},
		Cell{X: 2, Y: 4}}}

	result := arrange.Left()

	if !isSpaceEqual(expected, result) {
		fmt.Println(arrange)
		fmt.Println(result)
		t.Fail()
	}
}

func Test_Right(t *testing.T) {
	arrange := Piece{Space: [4]Cell{
		Cell{X: 3, Y: 1},
		Cell{X: 3, Y: 2},
		Cell{X: 3, Y: 3},
		Cell{X: 3, Y: 4}}}
	expected := Piece{Space: [4]Cell{
		Cell{X: 4, Y: 1},
		Cell{X: 4, Y: 2},
		Cell{X: 4, Y: 3},
		Cell{X: 4, Y: 4}}}

	result := arrange.Right()

	if !isSpaceEqual(expected, result) {
		fmt.Println(arrange)
		fmt.Println(result)
		t.Fail()
	}
}

func Test_Down(t *testing.T) {
	arrange := Piece{Space: [4]Cell{
		Cell{X: 3, Y: 1},
		Cell{X: 3, Y: 2},
		Cell{X: 3, Y: 3},
		Cell{X: 3, Y: 4}}}
	expected := Piece{Space: [4]Cell{
		Cell{X: 3, Y: 0},
		Cell{X: 3, Y: 1},
		Cell{X: 3, Y: 2},
		Cell{X: 3, Y: 3}}}

	result := arrange.Down()

	if !isSpaceEqual(expected, result) {
		fmt.Println(arrange)
		fmt.Println(result)
		t.Fail()
	}
}
