package main

import (
	//"fmt"
	//"sort"
	"strings"
)

type Field [][]bool

func (f *Field) init(raw string) Field {
	rows := strings.Split(raw, ";")
	height := len(rows)
	var field Field
	field = make([][]bool, height)
	for rowIndex, row := range rows {
		y := height - rowIndex
		cells := strings.Split(row, ",")
		var colums = make([]bool, len(cells))
		for columIndex, value := range cells {
			if value == "2" {
				colums[columIndex] = true
			} else {
				colums[columIndex] = false
			}
		}
		field[y-1] = colums
	}
	return field
}

func (f Field) Width() int { return len(f[0]) }

func (f Field) Height() int { return len(f) }

func (f Field) IsFit(pick, up int) bool {
	if pick+up <= f.Height() {
		return true
	}
	return false
}

func (f Field) Picks() Picks {
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

func (f Field) Equal(b Field) bool {
	if f.Height() != b.Height() {
		return false
	}
	for i := range f {
		if len(f[i]) != len(b[i]) {
			return false
		}
		for j := range f[i] {
			if f[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func (f Field) Trim(trim int) Field {
	var trimed = make([][]bool, len(f))
	newSize := len(f[0]) - trim
	for rowIndex, row := range f {
		colums := make([]bool, newSize)
		copy(colums, row[:])
		trimed[rowIndex] = colums
	}
	return trimed
}

func (f Field) Positions(piece string, st Strategy) []Position {
	w := f.Width()
	picks := f.Picks()
	hBlocked, hLeft, hRight := f.FindHoles(picks)
	positions := make([]Position, 0)
	rotationMax := 1

	switch piece {
	case "I", "Z", "S":
		rotationMax = 2
	case "J", "L", "T":
		rotationMax = 4
	}

	for r := 0; r < rotationMax; r++ {
		positions = append(positions, f.topPositions(st, r, w, picks, piece, hBlocked)...)
		if len(hLeft) > 0 {
			positions = append(positions, f.leftPositions(st, r, piece, hLeft)...)
		}
		if len(hRight) > 0 {
			positions = append(positions, f.rightPositions(st, r, piece, hRight)...)
		}
	}
	return positions
}

func (f Field) topPositions(st Strategy, r, w int, picks Picks, piece string, holes []Cell) []Position {
	pos := make([]Position, 0)
	for i := 0; i < w; i++ {
		fieldAfter := f.After(i, r, piece)
		if fieldAfter != nil {
			p := Position{Rotation: r, X: i}
			p.InitTop(picks, fieldAfter, holes, st)
			pos = append(pos, p)
		}
	}
	return pos
}

func (f Field) leftPositions(st Strategy, r int, piece string, holes []Cell) []Position {
	pos := make([]Position, 0)
	for _, h := range holes {
		fieldAfterLeft := f.AfterLeftFix(r, piece, h)
		if fieldAfterLeft != nil {
			p := Position{Rotation: r, X: int(h.X)}
			p.InitLeft(fieldAfterLeft, st)
			pos = append(pos, p)
		}
	}
	return pos
}

func (f Field) rightPositions(st Strategy, r int, piece string, holes []Cell) []Position {
	pos := make([]Position, 0)
	for _, h := range holes {
		fieldAfterRight := f.AfterRightFix(r, piece, h)
		if fieldAfterRight != nil {
			p := Position{Rotation: r, X: int(h.X)}
			p.InitRight(fieldAfterRight, st)
			pos = append(pos, p)
		}
	}
	return pos
}

func (f Field) WillBurn() int {
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

func (f Field) Burn() {
	for i, row := range f {
		check := true
		for _, col := range row {
			if !col {
				check = false
			}
		}
		if check && i < len(f) { //delete line
			//fmt.Println(len(f), i)
			f = append(f[:i], f[i+1:]...)
		}
	}
}

func (f Field) FindHoles(picks Picks) ([]Cell, []Cell, []Cell) {
	blocked := make([]Cell, 0)
	left := make([]Cell, 0)
	right := make([]Cell, 0)
	for i, pick := range picks {
		for j := 0; j < pick; j++ {
			if !f[j][i] {
				hole := Cell{X: i, Y: j}
				if i-2 > -1 && !f[j][i-1] && !f[j][i-2] {
					left = append(left, hole)
				} else if i+2 < f.Width() && !f[j][i+1] && !f[j][i+2] {
					right = append(right, hole)
				} else {
					blocked = append(blocked, hole)
				}
			}
		}
	}
	return blocked, left, right
}

func (f Field) After(x, r int, piece string) Field {
	valid := false
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
					valid = true
				}
			}
		case 1:
			pick := picks[x]
			if f.IsFit(pick, 4) {
				a[pick][x] = true
				a[pick+1][x] = true
				a[pick+2][x] = true
				a[pick+3][x] = true
				valid = true
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
					valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(l, 3) {
						a[l][x] = true
						a[l+1][x] = true
						a[l+2][x] = true
						a[l+2][x+1] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick][x+2] = true
						a[pick-1][x+2] = true
						valid = true
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
					valid = true
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
					valid = true
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
					valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick-1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick][x+2] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(l2, 3) {
						a[l2+2][x] = true
						a[l2][x+1] = true
						a[l2+1][x+1] = true
						a[l2+2][x+1] = true
						valid = true
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
				valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick-1][x] = true
						a[pick-1][x+1] = true
						a[pick][x+1] = true
						a[pick][x+2] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a[pick+1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick-1][x+1] = true
						valid = true
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
					valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a[pick-1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick+1][x] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick-1][x+1] = true
						a[pick][x+2] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a[pick+1][x+1] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick-1][x+1] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick-1][x+1] = true
						a[pick-1][x+2] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a[pick-1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick+1][x+1] = true
						valid = true
					}
				}
			}
		}
	}
	if valid {
		return a
	}
	return nil
}

