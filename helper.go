package main

import (
	"sort"
	//"fmt"
)

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

func _getAllPossiblePositions(piece string, field [][]bool) []Position {
	picks := _getPicks(field)
	picksSum := _sum(picks)
	var positions []Position
	rotationMax := 1
	switch piece {
	case "I", "Z", "S":
		rotationMax = 2
	case "J", "L", "T":
		rotationMax = 4
	}

	for r := 0; r < rotationMax; r++ {
		for i := 0; i < Width; i++ {
			//fmt.Println(piece,r,i)
			fieldAfter := _fieldAfter(field, i, r, piece)
			//columsAfter, maxY := _getColumnsAfter(picks, i, r, piece)
			if !_eq2(field, fieldAfter) {
				//fmt.Println(piece,r,i)
				columsAfter := _getPicks(fieldAfter)
				growMin, growMax := _getGrow(picks, columsAfter)
				damage := _sum(columsAfter) - picksSum
				p := Position{
					Rotation:     r,
					X:            i,
					IsBurn:       _isBurn(fieldAfter),
					Damadge:      damage,
					Score:        (3 * damage) + growMin + growMax,
					ColumnsAfter: columsAfter,
					GrowY:        growMax,
					FieldAfter:   fieldAfter}
				positions = append(positions, p)
			}
		}
	}
	return positions
}

func _calculateMoves(time int) Position {
	var allPositins []Position
	roofIsnear := false
	safePlay := false
	leftRows := Height - MyPlayer.MaxY
	//fmt.Println(Height,MyPlayer.MaxY,leftRows)
	if leftRows <= 5 {
		roofIsnear = true
	}
	if leftRows >= 10 {
		safePlay = true
	}

	//TODO: choose plasements clother to the wall
	//TODO: build 2-wide hole

	// try to burn more
	if MyPlayer.Combo > 0 || !safePlay {

		pos, isFound := _keepUpBurn(allPositins)
		if isFound {
			return pos
		}
	}

	// build minimum damadge with 2-wide hole on the right
	if safePlay {
		// change width and cut field
		Width = Width - 2
		shortField := _trimField(MyPlayer.Field, 2)
		allPositins = _getAllPossiblePositions(CurrentPiece, shortField)
		sort.Sort(ByScore(allPositins))
	}

	// play save try to burn rows and get lowest Y
	if roofIsnear {
		//TODO check if burn and check if no damadge
		sort.Sort(ByMaxY(allPositins))
	}

	return allPositins[0]
}

func _keepUpBurn(positions []Position) (Position, bool) {
	var emptyPos Position
	var burnedPositions []Position

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
			nextPiecePositions := _getAllPossiblePositions(NextPiece, pos.FieldAfter)
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
	var field = make([][]bool, Height)
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
func _isRight(i, right int) bool {
	if i+right < Width {
		return true
	}
	return false
}

func _isUp(i, up int) bool {
	//fmt.Println(i+up,Height)
	if i+up <= Height {
		return true
	}
	return false
}

func _isLeft(i, left int) bool {
	if i-left > 0 {
		return true
	}
	return false
}

func _getColumnsAfter(c []int, i, r int, piece string) ([]int, int) {
	a := make([]int, len(c))
	copy(a, c[:])
	y := 0
	switch piece {
	case "I":
		switch r {
		case 0:
			if _isRight(i, 3) {
				pick := _getPick(i, 3)
				a[i] = pick + 1
				a[i+1] = pick + 1
				a[i+2] = pick + 1
				a[i+3] = pick + 1
				y = pick + 1
			}
		case 1:
			a[i] = c[i] + 4
			y = c[i] + 4
		}
	case "J":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				a[i] = pick + 2
				a[i+1] = pick + 1
				a[i+2] = pick + 1
				y = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				if c[i+1] >= c[i]+2 {
					a[i] = c[i+1] + 1
					a[i+1] = c[i+1] + 1
					y = c[i+1] + 1
				} else {
					a[i] = c[i] + 3
					a[i+1] = c[i] + 3
					y = c[i] + 3
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i+2] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				a[i] = pick + 1
				a[i+1] = pick + 3
				y = pick + 3
			}
		}
	case "L":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				a[i] = pick + 1
				a[i+1] = pick + 1
				a[i+2] = pick + 2
				y = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				a[i] = pick + 3
				a[i+1] = pick + 1
				y = pick + 3
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				if c[i] >= c[i+1]+2 {
					a[i] = c[i] + 1
					a[i+1] = c[i] + 1
					y = c[i] + 1
				} else {
					a[i] = c[i+1] + 3
					a[i+1] = c[i+1] + 3
					y = c[i+1] + 3
				}
			}
		}
	case "O":
		if _isRight(i, 1) {
			pick := _getPick(i, 1)
			a[i] = pick + 2
			a[i+1] = pick + 2
			y = pick + 2
		}
	case "S":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i] || pick == c[i+1] {
					a[i] = pick + 1
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i+1] {
					a[i] = pick + 3
					a[i+1] = pick + 2
					y = pick + 3
				} else {
					a[i] = pick + 2
					a[i+1] = pick + 1
					y = pick + 2
				}
			}
		}
	case "T":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				a[i] = pick + 1
				a[i+1] = pick + 2
				a[i+2] = pick + 1
				y = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i] {
					a[i] = pick + 3
					a[i+1] = pick + 2
					y = pick + 3
				} else {
					a[i] = pick + 2
					a[i+1] = pick + 1
					y = pick + 2
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i+1] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i+1] {
					a[i] = pick + 2
					a[i+1] = pick + 3
					y = pick + 3

				} else {
					a[i] = pick + 1
					a[i+1] = pick + 2
					y = pick + 2
				}
			}
		}
	case "Z":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i+1] || pick == c[i+2] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 1
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick
					y = pick + 1
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i] {
					a[i] = pick + 2
					a[i+1] = pick + 3
					y = pick + 3
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 2
					y = pick + 2
				}
			}
		}
	}
	return a, y
}

