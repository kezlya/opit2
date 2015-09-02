package main

import (
	"fmt"
	"testing"
)

func PrintVisual(f Field) {
	fmt.Println()
	y := len(f) - 1
	for i := range f {
		fmt.Print(y-i, "	")
		for _, c := range f[y-i] {
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

func PrintVisuals(a, b Field) {
	fmt.Println()
	y := len(a) - 1
	for i := range a {
		fmt.Print(y-i, "	")
		for _, c := range a[y-i] {
			if c {
				fmt.Print("⬛ ")
			} else {
				fmt.Print("⬜ ")
			}
		}
		fmt.Print("   ")
		for _, c := range b[y-i] {
			if c {
				fmt.Print("⬛ ")
			} else {
				fmt.Print("⬜ ")
			}
		}
		fmt.Println()
	}
	fmt.Println("	 0 1 2 3 4 5 6 7 8 9    0 1 2 3 4 5 6 7 8 9")
}

func Test_convertField(t *testing.T) {
	cleanInput := "0,0,0,1,1,1,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,2,0,0,0;0,0,0,0,0,0,2,0,0,0;0,0,0,0,0,0,2,2,0,0;0,0,2,2,0,0,2,2,0,0;0,2,2,2,0,0,2,2,0,2;2,2,2,2,2,0,2,2,2,2;2,2,0,2,2,2,2,2,2,2;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0"
	expect := Field{{true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, false}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, false, false, true, true, false, true}, {false, false, true, true, false, false, true, true, false, false}, {false, false, false, false, false, false, true, true, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	var result Field
	result = result.init(cleanInput)

	if !expect.Equal(result) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], result[y-i])
		}
	}
}

func Test_IsBurn(t *testing.T) {
	arange := Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, false, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, false, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expect := 5

	result := arange.WillBurn()
	if expect != result {
		t.Fail()
		fmt.Println(expect, "!=", result)
	}
}

func Test_Burn(t *testing.T) {
	arange := Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, false, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, false, true, true, false, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expect := Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, false, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, false, true, true, false, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, false, true}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, true, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	arange.Burn()

	if !arange.Equal(expect) {
		t.Fail()
		y := len(expect) - 1
		for i := range expect {
			fmt.Println(expect[y-i], arange[y-i])
		}
	}
}

