package main

import (
	"fmt"
	"testing"
)

func Test_Damage(t *testing.T) {
	before := Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, true, false, true, true, false, true, true, true, true}, {false, false, false, true, true, false, true, true, true, true}, {false, false, false, false, true, false, false, false, true, true}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	after := Field{{true, false, true, true, true, true, true, true, true, true}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, true}, {false, true, false, true, true, false, true, true, true, true}, {false, false, false, true, true, true, true, true, true, true}, {false, false, false, false, true, true, true, false, true, true}, {false, false, false, false, false, false, true, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}

	piksBefore := before.Picks()
	piksAfter := after.Picks()

	damage, lowY, highY, hole := piksBefore.Damage(piksAfter)

	if damage != 5 {
		fmt.Println("damage is wrong", damage)
		t.Fail()
	}

	if lowY != 3 {
		fmt.Println("lowY is wrong", lowY)
		t.Fail()
	}

	if highY != 7 {
		fmt.Println("highY is wrong", highY)
		t.Fail()
	}

	if hole != 2 {
		fmt.Println("hole is wrong", hole)
		t.Fail()
	}
}
