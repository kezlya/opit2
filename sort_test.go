package main

import (
	"fmt"
	"testing"
)

func Test_Sort(t *testing.T) {

	pos0 := Position{Damage: 4, LowY: 7, HighY: 5}
	pos1 := Position{Damage: 4, LowY: 7, HighY: 10}
	pos2 := Position{Damage: 4, LowY: 13, HighY: 3}
	pos3 := Position{Damage: 5, LowY: 1, HighY: 3}
	pos4 := Position{Damage: 5, LowY: 5, HighY: 1}
	pos5 := Position{Damage: 5, LowY: 11, HighY: 4}
	pos6 := Position{Damage: 6, LowY: 1, HighY: 2}

	positions := []Position{pos4, pos0, pos5, pos2, pos6, pos1, pos3}

	OrderedBy(DAMAGE, LOWY, HIGHY).Sort(positions)

	if !(positions[0].HighY == 5 && positions[1].HighY == 10 && positions[2].HighY == 3 && positions[3].HighY == 3 && positions[4].HighY == 1 && positions[5].HighY == 4) {
		t.Fail()
		for _, pos := range positions {
			fmt.Println("damadge:", pos.Damage, "LowY:", pos.LowY, "HighY:", pos.HighY)
		}
	}
}