func (f Field) AfterLeftFix(r int, piece string, h Cell) Field {
	valid := false
	fw := f.Width()
	fh := f.Height()
	a := make([][]bool, fh)
	for i, row := range f {
		a[i] = make([]bool, fw)
		copy(a[i], row[:])
	}

	switch piece {
	/*case "I":
	switch r {
	case 0:
		if picks.IsRight(x, 3) {
			pick := picks.MaxR(x, 3)
			if f.IsFit(pick, 1) {
				a[pick][x] = true
				a[pick][x+1] = true
				a[pick][x+2] = true
				a[pick][x+3] = true
				valid = true
			}
		}
	case 1:
		pick := picks[x]
		if f.IsFit(pick, 4) {
			a[pick][x] = true
			a[pick+1][x] = true
			a[pick+2][x] = true
			a[pick+3][x] = true
			valid = true
		}
	}
	*/
	/*case "J":
	switch r {
	case 0:
		if picks.IsRight(x, 2) {
			pick := picks.MaxR(x, 2)
			if f.IsFit(pick, 2) {
				a[pick][x] = true
				a[pick+1][x] = true
				a[pick][x+1] = true
				a[pick][x+2] = true
				valid = true
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
					valid = true
				}
			} else {
				if f.IsFit(l, 3) {
					a[l][x] = true
					a[l+1][x] = true
					a[l+2][x] = true
					a[l+2][x+1] = true
					valid = true
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
					valid = true
				}
			} else {
				if f.IsFit(pick, 1) {
					a[pick][x] = true
					a[pick][x+1] = true
					a[pick][x+2] = true
					a[pick-1][x+2] = true
					valid = true
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
				valid = true
			}
		}
	}*/
	/*case "L":
		switch r {
		case 0:
			if picks.IsRight(x, 2) {
				pick := picks.MaxR(x, 2)
				if f.IsFit(pick, 2) {
					a[pick][x] = true
					a[pick][x+1] = true
					a[pick][x+2] = true
					a[pick+1][x+2] = true
					valid = true
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
					valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick-1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick][x+2] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(l2, 3) {
						a[l2+2][x] = true
						a[l2][x+1] = true
						a[l2+1][x+1] = true
						a[l2+2][x+1] = true
						valid = true
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
				valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick-1][x] = true
						a[pick-1][x+1] = true
						a[pick][x+1] = true
						a[pick][x+2] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a[pick+1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick-1][x+1] = true
						valid = true
					}
				}
			}
		}*/
	case "T":
		switch r {
		case 0:
			if !f[h.Y][h.X-1] && !f[h.Y+1][h.X-1] && !f[h.Y][h.X-2] {
				a[h.Y][h.X] = true
				a[h.Y][h.X-1] = true
				a[h.Y+1][h.X-1] = true
				a[h.Y][h.X-2] = true
				valid = true
			}
		case 1:
			if h.Y-1 > -1 && !f[h.Y][h.X-1] && !f[h.Y+1][h.X-1] && !f[h.Y-1][h.X-1] {
				a[h.Y][h.X] = true
				a[h.Y][h.X-1] = true
				a[h.Y+1][h.X-1] = true
				a[h.Y-1][h.X-1] = true
				valid = true
			}
		case 2:
			if h.Y-1 > -1 && !f[h.Y][h.X-1] && !f[h.Y][h.X-2] && !f[h.Y-1][h.X-1] {
				a[h.Y][h.X] = true
				a[h.Y][h.X-1] = true
				a[h.Y][h.X-2] = true
				a[h.Y-1][h.X-1] = true
				valid = true
			}
		case 3:
			if h.Y+2 < fh && !f[h.Y+1][h.X] && !f[h.Y+1][h.X-1] && !f[h.Y+2][h.X] {
				a[h.Y][h.X] = true
				a[h.Y+1][h.X] = true
				a[h.Y+1][h.X-1] = true
				a[h.Y+2][h.X] = true
				valid = true
			}
		}
		/*case "Z":
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick-1][x+1] = true
						a[pick-1][x+2] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a[pick-1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick+1][x+1] = true
						valid = true
					}
				}
			}
		}
		*/
	}
	if valid {
		return a
	}
	return nil
}