func Test_Picks(t *testing.T) {
	arrange := Field{{true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, false, true, true, true, true, true, true, false, true}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	expect := Picks{12, 13, 14, 14, 14, 17, 14, 14, 13, 14}

	result := arrange.Picks()

	if !expect.Equal(result) {
		fmt.Println(result)
		fmt.Println(expect)
		t.Fail()
	}
}

func Test_FindHoles(t *testing.T) {
	arrange := Field{{true, true, true, true, true, false, true, true, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, false, true, true, true, true, false, true, true, true}, {true, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, false, true, true, false, true, false, false, false, false}, {false, true, true, true, false, true, true, true, true, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, false, true, false, false, false, false, false, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
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
	}
	expectedFixable := []Cell{
		Cell{X: 2, Y: 13},
		Cell{X: 8, Y: 9},
		Cell{X: 3, Y: 11},
		Cell{X: 6, Y: 9},
		Cell{X: 7, Y: 9},
	}
	expectedHoles := append(expectedBlocked, expectedFixable...)

	blocked, fixable := arrange.FindHoles(arrange.Picks())
	holes := append(blocked, fixable...)

	if len(blocked) != len(expectedBlocked) {
		fmt.Println("blocked: ", len(blocked), len(expectedBlocked))
		PrintVisual(arrange)
		t.Fail()
	}

	if len(fixable) != len(expectedFixable) {
		fmt.Println("fixable: ", len(fixable), len(expectedFixable))
		PrintVisual(arrange)
		t.Fail()
	}

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

func Test_IsValid(t *testing.T) {
	arrange := Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, true, true, true, false, true, true, true, true, false}, {false, true, true, false, false, false, false, true, false, false}, {false, false, true, true, false, false, false, true, false, false}, {false, false, false, true, false, false, false, true, false, false}, {false, false, true, true, false, false, true, true, true, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	good1 := Cell{X: 2, Y: 6}
	good2 := Cell{X: 5, Y: 17}
	good3 := Cell{X: 0, Y: 17}
	good4 := Cell{X: 1, Y: 0}
	bad1 := Cell{X: 3, Y: 3}
	bad2 := Cell{X: -1, Y: 3}
	bad3 := Cell{X: 3, Y: -1}
	bad4 := Cell{X: 2, Y: 18}

	if !arrange.IsValid(&map[string]Cell{"x": good1}) {
		fmt.Println("should be valid", good1)
		t.Fail()
	}

	if !arrange.IsValid(&map[string]Cell{"x": good2}) {
		fmt.Println("should be valid", good2)
		t.Fail()
	}

	if !arrange.IsValid(&map[string]Cell{"x": good3}) {
		fmt.Println("should be valid", good3)
		t.Fail()
	}

	if !arrange.IsValid(&map[string]Cell{"x": good4}) {
		fmt.Println("should be valid", good4)
		t.Fail()
	}

	if arrange.IsValid(&map[string]Cell{"x": bad1}) {
		fmt.Println("should not be valid", bad1)
		t.Fail()
	}

	if arrange.IsValid(&map[string]Cell{"x": bad2}) {
		fmt.Println("should not be valid", bad2)
		t.Fail()
	}

	if arrange.IsValid(&map[string]Cell{"x": bad3}) {
		fmt.Println("should not be valid", bad3)
		t.Fail()
	}

	if arrange.IsValid(&map[string]Cell{"x": bad4}) {
		fmt.Println("should not be valid", bad4)
		t.Fail()
	}

	if arrange.IsValid(&map[string]Cell{"1": good1, "2": good2, "3": good3, "4": good4, "5": bad1, "6": bad2, "7": bad3, "8": bad4}) {
		fmt.Println("should not be valid")
		t.Fail()
	}
}

func Test_FixHoles_T(t *testing.T) {
	var arangePathField = Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, true, true, true, false, true, true, true, true, false}, {false, true, true, false, false, false, false, true, false, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, false, true, false, false, false, true, false, false}, {false, false, true, true, false, false, true, true, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	piece := Piece{Name: "T", Rotation: 0}
	piece.InitSpace(Cell{X: 3, Y: 19})

	hole0 := Cell{X: 3, Y: 16}
	hole1 := Cell{X: 2, Y: 13}
	hole2 := Cell{X: 3, Y: 11}
	hole3 := Cell{X: 7, Y: 12}
	hole_bad1 := Cell{X: 2, Y: 6}
	hole_bad2 := Cell{X: 8, Y: 11}

	good_positions := arangePathField.FixHoles(piece, []Cell{hole0, hole1, hole2, hole3})
	bad_positions := arangePathField.FixHoles(piece, []Cell{hole_bad1, hole_bad2})

	if len(good_positions) != 12 {
		for _, pos := range good_positions {
			fmt.Println("good positions failed")
			fmt.Println(pos.Moves)
			PrintVisual(arangePathField)
		}
		t.Fail()
	}

	if len(bad_positions) != 0 {
		for _, pos := range bad_positions {
			fmt.Println("bad positions failed")
			fmt.Println(pos.Moves)
			PrintVisual(arangePathField)
		}
		t.Fail()
	}
}

func Test_ValidatePosition_I(t *testing.T) {
	var arangeField = Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, true, true, true, false, true, true, true, true, false}, {false, true, true, false, false, false, false, true, false, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, false, true, false, false, false, true, false, false}, {false, false, true, true, false, false, true, true, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	piece := Piece{Name: "I", Rotation: 0}
	piece.InitSpace(Cell{X: 3, Y: 19})

	validPieces := arangeField.ValidPosition(piece)

	if len(validPieces) != 13 {
		PrintVisual(arangeField)
		fmt.Println("should be 13 but returned", len(validPieces))
		t.Fail()
	}
}

func Test_ValidatePosition_T(t *testing.T) {
	var arangeField = Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, true, true, true, false, true, true, true, true, false}, {false, true, true, false, false, false, false, true, false, false}, {false, false, true, true, false, false, false, false, false, false}, {false, false, false, true, false, false, false, true, false, false}, {false, false, true, true, false, false, true, true, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, true, false}, {false, false, true, true, false, false, false, false, true, false}, {false, false, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	piece := Piece{Name: "T", Rotation: 0}
	piece.InitSpace(Cell{X: 3, Y: 19})

	validPieces := arangeField.ValidPosition(piece)

	if len(validPieces) != 23 {
		PrintVisual(arangeField)
		/*for _, p := range validPieces {
			fmt.Println(p.Rotation, p.CurrentX)
		}*/

		fmt.Println("should be 23 but returned", len(validPieces))
		t.Fail()
	}
}
