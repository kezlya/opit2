package main

import (
	"fmt"
	"testing"
)

func Test_Sort(t *testing.T) {

	pos0 := Piece{Score: Score{Total: 5}}
	pos1 := Piece{Score: Score{Total: 10}}
	pos2 := Piece{Score: Score{Total: 3}}
	pos3 := Piece{Score: Score{Total: 3}}
	pos4 := Piece{Score: Score{Total: 1}}
	pos5 := Piece{Score: Score{Total: 4}}
	pos6 := Piece{Score: Score{Total: 2}}

	pp := []Piece{pos4, pos0, pos5, pos2, pos6, pos1, pos3}

	OrderedBy(SCORE).Sort(pp)

	if !(pp[0].Score.Total == 1 &&
		pp[1].Score.Total == 2 &&
		pp[2].Score.Total == 3 &&
		pp[3].Score.Total == 3 &&
		pp[4].Score.Total == 4 &&
		pp[5].Score.Total == 5) {
		t.Fail()
		for _, pos := range pp {
			fmt.Println("score:", pos.Score.Total)
		}
	}
}
