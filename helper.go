package main

import (
	"bytes"
	"fmt"
	"sort"
)

//TODO rewrite to extentions for colums

type ByDamage []Position
type ByMaxY []Position
type ByScore []Position

func (a ByDamage) Len() int           { return len(a) }
func (a ByDamage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDamage) Less(i, j int) bool { return a[i].Damadge < a[j].Damadge }

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

func (a ByMaxY) Len() int           { return len(a) }
func (a ByMaxY) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMaxY) Less(i, j int) bool { return a[i].GrowY < a[j].GrowY }

func _availablePositions(piece string, field Field) []Position {
	w := field.Width()
	picks := field.Picks()
	var positions []Position
	rotationMax := 1

	switch piece {
	case "I", "Z", "S":
		rotationMax = 2
	case "J", "L", "T":
		rotationMax = 4
	}

	for r := 0; r < rotationMax; r++ {
		for i := 0; i < w; i++ {
			//fmt.Println(piece,r,i)
			fieldAfter := _fieldAfter(field, i, r, piece)
			//columsAfter, maxY := _getColumnsAfter(picks, i, r, piece)
			if !field.Equal(fieldAfter) {
				//fmt.Println(piece,r,i)
				picksAfter := fieldAfter.Picks()
				growMin, growMax := _getGrow(picks, picksAfter)
				//damage := _sum(columsAfter) - picksSum //kill this and get data from grow
				p := Position{
					Rotation: r,
					X:        i,
					IsBurn:   _isBurn(fieldAfter),
					//Damadge:      damage,
					//Score:        (3 * damage) + growMin + growMax,
					//ColumnsAfter: picksAfter,
					GrowY:      growMax + growMin, //very wrong redu this
					FieldAfter: fieldAfter}
				positions = append(positions, p)
			}
		}
	}
	return positions
}

func _calculateMoves() Position {
	//TODO: choose plasements clother to the wall

	// try to burn more
	if MyPlayer.Combo > 0 {
		pos, isFound := _keepUpBurn()
		if isFound {
			return pos
		}
	}

	if MyPlayer.State == "safe" {
		shortField := _trimField(MyPlayer.Field, 2)
		shortPositions := _availablePositions(CurrentPiece, shortField)
		sort.Sort(ByDamage(shortPositions))
		return shortPositions[0] //TODO: check lowest fit and predict next piece
	}

	positions := _availablePositions(CurrentPiece, MyPlayer.Field)

	if MyPlayer.State == "normal" {
		sort.Sort(ByDamage(positions))
		//TODO: check lowest fit and predict next piece
	}

	// play save try to burn rows and get lowest Y
	if MyPlayer.State == "dangerous" {
		//TODO check if burn and check if no damadge
		sort.Sort(ByMaxY(positions))
		//TODO: check lowest fit and predict next piece
	}

	return positions[0]
}

func _keepUpBurn() (Position, bool) {
	var emptyPos Position
	var burnedPositions []Position
	positions := _availablePositions(CurrentPiece, MyPlayer.Field)

	for _, pos := range positions {
		if pos.IsBurn > 0 {
			burnedPositions = append(burnedPositions, pos)
		}
	}
	burnedPositionsTotal := len(burnedPositions)

	if burnedPositionsTotal == 1 {
		return burnedPositions[0], true
	}

	//see if next peacie will burn rows
	if burnedPositionsTotal > 1 {
		//sort first
		sort.Sort(ByDamage(burnedPositions))

		bIndex := 0
		for current_i, pos := range burnedPositions {
			nextPiecePositions := _availablePositions(NextPiece, pos.FieldAfter)
			for _, nextPos := range nextPiecePositions {
				if nextPos.IsBurn > 0 {
					bIndex = current_i
					break
				}
			}
		}
		return burnedPositions[bIndex], true
	}
	return emptyPos, false
}

func _trimField(f [][]bool, trim int) [][]bool {
	var field = make([][]bool, len(f))
	newSize := len(f[0]) - trim
	for rowIndex, row := range f {
		colums := make([]bool, newSize)
		copy(colums, row[:])
		field[rowIndex] = colums
	}
	return field
}

/*
func _isHole(cols []int, piece string) bool {
	for i, c := range cols {
		if _isRight(i, 1) && (c-cols[i+1] < -2 || c-cols[i+1] > 2) && piece != "I" && NextPiece != "I" {
			return true
		}
	}
	return false
}
*/

