package main

import (
	"testing"
)

var benchGrid = Grid{
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
}
var benchField = benchGrid.ToField()
var benchMyPlayer = Player{Field: benchField}
var benchTpiece = InitPiece("T", 3, 19)
var benchGame = Game{
	Strategy:     strategy,
	CurrentPiece: benchTpiece,
	NextPiece:    benchTpiece,
	MyPlayer:     &benchMyPlayer,
}

func Benchmark_FindPositions(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchField.FindPositions(benchTpiece)
	}
}

func Benchmark_CalculateMoves(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchGame.calculateMoves()
	}
}