func _getPicks(f [][]bool) []int {
	result := make([]int, len(f[0]))
	for i, row := range f {
		for j, col := range row {
			if i+1 > result[j] && col == true {
				result[j] = i + 1
			}
		}
	}
	return result
}

func _fieldAfter(f [][]bool, i, r int, piece string) [][]bool {
	a := make([][]bool, len(f))
	for i, row := range f {
		a[i] = make([]bool, len(row))
		copy(a[i], row[:])
	}

	switch piece {
	case "I":
		switch r {
		case 0:
			if _isRight(i, 3) {
				pick := _getPick(i, 3)
				if _isUp(pick, 1) {
					a[pick][i] = true
					a[pick][i+1] = true
					a[pick][i+2] = true
					a[pick][i+3] = true
				}
			}
		case 1:
			pick := MyPlayer.Columns[i]
			if _isUp(pick, 4) {
				a[pick][i] = true
				a[pick+1][i] = true
				a[pick+2][i] = true
				a[pick+3][i] = true
			}
		}
	case "J":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if _isUp(pick, 2) {
					a[pick][i] = true
					a[pick+1][i] = true
					a[pick][i+1] = true
					a[pick][i+2] = true
				}
			}
		case 1:
			if _isRight(i, 1) {
				l := MyPlayer.Columns[i]
				r := MyPlayer.Columns[i+1]
				if r >= l+2 {
					if _isUp(r, 1) {
						a[r][i] = true
						a[r][i+1] = true
						a[r-1][i] = true
						a[r-2][i] = true
					}
				} else {
					if _isUp(l, 3) {
						a[l][i] = true
						a[l+1][i] = true
						a[l+2][i] = true
						a[l+2][i+1] = true
					}
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				r := MyPlayer.Columns[i+2]
				if pick == r {
					if _isUp(r, 2) {
						a[pick+1][i] = true
						a[pick+1][i+1] = true
						a[pick+1][i+2] = true
						a[pick][i+2] = true
					}
				} else {
					if _isUp(r, 1) {
						a[pick][i] = true
						a[pick][i+1] = true
						a[pick][i+2] = true
						a[pick-1][i+2] = true
					}
				}
			}
		case 3:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if _isUp(pick, 3) {
					a[pick][i] = true
					a[pick][i+1] = true
					a[pick+1][i+1] = true
					a[pick+2][i+1] = true
				}
			}
		}
	case "L":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if _isUp(pick, 2) {
					a[pick][i] = true
					a[pick][i+1] = true
					a[pick][i+2] = true
					a[pick+1][i+2] = true
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if _isUp(pick, 3) {
					a[pick][i] = true
					a[pick+1][i] = true
					a[pick+2][i] = true
					a[pick][i+1] = true
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				l := MyPlayer.Columns[i]
				if pick == l {
					if _isUp(l, 2) {
						a[pick][i] = true
						a[pick+1][i] = true
						a[pick+1][i+1] = true
						a[pick+1][i+2] = true
					}
				} else {
					if _isUp(l, 1) {
						a[pick-1][i] = true
						a[pick][i] = true
						a[pick][i+1] = true
						a[pick][i+2] = true
					}
				}
			}
		case 3:
			if _isRight(i, 1) {
				l := MyPlayer.Columns[i]
				r := MyPlayer.Columns[i+1]
				if l >= r+2 {
					if _isUp(l, 1) {
						a[l][i] = true
						a[l][i+1] = true
						a[l-1][i+1] = true
						a[l-2][i+1] = true
					}
				} else {
					if _isUp(r, 3) {
						a[r+2][i] = true
						a[r][i+1] = true
						a[r+1][i+1] = true
						a[r+2][i+1] = true
					}
				}
			}
		}
	case "O":
		if _isRight(i, 1) {
			pick := _getPick(i, 1)
			if _isUp(pick, 2) {
				a[pick][i] = true
				a[pick+1][i] = true
				a[pick][i+1] = true
				a[pick+1][i+1] = true
			}
		}
	case "S":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				l := MyPlayer.Columns[i]
				l1 := MyPlayer.Columns[i+1]
				if pick == l || pick == l1 {
					if _isUp(pick, 2) {
						a[pick][i] = true
						a[pick][i+1] = true
						a[pick+1][i+1] = true
						a[pick+1][i+2] = true
					}
				} else {
					if _isUp(pick, 1) {
						a[pick-1][i] = true
						a[pick-1][i+1] = true
						a[pick][i+1] = true
						a[pick][i+2] = true
					}
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				r := MyPlayer.Columns[i+1]
				if pick == r {
					if _isUp(pick, 3) {
						a[pick+2][i] = true
						a[pick+1][i] = true
						a[pick+1][i+1] = true
						a[pick][i+1] = true
					}
				} else {
					if _isUp(pick, 2) {
						a[pick+1][i] = true
						a[pick][i] = true
						a[pick][i+1] = true
						a[pick-1][i+1] = true
					}
				}
			}
		}
	case "T":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if _isUp(pick, 2) {
					a[pick][i] = true
					a[pick][i+1] = true
					a[pick+1][i+1] = true
					a[pick][i+2] = true
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				l := MyPlayer.Columns[i]
				if pick == l {
					if _isUp(pick, 3) {
						a[pick][i] = true
						a[pick+1][i] = true
						a[pick+1][i+1] = true
						a[pick+2][i] = true
					}
				} else {
					if _isUp(pick, 2) {
						a[pick-1][i] = true
						a[pick][i] = true
						a[pick][i+1] = true
						a[pick+1][i] = true
					}
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				c := MyPlayer.Columns[i+1]
				if pick == c {
					if _isUp(pick, 2) {
						a[pick+1][i] = true
						a[pick][i+1] = true
						a[pick+1][i+1] = true
						a[pick+1][i+2] = true
					}
				} else {
					if _isUp(pick, 1) {
						a[pick][i] = true
						a[pick][i+1] = true
						a[pick-1][i+1] = true
						a[pick][i+2] = true
					}
				}
			}
		case 3:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				r := MyPlayer.Columns[i+1]
				if pick == r {
					if _isUp(pick, 3) {
						a[pick+2][i+1] = true
						a[pick+1][i] = true
						a[pick+1][i+1] = true
						a[pick][i+1] = true
					}
				} else {
					if _isUp(pick, 2) {
						a[pick+1][i+1] = true
						a[pick][i] = true
						a[pick][i+1] = true
						a[pick-1][i+1] = true
					}
				}
			}
		}
	case "Z":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				l1 := MyPlayer.Columns[i+1]
				l2 := MyPlayer.Columns[i+2]
				if pick == l1 || pick == l2 {
					if _isUp(pick, 2) {
						a[pick+1][i] = true
						a[pick+1][i+1] = true
						a[pick][i+1] = true
						a[pick][i+2] = true
					}
				} else {
					if _isUp(pick, 1) {
						a[pick][i] = true
						a[pick][i+1] = true
						a[pick-1][i+1] = true
						a[pick-1][i+2] = true
					}
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				l := MyPlayer.Columns[i]
				if pick == l {
					if _isUp(pick, 3) {
						a[pick][i] = true
						a[pick+1][i] = true
						a[pick+1][i+1] = true
						a[pick+2][i+1] = true
					}
				} else {
					if _isUp(pick, 2) {
						a[pick-1][i] = true
						a[pick][i] = true
						a[pick][i+1] = true
						a[pick+1][i+1] = true
					}
				}
			}
		}
	}
	return a
}

func _getPick(i, v int) int {
	pick := MyPlayer.Columns[i]
	for n := 1; n <= v; n++ {
		if pick < MyPlayer.Columns[i+n] {
			pick = MyPlayer.Columns[i+n]
		}
	}
	return pick
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

func _getMaxY(c []int) int {
	maxY := 0
	for _, col := range c {
		if col > maxY {
			maxY = col
		}
	}
	return maxY
}

func _sum(c []int) int {
	sum := 0
	for i := 0; i < len(c); i++ {
		sum += c[i]
	}
	return sum
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

func _eq2(a, b [][]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func _eq1(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
