package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testGrid = Grid{
	{true, false, true, true, true, true, true, true, true, true},
	{true, true, false, true, true, false, true, true, true, true},
	{true, true, true, true, true, true, true, false, true, true},
	{true, true, true, true, true, true, true, false, true, true},
	{true, true, true, true, true, true, true, true, true, false},
	{true, true, true, true, false, true, true, false, true, true},
	{false, true, true, true, true, true, true, true, true, true},
	{true, true, true, true, false, true, true, true, true, true},
	{true, true, true, true, true, true, true, true, true, false},
	{true, true, true, true, true, true, true, false, true, true},
	{true, true, false, false, false, false, true, true, true, true},
	{true, true, true, false, false, false, true, true, true, true},
	{false, true, true, false, false, true, true, true, true, true},
	{false, true, true, false, false, false, true, true, false, true},
	{false, false, true, false, false, false, true, false, false, false},
	{false, false, true, false, false, false, true, false, false, false},
	{false, false, true, false, false, false, true, false, false, false},
	{false, false, true, false, false, false, true, false, false, false},
	{false, false, true, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
}

func (a Grid) assertEqualTo(b Grid, t *testing.T) {
	if !a.isEqual(b) {
		t.Fail()
		fmt.Println("Grids are not equal")
		a.visual()
		b.visual()
	}
}

func (a Grid) assertNotEqualTo(b Grid, t *testing.T) {
	if a.isEqual(b) {
		t.Fail()
		fmt.Println("Grids are equal")
		a.visual()
		b.visual()
	}
}

func (a Grid) isEqual(b Grid) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[i]); j++ {
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

func Test_InitGrid(t *testing.T) {
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
	grid := InitGrid(raw)

	//assert
	grid.assertEqualTo(expectedGrid, t)
}

func Test_Copy(t *testing.T) {
	//arrange

	//act
	newGrid := testGrid.Copy()

	//assert
	testGrid.assertEqualTo(newGrid, t)
	tg := reflect.ValueOf(testGrid).Pointer()
	ng := reflect.ValueOf(newGrid).Pointer()
	if tg == ng {
		t.Fail()
		fmt.Println("Grid was not copied")
		fmt.Println("Grid pointers", tg, "and", ng, "should be different")
	}
}

func Test_ToField(t *testing.T) {
	//arrange
	//expectedPicks := []int{0, 2, 2, 2, 2, 0, 0, 1, 0, 1}
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
	if field.CountBH != 1 {
		t.Fail()
		fmt.Println("Bad CountBH", field.CountBH)
	}
	if field.CountFH != 1 {
		t.Fail()
		fmt.Println("Bad CountFH", field.CountFH)
	}
}

func Test_Burn(t *testing.T) {
	//arrange
	grid := Grid{
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, false, true, true, false, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, true, true},
		{true, true, true, true, false, true, true, false, true, true},
		{true, true, true, true, true, true, true, true, true, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, true, true, true, false, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, false, true},
		{false, false, true, false, false, false, true, false, false, false},
		{false, false, true, false, false, false, true, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	expectedGrid := Grid{
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, false, true, true, false, true, true, true, true},
		{true, true, true, true, true, true, true, false, true, true},
		{true, true, true, true, false, true, true, false, true, true},
		{true, true, true, true, false, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, true, true, true, false, true, true, true, true},
		{false, true, true, true, true, true, true, true, true, true},
		{false, true, true, true, true, true, true, true, false, true},
		{false, false, true, false, false, false, true, false, false, false},
		{false, false, true, false, false, false, true, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}

	//ast
	result := grid.burn()

	//assert
	grid.assertEqualTo(expectedGrid, t)
	if result != 5 {
		t.Fail()
		fmt.Println("Bad Burn", result)
		grid.visual()
		expectedGrid.visual()
	}
}

func Test_isTshapeSpace(t *testing.T) {
	//arrange
	hole := &Cell{X: 3, Y: 0}
	grid := Grid{
		{false, true, true, false, true, true, true, false, false, false},
		{false, false, false, false, false, false, true, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}

	//act
	result := grid.isTshapeSpace(hole)

	//assert
	if !result {
		t.Fail()
		fmt.Println("Tshap not found")
		grid.visual()
	}
}

func Test_tSpinLevels(t *testing.T) {
	//arrange
	grid := Grid{
		{true, true, true, false, true, true, true, true, true, true},
		{false, false, false, false, false, false, true, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	field := grid.ToField()

	//act
	level1, level2, level3 := grid.tSpinLevels(field.MaxPick)

	//assert
	if level1 != 1 {
		t.Fail()
		fmt.Println("Level1 tSpace faild", level1)
		grid.visual()
	}
	if level2 != 0 {
		t.Fail()
		fmt.Println("Level1 tSpace faild", level2)
		grid.visual()
	}
	if level3 != 0 {
		t.Fail()
		fmt.Println("Level1 tSpace faild", level3)
		grid.visual()
	}
}
