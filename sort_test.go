package main

import (
	"fmt"
	"testing"
)

func Test_Sort(t *testing.T) {
	//arrange
	pos0 := &Piece{Score: &Score{Total: 5}, FieldAfter: &testField}
	pos1 := &Piece{Score: &Score{Total: 10}, FieldAfter: &testField}
	pos2 := &Piece{Score: &Score{Total: 3}, FieldAfter: &testField}
	pos3 := &Piece{Score: &Score{Total: 3}, FieldAfter: &testField}
	pos4 := &Piece{Score: &Score{Total: 1}, FieldAfter: &testField}
	pos5 := &Piece{Score: &Score{Total: 4}, FieldAfter: &testField}
	pos6 := &Piece{Score: &Score{Total: 2}, FieldAfter: &testField}
	pp := []*Piece{pos4, pos0, pos5, pos2, pos6, pos1, pos3}

	//act
	orderedBy(SCORE).Sort(pp)

	//arrange
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
