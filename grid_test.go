package main

import (
	"fmt"
	"testing"
)

func (a Grid) isEqual(b Grid) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func (g Grid) visual() {
	fmt.Println()
	for i := range g {
		y := len(g) - i - 1
		fmt.Print(y, "	")
		for _, c := range g[y] {
			if c {
				fmt.Print("⬛ ")
			} else {
				fmt.Print("⬜ ")
			}
		}
		fmt.Println()
	}
	fmt.Println("	 0 1 2 3 4 5 6 7 8 9")
}

func Test_GridFromString(t *testing.T) {
	//arrange
	raw := "0,0,0,1,1,1,0,0,0,0;" +
		"0,0,0,0,0,0,0,0,0,0;" +
		"0,0,0,0,0,0,0,0,0,0;" +
		"0,2,0,2,0,2,0,2,0,2"
	expectedGrid := Grid{
		{false, true, false, true, false, true, false, true, false, true},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}

	//act
	grid := GridFromString(raw)

	//assert
	if !grid.isEqual(expectedGrid) {
		t.Fail()
		fmt.Println("Something went wrong while conversion")
		grid.visual()
		expectedGrid.visual()
	}
}

func Test_ToField(t *testing.T) {
	//arrange
	expectedPicks := []int{0, 2, 2, 2, 2, 0, 0, 1, 0, 1}
	grid := Grid{
		{false, true, false, true, false, false, false, true, false, true},
		{false, true, true, true, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}

	//act
	field := grid.ToField()

	//assert
	if field.Width != 10 {
		t.Fail()
		fmt.Println("Bad Width", field.Width)
	}
	if field.Height != 4 {
		t.Fail()
		fmt.Println("Bad Height", field.Height)
	}
	if field.Empty != 2 {
		t.Fail()
		fmt.Println("Bad Empty", field.Empty)
	}
	if field.MaxPick != 2 {
		t.Fail()
		fmt.Println("Bad MaxY", field.MaxPick)
	}
	if !field.Picks.isEqual(expectedPicks) {
		t.Fail()
		fmt.Println("Bad Picks", field.Picks)
	}
	if field.CountBH != 1 {
		t.Fail()
		fmt.Println("Bad CountBH", field.CountBH)
	}
	if field.CountFH != 1 {
		t.Fail()
		fmt.Println("Bad CountFH", field.CountFH)
	}
}

func Test_findHoles(t *testing.T) {
	//arrange
	grid := Grid{
		{true, true, true, true, true, false, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, false, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, true},
		{true, false, true, true, true, true, false, true, true, true},
		{true, true, true, false, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, true},
		{false, false, true, true, false, true, false, false, false, false},
		{false, true, true, true, false, true, true, true, true, false},
		{false, true, true, false, false, false, false, false, false, false},
		{false, false, true, true, false, false, false, false, false, false},
		{false, false, false, true, false, false, false, false, false, false},
		{false, false, true, true, false, false, false, false, false, false},
		{false, false, true, true, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	field := grid.ToField()
	expectedBlocked := []Cell{
		Cell{X: 1, Y: 6},
		Cell{X: 1, Y: 9},
		Cell{X: 2, Y: 1},
		Cell{X: 3, Y: 7},
		Cell{X: 4, Y: 4},
		Cell{X: 5, Y: 0},
		Cell{X: 6, Y: 6},
		Cell{X: 7, Y: 2},
		Cell{X: 8, Y: 3},
		Cell{X: 6, Y: 9},
		Cell{X: 7, Y: 9},
		Cell{X: 8, Y: 9},
	}
	expectedFixable := []Cell{
		Cell{X: 2, Y: 13},
		Cell{X: 3, Y: 11},
	}
	expectedHoles := append(expectedBlocked, expectedFixable...)

	//act
	blocked, fixable := grid.findHoles(field.Picks)
	holes := append(blocked, fixable...)

	//assert
	for _, h := range holes {
		found := false
		for _, eh := range expectedHoles {
			if eh.X == h.X && eh.Y == h.Y {
				found = true
			}
		}
		if !found {
			fmt.Println("not found")
			fmt.Println("x:", h.X, " y:", h.Y)
			t.Fail()
		}
	}
}
