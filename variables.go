package main

var Timebank, TimePerMove, Width, OriginalHeight, Height, Pick, Round, CurrentPieceX, CurrentPieceY int
var Players [2]Player
var MyPlayer *Player
var CurrentPiece, NextPiece string

type Player struct {
	Name    string
	Columns []int
	MaxY    int
	Field   [][]bool
	Points  int
	Combo   int
}

type Position struct {
	Rotation     int
	X            int
	IsBurn       int
	Damadge      int
	GrowY        int
	Score        int
	ColumnsAfter []int
	FieldAfter   [][]bool
}

type Field [][] bool
func (f Field) Width() int {	return len(f[0])}
func (f Field) Height() int {	return len(f)}
func (f Field) IsFit(pick, up int) bool {
	//fmt.Println(i+up,Height)
	if pick+up <= f.Height() {
		return true
	}
	return false
}
func (f Field) Columns Picks {
	result := make([]int, f.Width())
	for i, row := range f {
		for j, col := range row {
			if i+1 > result[j] && col == true {
				result[j] = i + 1
			}
		}
	}
	return result
}

type Picks []int
func (p Picks) Max() int {
	result := 0
	for _,pick := range p {
		if result < pick {
		result = pick
		}
	}
	return result
}
func (p Picks) Max(x,n int) int {
	pick := p[x]
	for i := 1; i <= n; i++ {
		if pick < p[x+i] {
			pick = p[x+i]
		}
	}
	return pick
}
func (p Picks) IsRight(x,n int) bool {
	if x+n < len(p) {
		return true
	}
	return false
}