func _fieldAfter(f Field, x, r int, piece string) Field {
	picks := f.Picks()
	w := f.Width()
	a := make([][]bool, f.Height())
	for i, row := range f {
		a[i] = make([]bool, w)
		copy(a[i], row[:])
	}

	switch piece {
	case "I":
		switch r {
		case 0:
			if picks.IsRight(x, 3) {
				pick := picks.MaxR(x, 3)
				if f.IsFit(pick, 1) {
					a[pick][x] = true
					a[pick][x+1] = true
					a[pick][x+2] = true
					a[pick][x+3] = true
				}
			}
		case 1:
			pick := picks[x]
			if f.IsFit(pick, 4) {
				a[pick][x] = true
				a[pick+1][x] = true
				a[pick+2][x] = true
				a[pick+3][x] = true
			}
		}
	case "J":
		switch r {
		case 0:
			if picks.IsRight(x, 2) {
				pick := picks.MaxR(x, 2)
				if f.IsFit(pick, 2) {
					a[pick][x] = true
					a[pick+1][x] = true
					a[pick][x+1] = true
					a[pick][x+2] = true
				}
			}
		case 1:
			if picks.IsRight(x, 1) {
				l := picks[x]
				l2 := picks[x+1]
				if l2 >= l+2 {
					if f.IsFit(l2, 1) {
						a[l2][x] = true
						a[l2][x+1] = true
						a[l2-1][x] = true
						a[l2-2][x] = true
					}
				} else {
					if f.IsFit(l, 3) {
						a[l][x] = true
						a[l+1][x] = true
						a[l+2][x] = true
						a[l+2][x+1] = true
					}
				}
			}
		case 2:
			if picks.IsRight(x, 2) {
				pick := picks.MaxR(x, 2)
				l3 := picks[x+2]
				if pick == l3 {
					if f.IsFit(l3, 2) {
						a[pick+1][x] = true
						a[pick+1][x+1] = true
						a[pick+1][x+2] = true
						a[pick][x+2] = true
					}
				} else {
					if f.IsFit(l3, 1) {
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick][x+2] = true
						a[pick-1][x+2] = true
					}
				}
			}
		case 3:
			if picks.IsRight(x, 1) {
				pick := picks.MaxR(x, 1)
				if f.IsFit(pick, 3) {
					a[pick][x] = true
					a[pick][x+1] = true
					a[pick+1][x+1] = true
					a[pick+2][x+1] = true
				}
			}
		}
	case "L":
		switch r {
		case 0:
			if picks.IsRight(x, 2) {
				pick := picks.MaxR(x, 2)
				if f.IsFit(pick, 2) {
					a[pick][x] = true
					a[pick][x+1] = true
					a[pick][x+2] = true
					a[pick+1][x+2] = true
				}
			}
		case 1:
			if picks.IsRight(x, 1) {
				pick := picks.MaxR(x, 1)
				if f.IsFit(pick, 3) {
					a[pick][x] = true
					a[pick+1][x] = true
					a[pick+2][x] = true
					a[pick][x+1] = true
				}
			}
		case 2:
			if picks.IsRight(x, 2) {
				pick := picks.MaxR(x, 2)
				l := picks[x]
				if pick == l {
					if f.IsFit(l, 2) {
						a[pick][x] = true
						a[pick+1][x] = true
						a[pick+1][x+1] = true
						a[pick+1][x+2] = true
					}
				} else {
					if f.IsFit(l, 1) {
						a[pick-1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick][x+2] = true
					}
				}
			}
		case 3:
			if picks.IsRight(x, 1) {
				l := picks[x]
				l2 := picks[x+1]
				if l >= l2+2 {
					if f.IsFit(l, 1) {
						a[l][x] = true
						a[l][x+1] = true
						a[l-1][x+1] = true
						a[l-2][x+1] = true
					}
				} else {
					if f.IsFit(l2, 3) {
						a[l2+2][x] = true
						a[l2][x+1] = true
						a[l2+1][x+1] = true
						a[l2+2][x+1] = true
					}
				}
			}
		}
	case "O":
		if picks.IsRight(x, 1) {
			pick := picks.MaxR(x, 1)
			if f.IsFit(pick, 2) {
				a[pick][x] = true
				a[pick+1][x] = true
				a[pick][x+1] = true
				a[pick+1][x+1] = true
			}
		}
	case "S":
		switch r {
		case 0:
			if picks.IsRight(x, 2) {
				pick := picks.MaxR(x, 2)
				l := picks[x]
				l1 := picks[x+1]
				if pick == l || pick == l1 {
					if f.IsFit(pick, 2) {
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick+1][x+1] = true
						a[pick+1][x+2] = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick-1][x] = true
						a[pick-1][x+1] = true
						a[pick][x+1] = true
						a[pick][x+2] = true
					}
				}
			}
		case 1:
			if picks.IsRight(x, 1) {
				pick := picks.MaxR(x, 1)
				l2 := picks[x+1]
				if pick == l2 {
					if f.IsFit(pick, 3) {
						a[pick+2][x] = true
						a[pick+1][x] = true
						a[pick+1][x+1] = true
						a[pick][x+1] = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a[pick+1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick-1][x+1] = true
					}
				}
			}
		}
	case "T":
		switch r {
		case 0:
			if picks.IsRight(x, 2) {
				pick := picks.MaxR(x, 2)
				if f.IsFit(pick, 2) {
					a[pick][x] = true
					a[pick][x+1] = true
					a[pick+1][x+1] = true
					a[pick][x+2] = true
				}
			}
		case 1:
			if picks.IsRight(x, 1) {
				pick := picks.MaxR(x, 1)
				l := picks[x]
				if pick == l {
					if f.IsFit(pick, 3) {
						a[pick][x] = true
						a[pick+1][x] = true
						a[pick+1][x+1] = true
						a[pick+2][x] = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a[pick-1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick+1][x] = true
					}
				}
			}
		case 2:
			if picks.IsRight(x, 2) {
				pick := picks.MaxR(x, 2)
				c := picks[x+1]
				if pick == c {
					if f.IsFit(pick, 2) {
						a[pick+1][x] = true
						a[pick][x+1] = true
						a[pick+1][x+1] = true
						a[pick+1][x+2] = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick-1][x+1] = true
						a[pick][x+2] = true
					}
				}
			}
		case 3:
			if picks.IsRight(x, 1) {
				pick := picks.MaxR(x, 1)
				l2 := picks[x+1]
				if pick == l2 {
					if f.IsFit(pick, 3) {
						a[pick+2][x+1] = true
						a[pick+1][x] = true
						a[pick+1][x+1] = true
						a[pick][x+1] = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a[pick+1][x+1] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick-1][x+1] = true
					}
				}
			}
		}
	case "Z":
		switch r {
		case 0:
			if picks.IsRight(x, 2) {
				pick := picks.MaxR(x, 2)
				l1 := picks[x+1]
				l2 := picks[x+2]
				if pick == l1 || pick == l2 {
					if f.IsFit(pick, 2) {
						a[pick+1][x] = true
						a[pick+1][x+1] = true
						a[pick][x+1] = true
						a[pick][x+2] = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick-1][x+1] = true
						a[pick-1][x+2] = true
					}
				}
			}
		case 1:
			if picks.IsRight(x, 1) {
				pick := picks.MaxR(x, 1)
				l := picks[x]
				if pick == l {
					if f.IsFit(pick, 3) {
						a[pick][x] = true
						a[pick+1][x] = true
						a[pick+1][x+1] = true
						a[pick+2][x+1] = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a[pick-1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick+1][x+1] = true
					}
				}
			}
		}
	}
	return a
}