func (f Field) AfterRightFix(r int, piece string, h Cell) Field {
	valid := false
	fw := f.Width()
	fh := f.Height()
	a := make([][]bool, fh)
	for i, row := range f {
		a[i] = make([]bool, fw)
		copy(a[i], row[:])
	}

	switch piece {
	/*case "I":
	switch r {
	case 0:
		if picks.IsRight(x, 3) {
			pick := picks.MaxR(x, 3)
			if f.IsFit(pick, 1) {
				a[pick][x] = true
				a[pick][x+1] = true
				a[pick][x+2] = true
				a[pick][x+3] = true
				valid = true
			}
		}
	case 1:
		pick := picks[x]
		if f.IsFit(pick, 4) {
			a[pick][x] = true
			a[pick+1][x] = true
			a[pick+2][x] = true
			a[pick+3][x] = true
			valid = true
		}
	}
	*/
	/*case "J":
	switch r {
	case 0:
		if picks.IsRight(x, 2) {
			pick := picks.MaxR(x, 2)
			if f.IsFit(pick, 2) {
				a[pick][x] = true
				a[pick+1][x] = true
				a[pick][x+1] = true
				a[pick][x+2] = true
				valid = true
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
					valid = true
				}
			} else {
				if f.IsFit(l, 3) {
					a[l][x] = true
					a[l+1][x] = true
					a[l+2][x] = true
					a[l+2][x+1] = true
					valid = true
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
					valid = true
				}
			} else {
				if f.IsFit(pick, 1) {
					a[pick][x] = true
					a[pick][x+1] = true
					a[pick][x+2] = true
					a[pick-1][x+2] = true
					valid = true
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
				valid = true
			}
		}
	}*/
	/*case "L":
		switch r {
		case 0:
			if picks.IsRight(x, 2) {
				pick := picks.MaxR(x, 2)
				if f.IsFit(pick, 2) {
					a[pick][x] = true
					a[pick][x+1] = true
					a[pick][x+2] = true
					a[pick+1][x+2] = true
					valid = true
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
					valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick-1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick][x+2] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(l2, 3) {
						a[l2+2][x] = true
						a[l2][x+1] = true
						a[l2+1][x+1] = true
						a[l2+2][x+1] = true
						valid = true
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
				valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick-1][x] = true
						a[pick-1][x+1] = true
						a[pick][x+1] = true
						a[pick][x+2] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a[pick+1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick-1][x+1] = true
						valid = true
					}
				}
			}
		}*/
	case "T":
		switch r {
		case 0:
			if !f[h.Y][h.X+1] && !f[h.Y+1][h.X+1] && !f[h.Y][h.X+2] {
				a[h.Y][h.X] = true
				a[h.Y][h.X+1] = true
				a[h.Y+1][h.X+1] = true
				a[h.Y][h.X+2] = true
				valid = true
			}
		case 1:
			if h.Y+2 < fh && !f[h.Y+1][h.X] && !f[h.Y+1][h.X+1] && !f[h.Y+2][h.X] {
				a[h.Y][h.X] = true
				a[h.Y+1][h.X] = true
				a[h.Y+1][h.X+1] = true
				a[h.Y+2][h.X] = true
				valid = true
			}
		case 2:
			if h.Y-1 > -1 && !f[h.Y][h.X+1] && !f[h.Y][h.X+2] && !f[h.Y-1][h.X+1] {
				a[h.Y][h.X] = true
				a[h.Y][h.X+1] = true
				a[h.Y][h.X+2] = true
				a[h.Y-1][h.X+1] = true
				valid = true
			}
		case 3:
			if h.Y-1 > -1 && !f[h.Y][h.X+1] && !f[h.Y+1][h.X+1] && !f[h.Y-1][h.X+1] {
				a[h.Y][h.X] = true
				a[h.Y][h.X+1] = true
				a[h.Y+1][h.X+1] = true
				a[h.Y-1][h.X+1] = true
				valid = true
			}
		}
		/*case "Z":
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick-1][x+1] = true
						a[pick-1][x+2] = true
						valid = true
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
						valid = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a[pick-1][x] = true
						a[pick][x] = true
						a[pick][x+1] = true
						a[pick+1][x+1] = true
						valid = true
					}
				}
			}
		}
		*/
	}
	if valid {
		return a
	}
	return nil
}

