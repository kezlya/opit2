package main

import (
	"testing"
)

var fBench = Field{
	Grid: [][]bool{
		{true, true, true, false, true, true, true, true, true, false},
		{true, true, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, false, true, true, true, false},
		{true, true, true, true, true, true, false, true, true, true},
		{true, true, true, true, true, false, false, true, true, true},
		{true, false, true, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true, true, true},
		{true, true, true, true, true, false, true, true, true, true},
		{true, true, true, true, true, true, true, true, true, true},
		{true, false, true, true, true, false, false, false, false, true},
		{true, false, false, false, true, false, true, false, false, true},
		{true, false, false, true, false, true, true, false, false, false},
		{true, false, false, true, false, false, true, false, false, false},
		{true, false, false, true, false, false, false, false, false, false},
		{true, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	},
}

func Benchmark_Picks(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fBench.Picks()
	}
}

func Benchmark_FindHoles(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fBench.FindHoles()
	}
}