func _getGrow(b, a []int) (int, int) {
	maxY := 0
	minY := 1000
	for i, col := range b {
		if (a[i] - col) > 0 {

			if col > maxY {
				maxY = col
			}

			if a[i] < minY {
				minY = a[i]
			}
		}
	}
	return minY, maxY
}

func _isBurn(f [][]bool) int {
	burn := 0
	for _, row := range f {
		check := true
		for _, col := range row {
			if !col {
				check = false
				break
			}
		}
		if check {
			burn++
		}
	}
	return burn
}

func _printMoves(pos Position) {
	var buffer bytes.Buffer
	for i := 0; i < pos.Rotation; i++ {
		buffer.WriteString("turnright,")
	}
	if pos.Rotation == 1 {
		CurrentPieceX = CurrentPieceX + 1
		if CurrentPiece == "I" {
			CurrentPieceX = CurrentPieceX + 1
		}
	}
	if CurrentPieceX > pos.X {
		for i := 0; i < CurrentPieceX-pos.X; i++ {
			buffer.WriteString("left,")
		}
	}
	if CurrentPieceX < pos.X {
		for i := 0; i < pos.X-CurrentPieceX; i++ {
			buffer.WriteString("right,")
		}
	}
	buffer.WriteString("drop")
	fmt.Println(buffer.String())
}

func _roundOne() {
	fmt.Println("drop")
}