func (f Field) IsValid(cells *map[string]Cell) bool {
	for _, c := range *cells {
		if c.X < 0 || c.X >= f.Width() || c.Y < 0 {
			return false
		}
		if c.Y < f.Height() && f[c.Y][c.X] {
			return false
		}
	}
	return true
}

/*
func (f Field) CalculatePath(pos Position, piece Piece) string {
	path := []string{}
	// Top positions only

	maxTry := 10
	np := Piece{}
	switch piece.Name {
	case "I":
		if pos.Rotation == 1 && pos.X <= piece.CurrentX {
			pos.Rotation = 3
		}

		for (pos.X != piece.CurrentX || pos.Rotation != piece.Rotation) && maxTry > 0 {
			//fmt.Println(pos.X, piece.CurrentX, "	", pos.Rotation, piece.Rotation)
			if pos.X < piece.CurrentX {
				np = piece.Left()
				if f.IsValid(np.Space) {
					//fmt.Println("valid")
					piece = np
					path = append(path, "left")
				}

			}

			if pos.X == piece.CurrentX {
				if pos.Rotation == 1 {
					np = piece.Turnright()
					if f.IsValid(np.Space) {
						piece = np
						path = append(path, "turnright")
					}
				}
				if pos.Rotation == 3 {

					np = piece.Turnleft()
					if f.IsValid(np.Space) {
						piece = np
						//fmt.Println(np.CurrentX, np.Rotation)
						//fmt.Println(piece.CurrentX, piece.Rotation)
						path = append(path, "turnleft")
					}

				}
			}

			if pos.X > piece.CurrentX {
				np := piece.Right()
				if f.IsValid(np.Space) {
					piece = np
					path = append(path, "right")
				}
			}
			maxTry--
		}

	}


			var buffer bytes.Buffer
		for i := 0; i < pos.Rotation; i++ {
			buffer.WriteString("turnright,")
		}
		if pos.Rotation == 1 {
			g.CurrentPieceX = g.CurrentPieceX + 1
			if g.CurrentPiece == "I" {
				g.CurrentPieceX = g.CurrentPieceX + 1
			}
		}
		if g.CurrentPieceX > pos.X {
			for i := 0; i < g.CurrentPieceX-pos.X; i++ {
				buffer.WriteString("left,")
			}
		}
		if g.CurrentPieceX < pos.X {
			for i := 0; i < pos.X-g.CurrentPieceX; i++ {
				buffer.WriteString("right,")
			}
		}
		buffer.WriteString("drop")
		fmt.Println(buffer.String())


	pos.Moves = "left,left" //  ???????????????????
	path = append(path, "drop")
	return strings.Join(path, ",")
}
*/
func (f Field) FixHole(piece Piece, hole Cell) ([]Position, int) {
	bag := &Bag{
		Options: make(map[int]*Piece),
		Hole:    hole,
	}
	bag.Options[piece.Key] = &piece

	f.Search("down", &piece, bag)
	f.Search("left", &piece, bag)
	f.Search("right", &piece, bag)

	f.Search("turnleft", &piece, bag)
	f.Search("turnright", &piece, bag)

	//fmt.Println(bag.Options)
	/*
		var keys []int
		for k := range bag.Options {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		// To perform the opertion you want
		for _, k := range keys {
			fmt.Println("Key:", k, "Value:", bag.Options[k])
		}
	*/
	return CheckFix(bag), bag.Total
}

