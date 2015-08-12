package main

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
