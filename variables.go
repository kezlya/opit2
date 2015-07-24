package main

var Timebank, TimePerMove, Width, Height, Round, CurrentPieceX, CurrentPieceY int
var Players [2]Player
var MyPlayer *Player
var CurrentPiece, NextPiece string

type Player struct {
	Name    string
	Columns []int
	Points  int
	Combo   int
}

type Position struct {
	Rotation int
	X        int
}