func (f Field) Search(dir string, piece *Piece, bag *Bag) {
	bag.Total++
	var ok bool
	var el *Piece = nil
	//var np *Piece = nil

	switch dir {
	case "left":
		nextKey := piece.Key - 100
		el, ok = bag.Options[nextKey]
		if ok {
			if el != nil && len(piece.Moves)+1 < len(el.Moves) {
				bag.Options[nextKey].Moves = append(piece.Moves, dir)
				//f.Search("left", bag.Options[nextKey], bag)
				//f.Search("down", bag.Options[nextKey], bag)
				//f.Search("turnleft", bag.Options[nextKey], bag)
				//f.Search("turnright", bag.Options[nextKey], bag)
				//replace moves in added positions
			}
			return
		} else {
			np := piece.Left()
			if f.IsValid(&np.Space) {
				np.Moves = append(piece.Moves, dir)
				bag.Options[np.Key] = &np
				//AddFix(bag, &np)

				f.Search("down", &np, bag)
				f.Search("left", &np, bag)

				f.Search("turnleft", &np, bag)
				f.Search("turnright", &np, bag)

			} else {
				bag.Options[np.Key] = nil
			}
		}
	case "right":
		nextKey := piece.Key + 100
		el, ok = bag.Options[nextKey]
		if ok {
			if el != nil && len(piece.Moves)+1 < len(el.Moves) {
				bag.Options[nextKey].Moves = append(piece.Moves, dir)
				//f.Search("right", bag.Options[nextKey], bag)
				//f.Search("down", bag.Options[nextKey], bag)
				//f.Search("turnleft", bag.Options[nextKey], bag)
				//f.Search("turnright", bag.Options[nextKey], bag)
			}
			return
		} else {
			np := piece.Right()
			if f.IsValid(&np.Space) {
				np.Moves = append(piece.Moves, dir)
				bag.Options[np.Key] = &np
				//AddFix(bag, &np)

				f.Search("down", &np, bag)
				f.Search("right", &np, bag)

				f.Search("turnleft", &np, bag)
				f.Search("turnright", &np, bag)

			} else {
				bag.Options[np.Key] = nil
			}
		}
	case "down":
		nextKey := piece.Key - 1
		el, ok = bag.Options[nextKey]
		if ok {
			if el != nil && len(piece.Moves)+1 < len(el.Moves) {
				bag.Options[nextKey].Moves = append(piece.Moves, dir)
				//f.Search("left", bag.Options[nextKey], bag)
				//f.Search("right", bag.Options[nextKey], bag)
				//f.Search("turnleft", bag.Options[nextKey], bag)
				//f.Search("turnright", bag.Options[nextKey], bag)
			}
			return
		} else {
			np := piece.Down()
			if f.IsValid(&np.Space) {
				np.Moves = append(piece.Moves, dir)
				bag.Options[np.Key] = &np
				//AddFix(bag, &np)

				f.Search("left", &np, bag)
				f.Search("right", &np, bag)
				f.Search("turnleft", &np, bag)
				f.Search("turnright", &np, bag)

			} else {
				bag.Options[np.Key] = nil
			}
		}

	case "turnleft":
		np := piece.Turnleft()
		el, ok = bag.Options[np.Key]
		if ok {
			if el != nil && len(piece.Moves)+1 < len(el.Moves) {
				bag.Options[np.Key].Moves = append(piece.Moves, dir)
				//f.Search("left", bag.Options[np.Key], bag)
				//f.Search("right", bag.Options[np.Key], bag)
				//f.Search("down", bag.Options[np.Key], bag)
				//f.Search("turnleft", bag.Options[np.Key], bag)

			}
			return
		} else {
			if f.IsValid(&np.Space) {
				np.Moves = append(piece.Moves, dir)
				bag.Options[np.Key] = &np
				//AddFix(bag, &np)

				f.Search("down", &np, bag)
				f.Search("left", &np, bag)
				f.Search("right", &np, bag)

				f.Search("turnleft", &np, bag)

			} else {
				bag.Options[np.Key] = nil
			}
		}

	case "turnright":
		np := piece.Turnright()
		el, ok = bag.Options[np.Key]
		if ok {
			if el != nil && len(piece.Moves)+1 < len(el.Moves) {
				bag.Options[np.Key].Moves = append(piece.Moves, dir)
				//f.Search("left", bag.Options[np.Key], bag)
				//f.Search("right", bag.Options[np.Key], bag)
				//f.Search("down", bag.Options[np.Key], bag)
				//f.Search("turnright", bag.Options[np.Key], bag)

			}
			return
		} else {
			if f.IsValid(&np.Space) {
				np.Moves = append(piece.Moves, dir)
				bag.Options[np.Key] = &np
				//AddFix(bag, &np)

				f.Search("down", &np, bag)
				f.Search("left", &np, bag)
				f.Search("right", &np, bag)

				f.Search("turnright", &np, bag)

			} else {
				bag.Options[np.Key] = nil
			}
		}
	}
}

func CheckFix(bag *Bag) []Position {
	positions := make([]Position, 0)

	for _, piece := range bag.Options {
		if piece != nil {
			for _, cell := range piece.Space {
				if cell.X == bag.Hole.X && cell.Y == bag.Hole.Y {
					pos := Position{
						Rotation: piece.Rotation,
						X:        piece.CurrentX,
						Moves:    strings.Join(piece.Moves, ","),
					}
					positions = append(positions, pos)
				}
			}
		}
	}
	return positions
}

type Bag struct {
	Options map[int]*Piece
	Total   int
	Hole    Cell
}
