package main

import (
	//"fmt"
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

func (f Field) Positions(piece Piece, st Strategy) []Position {
	positions := make([]Position, 0)
	picks := f.Picks()
	hBlocked, hFixable := f.FindHoles(picks)

	validPieces := f.ValidPosition(piece)

	for _, validPiece := range validPieces {
		fieldAfter := f.After(validPiece.CurrentX, validPiece.Rotation, piece.Name)
		if fieldAfter != nil {
			//fmt.Println(validPiece.Rotation, validPiece.CurrentX)
			p := Position{}
			p.Init(picks, fieldAfter, hBlocked, st)
			p.Moves = strings.TrimPrefix(validPiece.Moves, ",")
			positions = append(positions, p)
		}
	}
	if len(hFixable) > 0 && (piece.Name == "I" || piece.Name == "T" || piece.Name == "J") {
		//fmt.Println(len(hFixable))
		fixes := f.FixHoles(piece, hFixable)
		for _, fix := range fixes {
			//fmt.Println(fix.Key, fix.Moves)
			p := Position{}
			p.Init(picks, f.AfterHole(fix.Space), hBlocked, st)
			p.Moves = strings.TrimPrefix(fix.Moves, ",")
			positions = append(positions, p)
		}
	}
	return positions
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

func (f Field) FindHoles(picks Picks) ([]Cell, []Cell) {
	blocked := make([]Cell, 0)
	fixable := make([]Cell, 0)
	for i, pick := range picks {
		for j := 0; j < pick; j++ {
			if !f[j][i] {
				hole := Cell{X: i, Y: j}
				if (i-2 > -1 && !f[j][i-1] && !f[j][i-2]) ||
					(i+2 < f.Width() && !f[j][i+1] && !f[j][i+2]) {
					fixable = append(fixable, hole)
				} else {
					blocked = append(blocked, hole)
				}
			}
		}
	}
	return blocked, fixable
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
		case 0, 2:
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
		case 1, 3:
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
		case 0, 2:
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
		case 1, 3:
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
		case 0, 2:
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
		case 1, 3:
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

func (f Field) AfterHole(space map[string]Cell) Field {
	if len(space) != 4 {
		return nil
	}

	w := f.Width()
	a := make([][]bool, f.Height())
	for i, row := range f {
		a[i] = make([]bool, w)
		copy(a[i], row[:])
	}
	for _, cell := range space {
		if a[cell.Y][cell.X] {
			return nil
		} else {
			a[cell.Y][cell.X] = true
		}
	}
	return a
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

func (f Field) ValidPosition(piece Piece) []Piece {
	validPieces := make([]Piece, 0)
	bag := &Bag{Options: make(map[int]*Piece)}
	bag.Options[piece.Key] = &piece
	queue := make(map[int]bool)
	queue[piece.Key] = true
	nkey := 0

	for len(queue) > 0 {
		tmp := make(map[int]bool)
		for k, _ := range queue {
			nkey = f.Search("left", k, bag)
			if nkey > 0 {
				tmp[nkey] = false
			}
			nkey = f.Search("right", k, bag)
			if nkey > 0 {
				tmp[nkey] = false
			}
			if piece.Name != "O" {
				nkey = f.Search("turnleft", k, bag)
				if nkey > 0 {
					tmp[nkey] = false
				}
				nkey = f.Search("turnright", k, bag)
				if nkey > 0 {
					tmp[nkey] = false
				}
			}
		}
		queue = tmp
	}

	tempBag := &Bag{Options: make(map[int]*Piece)}
	for k, p := range bag.Options {
		if p == nil {
			delete(bag.Options, k)
			continue
		}
		if (p.Name == "I" || p.Name == "Z" || p.Name == "S") &&
			(p.Rotation == 3 || p.Rotation == 2) {
			_, ok := bag.Options[k-20000]
			if ok {
				delete(bag.Options, k)
				continue
			}
			_, ok = bag.Options[k-20000+1]
			if ok {
				delete(bag.Options, k)
				continue
			}
		}
		if (p.Name == "I" && p.Rotation != 0) ||
			(p.Name == "L" && p.Rotation != 2) ||
			(p.Name == "J" && p.Rotation != 2) ||
			(p.Name == "T" && p.Rotation != 2) {
			tempBag.Options[k] = p
			nkey = f.Search("down", k, tempBag)
			if nkey == 0 {
				delete(bag.Options, k)
				continue
			}
		}
		validPieces = append(validPieces, *p)
	}

	return validPieces
}

func (f Field) FixHoles(piece Piece, holes []Cell) []Piece {
	fixes := make([]Piece, 0)
	bag := &Bag{Options: make(map[int]*Piece)}
	bag.Options[piece.Key] = &piece
	queue := make(map[int]bool)
	queue[piece.Key] = true
	nkey := 0

	for len(queue) > 0 {
		tmp := make(map[int]bool)
		for k, _ := range queue {
			nkey = f.Search("down", k, bag)
			if nkey > 0 {
				tmp[nkey] = false
			}
			nkey = f.Search("left", k, bag)
			if nkey > 0 {
				tmp[nkey] = false
			}
			nkey = f.Search("right", k, bag)
			if nkey > 0 {
				tmp[nkey] = false
			}
			if piece.Name != "O" {
				nkey = f.Search("turnleft", k, bag)
				if nkey > 0 {
					tmp[nkey] = false
				}
				nkey = f.Search("turnright", k, bag)
				if nkey > 0 {
					tmp[nkey] = false
				}
			}
		}
		queue = tmp
	}
	stop := false
	for k, p := range bag.Options {
		if p == nil {
			delete(bag.Options, k)
			continue
		}
		el, ok := bag.Options[k-1]
		if ok && el != nil {
			delete(bag.Options, k)
			continue
		}
		stop = false
		for _, hole := range holes {
			for _, cell := range p.Space {
				if cell.X == hole.X && cell.Y == hole.Y {
					fixes = append(fixes, *p)
					stop = true
					break
				}
			}
			if stop {
				break
			}
		}
	}
	return fixes
}

func (f Field) Search(dir string, key int, bag *Bag) int {
	bag.Total++
	var ok bool
	var el *Piece = nil
	nMoves := bag.Options[key].Moves + "," + dir

	switch dir {
	case "left":
		nextKey := key - 100
		el, ok = bag.Options[nextKey]
		if ok {
			if el != nil && len(nMoves) < len(el.Moves) {
				bag.Options[nextKey].Moves = nMoves
				return 0
			}
			return 0
		}
		np := bag.Options[key].Left()
		if f.IsValid(&np.Space) {
			np.Moves = nMoves
			bag.Options[np.Key] = &np
			return np.Key
		}
		bag.Options[np.Key] = nil
		return 0

	case "right":
		nextKey := key + 100
		el, ok = bag.Options[nextKey]
		if ok {
			if el != nil && len(nMoves) < len(el.Moves) {
				bag.Options[nextKey].Moves = nMoves
				return 0
			}
			return 0
		}
		np := bag.Options[key].Right()
		if f.IsValid(&np.Space) {
			np.Moves = nMoves
			bag.Options[np.Key] = &np
			return np.Key
		}
		bag.Options[np.Key] = nil
		return 0
	case "down":
		nextKey := key - 1
		el, ok = bag.Options[nextKey]
		if ok {
			if el != nil && len(nMoves) < len(el.Moves) {
				bag.Options[nextKey].Moves = nMoves
				return 0
			}
			return 0
		}
		np := bag.Options[key].Down()
		if f.IsValid(&np.Space) {
			np.Moves = nMoves
			bag.Options[np.Key] = &np
			return np.Key
		}
		bag.Options[np.Key] = nil
		return 0
	case "turnleft":
		np := bag.Options[key].Turnleft()
		el, ok = bag.Options[np.Key]
		if ok {
			if el != nil && len(nMoves) < len(el.Moves) {
				bag.Options[np.Key].Moves = nMoves
				return 0
			}
			return 0
		}
		if f.IsValid(&np.Space) {
			np.Moves = nMoves
			bag.Options[np.Key] = &np
			return np.Key
		}
		bag.Options[np.Key] = nil
		return 0
	case "turnright":
		np := bag.Options[key].Turnright()
		el, ok = bag.Options[np.Key]
		if ok {
			if el != nil && len(nMoves) < len(el.Moves) {
				bag.Options[np.Key].Moves = nMoves
				return 0
			}
			return 0
		}
		if f.IsValid(&np.Space) {
			np.Moves = nMoves
			bag.Options[np.Key] = &np
			return np.Key
		}
		bag.Options[np.Key] = nil
		return 0
	}
	return 0
}
