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
	Burn         int
	Hole         int
	Damage       int
	LowY         int
	HighY        int
	Score        int
	ColumnsAfter Picks
	FieldAfter   Field
}

type Hole struct {
	X	int
	Y	int
}
