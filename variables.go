package main

var Timebank, TimePerMove, Width, OriginalHeight, Height, Pick, Round, CurrentPieceX, CurrentPieceY int
var Players [2]Player
var MyPlayer *Player
var CurrentPiece, NextPiece string

type Player struct {
	Name    string
	Columns []int
	MaxY	    int
	Field   [][]bool
	Points  int
	Combo   int
}

type Position struct {
	Rotation     int
	X            int
	IsBurn       int
	Damadge      int
	MaxY			 int
	ColumnsAfter []int
	FieldAfter   [][]bool
}
