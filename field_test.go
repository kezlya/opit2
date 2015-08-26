package main

import (
	"fmt"
	"testing"
)

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
	expectedBlocked := []Hole{
		Hole{X: 1, Y: 6},
		Hole{X: 1, Y: 9},
		Hole{X: 2, Y: 1},
		Hole{X: 3, Y: 7},
		Hole{X: 4, Y: 4},
		Hole{X: 5, Y: 0},
		Hole{X: 6, Y: 6},
		Hole{X: 7, Y: 2},
		Hole{X: 8, Y: 3},
	}

	expectedLeft := []Hole{
		Hole{X: 2, Y: 13},
		Hole{X: 8, Y: 9},
	}

	expectedRight := []Hole{
		Hole{X: 3, Y: 11},
		Hole{X: 6, Y: 9},
		Hole{X: 7, Y: 9},
	}
	expectedHoles := append(expectedBlocked, expectedLeft...)
	expectedHoles = append(expectedHoles, expectedRight...)

	blocked, left, right := arrange.FindHoles(arrange.Picks())
	holes := append(blocked, left...)
	holes = append(holes, right...)

	if len(blocked) != len(expectedBlocked) {
		fmt.Println("blocked: ", len(blocked), len(expectedBlocked))
		t.Fail()
	}

	if len(left) != len(expectedLeft) {
		fmt.Println("left: ", len(left), len(expectedLeft))
		t.Fail()
	}

	if len(right) != len(expectedRight) {
		fmt.Println("right: ", len(right), len(expectedRight))
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
