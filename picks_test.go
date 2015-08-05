package main

import (
	"testing"
)

func Test_isHoleYes(t *testing.T) {
	arrange := Picks{1, 1, 1, 1, 1, 1, 1, 2, 1, 4}
	result := arrange.IsHole()
	if !result {
		t.Fail()
	}
}

func Test_isHoleNo(t *testing.T) {
	arrange := Picks{2, 0, 2, 4, 3, 4, 2, 3, 4, 4}
	result := arrange.IsHole()
	if result {
		t.Fail()
	}
}
