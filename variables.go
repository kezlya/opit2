package main

var Timebank, TimePerMove, Width, OriginalHeight, Pick, Round, CurrentPieceX, CurrentPieceY int
var Players [2]Player
var MyPlayer *Player
var CurrentPiece, NextPiece string
var IsSafePlay, IsRoofNear bool

type Player struct {
	Name   string
	Field  Field
	Points int
	Combo  int
}

type Position struct {
	Rotation     int
	X            int
	IsBurn       int
	Damage       int
	LowY         int
	HighY        int
	Score        int
	ColumnsAfter Picks
	FieldAfter   Field
}
