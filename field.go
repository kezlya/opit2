package main

import (
	//"fmt"
	"strings"
)

type Field struct {
	Width   int
	Height  int
	Empty   int
	MaxPick int
	CountBH int
	CountFH int
	Burned  int

	Grid  Grid
	Picks Picks

	//TODO kill this after refactoring FixHoles
	HolesBlocked []Cell
	HolesFixable []Cell
}

func (f Field) Copy() Field {
	newField := f

	newGrid := make([][]bool, f.Height)
	for i, row := range f.Grid {
		newGrid[i] = make([]bool, f.Width)
		copy(newGrid[i], row[:])
	}
	newField.Grid = newGrid

	newPicks := make([]int, len(f.Picks))
	copy(newPicks, f.Picks[:])
	newField.Picks = newPicks

	//TODO if Holes property exist need to copy them as value type

	return newField
}

func (f Field) FindPositions(piece Piece) ([]Piece, int, int) {
	positions := make([]Piece, 0)
	countSearchCalls := 0

	drop := piece.CurrentY - f.MaxPick - 2
	if drop > 0 {
		piece = piece.Drop(drop)
	}

	bag := make(map[int]*Piece)
	queue := make(map[int]bool)
	nkey := 0

	bag[piece.Key] = &piece
	queue[piece.Key] = true

	for len(queue) > 0 {
		tmp := make(map[int]bool)
		//TODO impliment multithreading after bench it
		for k, _ := range queue {
			countSearchCalls++
			nkey = f.Search("down", k, bag)
			if nkey >= 0 {
				tmp[nkey] = false
			}
			countSearchCalls++
			nkey = f.Search("left", k, bag)
			if nkey >= 0 {
				tmp[nkey] = false
			}
			countSearchCalls++
			nkey = f.Search("right", k, bag)
			if nkey > 0 {
				tmp[nkey] = false
			}
			if piece.Name != "O" {
				countSearchCalls++
				nkey = f.Search("turnleft", k, bag)
				if nkey >= 0 {
					tmp[nkey] = false
				}
				countSearchCalls++
				nkey = f.Search("turnright", k, bag)
				if nkey >= 0 {
					tmp[nkey] = false
				}
			}
		}
		queue = tmp
	}
	bagLen := len(bag)

	for k, p := range bag {
		if p == nil {
			continue
		}
		el, ok := bag[k-1]
		if ok && el == nil {
			//check if piece position is over the top
			invalid := false
			for _, cell := range p.Space {
				if cell.Y >= f.Height {
					invalid = true
					break
				}
			}
			if !invalid {
				p.FieldAfter = f.AfterHole(p.Space)
				p.Moves = strings.TrimPrefix(p.Moves, ",")
				positions = append(positions, *p)
				//fmt.Println("   ", p.Key)
			}
		}
	}
	return positions, countSearchCalls, bagLen
}

//TODO kill this when ".After" removed
func (f Field) IsFit(pick, up int) bool {
	if pick+up <= f.Height {
		return true
	}
	return false
}

func (f Field) After(piece *Piece) (*Field, int) {
	picks := f.Picks
	x := piece.CurrentX
	y := 0
	r := piece.Rotation
	valid := false
	a := f.Copy()

	switch piece.Name {
	case "I":
		switch r {
		case 0, 2:
			if picks.IsRight(x, 3) {
				pick := picks.MaxR(x, 3)
				if f.IsFit(pick, 1) {
					a.Grid[pick][x] = true
					a.Grid[pick][x+1] = true
					a.Grid[pick][x+2] = true
					a.Grid[pick][x+3] = true
					y = pick
					valid = true
				}
			}
		case 1, 3:
			pick := picks[x]
			if f.IsFit(pick, 4) {
				a.Grid[pick][x] = true
				a.Grid[pick+1][x] = true
				a.Grid[pick+2][x] = true
				a.Grid[pick+3][x] = true
				y = pick
				valid = true
			}
		}
	case "J":
		switch r {
		case 0:
			if picks.IsRight(x, 2) {
				pick := picks.MaxR(x, 2)
				if f.IsFit(pick, 2) {
					a.Grid[pick][x] = true
					a.Grid[pick+1][x] = true
					a.Grid[pick][x+1] = true
					a.Grid[pick][x+2] = true
					y = pick
					valid = true
				}
			}
		case 1:
			if picks.IsRight(x, 1) {
				l := picks[x]
				l2 := picks[x+1]
				if l2 >= l+2 {
					if f.IsFit(l2, 1) {
						a.Grid[l2][x] = true
						a.Grid[l2][x+1] = true
						a.Grid[l2-1][x] = true
						a.Grid[l2-2][x] = true
						y = l2 - 2
						valid = true
					}
				} else {
					if f.IsFit(l, 3) {
						a.Grid[l][x] = true
						a.Grid[l+1][x] = true
						a.Grid[l+2][x] = true
						a.Grid[l+2][x+1] = true
						y = l
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
						a.Grid[pick+1][x] = true
						a.Grid[pick+1][x+1] = true
						a.Grid[pick+1][x+2] = true
						a.Grid[pick][x+2] = true
						y = pick
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a.Grid[pick][x] = true
						a.Grid[pick][x+1] = true
						a.Grid[pick][x+2] = true
						a.Grid[pick-1][x+2] = true
						y = pick - 1
						valid = true
					}
				}
			}
		case 3:
			if picks.IsRight(x, 1) {
				pick := picks.MaxR(x, 1)
				if f.IsFit(pick, 3) {
					a.Grid[pick][x] = true
					a.Grid[pick][x+1] = true
					a.Grid[pick+1][x+1] = true
					a.Grid[pick+2][x+1] = true
					y = pick
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
					a.Grid[pick][x] = true
					a.Grid[pick][x+1] = true
					a.Grid[pick][x+2] = true
					a.Grid[pick+1][x+2] = true
					y = pick
					valid = true
				}
			}
		case 1:
			if picks.IsRight(x, 1) {
				pick := picks.MaxR(x, 1)
				if f.IsFit(pick, 3) {
					a.Grid[pick][x] = true
					a.Grid[pick+1][x] = true
					a.Grid[pick+2][x] = true
					a.Grid[pick][x+1] = true
					y = pick
					valid = true
				}
			}
		case 2:
			if picks.IsRight(x, 2) {
				pick := picks.MaxR(x, 2)
				l := picks[x]
				if pick == l {
					if f.IsFit(l, 2) {
						a.Grid[pick][x] = true
						a.Grid[pick+1][x] = true
						a.Grid[pick+1][x+1] = true
						a.Grid[pick+1][x+2] = true
						y = pick
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a.Grid[pick-1][x] = true
						a.Grid[pick][x] = true
						a.Grid[pick][x+1] = true
						a.Grid[pick][x+2] = true
						y = pick - 1
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
						a.Grid[l][x] = true
						a.Grid[l][x+1] = true
						a.Grid[l-1][x+1] = true
						a.Grid[l-2][x+1] = true
						y = l - 2
						valid = true
					}
				} else {
					if f.IsFit(l2, 3) {
						a.Grid[l2+2][x] = true
						a.Grid[l2][x+1] = true
						a.Grid[l2+1][x+1] = true
						a.Grid[l2+2][x+1] = true
						y = l2
						valid = true
					}
				}
			}
		}
	case "O":
		if picks.IsRight(x, 1) {
			pick := picks.MaxR(x, 1)
			if f.IsFit(pick, 2) {
				a.Grid[pick][x] = true
				a.Grid[pick+1][x] = true
				a.Grid[pick][x+1] = true
				a.Grid[pick+1][x+1] = true
				y = pick
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
						a.Grid[pick][x] = true
						a.Grid[pick][x+1] = true
						a.Grid[pick+1][x+1] = true
						a.Grid[pick+1][x+2] = true
						y = pick
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a.Grid[pick-1][x] = true
						a.Grid[pick-1][x+1] = true
						a.Grid[pick][x+1] = true
						a.Grid[pick][x+2] = true
						y = pick - 1
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
						a.Grid[pick+2][x] = true
						a.Grid[pick+1][x] = true
						a.Grid[pick+1][x+1] = true
						a.Grid[pick][x+1] = true
						y = pick
						valid = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a.Grid[pick+1][x] = true
						a.Grid[pick][x] = true
						a.Grid[pick][x+1] = true
						a.Grid[pick-1][x+1] = true
						y = pick - 1
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
					a.Grid[pick][x] = true
					a.Grid[pick][x+1] = true
					a.Grid[pick+1][x+1] = true
					a.Grid[pick][x+2] = true
					y = pick
					valid = true
				}
			}
		case 1:
			if picks.IsRight(x, 1) {
				pick := picks.MaxR(x, 1)
				l := picks[x]
				if pick == l {
					if f.IsFit(pick, 3) {
						a.Grid[pick][x] = true
						a.Grid[pick+1][x] = true
						a.Grid[pick+1][x+1] = true
						a.Grid[pick+2][x] = true
						y = pick
						valid = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a.Grid[pick-1][x] = true
						a.Grid[pick][x] = true
						a.Grid[pick][x+1] = true
						a.Grid[pick+1][x] = true
						y = pick - 1
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
						a.Grid[pick+1][x] = true
						a.Grid[pick][x+1] = true
						a.Grid[pick+1][x+1] = true
						a.Grid[pick+1][x+2] = true
						y = pick
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a.Grid[pick][x] = true
						a.Grid[pick][x+1] = true
						a.Grid[pick-1][x+1] = true
						a.Grid[pick][x+2] = true
						y = pick - 1
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
						a.Grid[pick+2][x+1] = true
						a.Grid[pick+1][x] = true
						a.Grid[pick+1][x+1] = true
						a.Grid[pick][x+1] = true
						y = pick
						valid = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a.Grid[pick+1][x+1] = true
						a.Grid[pick][x] = true
						a.Grid[pick][x+1] = true
						a.Grid[pick-1][x+1] = true
						y = pick - 1
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
						a.Grid[pick+1][x] = true
						a.Grid[pick+1][x+1] = true
						a.Grid[pick][x+1] = true
						a.Grid[pick][x+2] = true
						y = pick
						valid = true
					}
				} else {
					if f.IsFit(pick, 1) {
						a.Grid[pick][x] = true
						a.Grid[pick][x+1] = true
						a.Grid[pick-1][x+1] = true
						a.Grid[pick-1][x+2] = true
						y = pick - 1
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
						a.Grid[pick][x] = true
						a.Grid[pick+1][x] = true
						a.Grid[pick+1][x+1] = true
						a.Grid[pick+2][x+1] = true
						y = pick
						valid = true
					}
				} else {
					if f.IsFit(pick, 2) {
						a.Grid[pick-1][x] = true
						a.Grid[pick][x] = true
						a.Grid[pick][x+1] = true
						a.Grid[pick+1][x+1] = true
						y = pick - 1
						valid = true
					}
				}
			}
		}
	}
	if valid {
		a.Burned = a.Grid.Burn()
		return &a, y
	}
	return nil, 0
}

func (f Field) AfterHole(space map[string]Cell) *Field {
	if len(space) != 4 {
		return nil
	}
	a := f.Copy()
	for _, cell := range space {
		if a.Grid[cell.Y][cell.X] {
			return nil
		} else {
			a.Grid[cell.Y][cell.X] = true
		}
	}
	a.Burned = a.Grid.Burn()
	return &a
}

func (f Field) IsValid(cells *map[string]Cell) bool {
	for _, c := range *cells {
		if c.X < 0 || c.X >= f.Width || c.Y < 0 {
			return false
		}
		if c.Y < f.Height && f.Grid[c.Y][c.X] {
			return false
		}
	}
	return true
}

func (f Field) ValidPosition(piece Piece) []Piece {
	validPieces := make([]Piece, 0)
	bag := make(map[int]*Piece)
	bag[piece.Key] = &piece
	queue := make(map[int]bool)
	queue[piece.Key] = true
	nkey := 0

	for len(queue) > 0 {
		tmp := make(map[int]bool)
		for k, _ := range queue {
			nkey = f.Search("left", k, bag)
			if nkey >= 0 {
				tmp[nkey] = false
			}
			nkey = f.Search("right", k, bag)
			if nkey > 0 {
				tmp[nkey] = false
			}
			if piece.Name != "O" {
				nkey = f.Search("turnleft", k, bag)
				if nkey >= 0 {
					tmp[nkey] = false
				}
				nkey = f.Search("turnright", k, bag)
				if nkey >= 0 {
					tmp[nkey] = false
				}
			}
		}
		queue = tmp
	}

	tempBag := make(map[int]*Piece)
	for k, p := range bag {
		if p == nil {
			delete(bag, k)
			continue
		}
		if (p.Name == "I" || p.Name == "Z" || p.Name == "S") &&
			(p.Rotation == 3 || p.Rotation == 2) {
			_, ok := bag[k-20000]
			if ok {
				delete(bag, k)
				continue
			}
			_, ok = bag[k-20000+1]
			if ok {
				delete(bag, k)
				continue
			}
		}
		if (p.Name == "I" && p.Rotation != 0) ||
			(p.Name == "L" && p.Rotation != 2) ||
			(p.Name == "J" && p.Rotation != 2) ||
			(p.Name == "T" && p.Rotation != 2) {
			tempBag[k] = p
			nkey = f.Search("down", k, tempBag)
			if nkey == -1 {
				delete(bag, k)
				continue
			}
		}
		fieldAfter, y := f.After(p)
		if fieldAfter == nil {
			delete(bag, k)
			continue
		}
		np := p.DropTo(y)
		np.FieldAfter = fieldAfter
		np.Moves = strings.TrimPrefix(p.Moves, ",")
		np.IsHole = false
		validPieces = append(validPieces, np)
	}

	return validPieces
}

func (f Field) FixHoles(piece Piece) []Piece {
	holes := f.HolesFixable
	fixes := make([]Piece, 0)
	bag := make(map[int]*Piece)
	queue := make(map[int]bool)
	nkey := 0

	drop := f.Height - f.Empty - piece.CurrentY - 1
	fp := piece.Drop(drop)
	bag[fp.Key] = &fp
	queue[fp.Key] = true

	for len(queue) > 0 {
		tmp := make(map[int]bool)
		for k, _ := range queue {
			nkey = f.Search("down", k, bag)
			if nkey >= 0 {
				tmp[nkey] = false
			}
			nkey = f.Search("left", k, bag)
			if nkey >= 0 {
				tmp[nkey] = false
			}
			nkey = f.Search("right", k, bag)
			// move to the right woulf never generate 0 key
			// 0 key is left bottom conner
			if nkey > 0 {
				tmp[nkey] = false
			}
			if piece.Name != "O" {
				nkey = f.Search("turnleft", k, bag)
				if nkey >= 0 {
					tmp[nkey] = false
				}
				nkey = f.Search("turnright", k, bag)
				if nkey >= 0 {
					tmp[nkey] = false
				}
			}
		}
		queue = tmp
	}
	found := false
	invalid := false
	maxY := f.Height

	for k, p := range bag {
		//fmt.Println(k)
		if p == nil {
			delete(bag, k)
			continue
		}
		el, ok := bag[k-1]
		if ok && el != nil {
			delete(bag, k)
			continue
		}
		if (p.Name == "I" || p.Name == "Z" || p.Name == "S") &&
			(p.Rotation == 3 || p.Rotation == 2) {
			el, ok := bag[k-20000]
			if ok && el != nil {
				delete(bag, k)
				continue
			}
			el, ok = bag[k-20000-1]
			if ok && el != nil {
				delete(bag, k)
				continue
			}
		}

		found = false
		invalid = false
		for _, hole := range holes {
			for _, cell := range p.Space {
				if cell.Y >= maxY {
					invalid = true
					break
				}
				if cell.X == hole.X && cell.Y == hole.Y {
					found = true
				}
			}
			if found && !invalid {
				p.FieldAfter = f.AfterHole(p.Space)
				p.Moves = strings.TrimPrefix(p.Moves, ",")
				p.IsHole = true
				fixes = append(fixes, *p)
				break
			}
		}
	}
	return fixes
}

func (f Field) Search(dir string, key int, bag map[int]*Piece) int {
	var ok bool
	var el *Piece = nil
	var np Piece
	nMoves := bag[key].Moves + "," + dir

	switch dir {
	case "left":
		nextKey := key - 100
		if nextKey%10000/100 < 0 {
			return -1
		}
		el, ok = bag[nextKey]
		if ok {
			el.shorterPath(nMoves)
			return -1
		}
		np = bag[key].Left()
	case "right":
		nextKey := key + 100
		if nextKey%10000/100 > f.Width {
			return -1
		}
		el, ok = bag[nextKey]
		if ok {
			el.shorterPath(nMoves)
			return -1
		}
		np = bag[key].Right()
	case "down":
		nextKey := key - 1
		if nextKey%100 < 0 {
			return -1
		}
		el, ok = bag[nextKey]
		if ok {
			el.shorterPath(nMoves)
			return -1
		}
		np = bag[key].Down()
	case "turnleft":
		np = bag[key].Turnleft()
		el, ok = bag[np.Key]
		if ok {
			el.shorterPath(nMoves)
			return -1
		}
	case "turnright":
		np = bag[key].Turnright()
		el, ok = bag[np.Key]
		if ok {
			el.shorterPath(nMoves)
			return -1
		}
	}

	if f.IsValid(&np.Space) {
		np.Moves = nMoves
		//TODO refactor this do not Call search on this keys
		// I S Z 0==2 and 1 == 3
		if np.Name == "I" || np.Name == "S" || np.Name == "Z" {
			_, ok1 := bag[np.Key-20000]
			_, ok2 := bag[np.Key+20000]
			if ok1 || ok2 {
				return -1
			}
		}
		bag[np.Key] = &np
		return np.Key
	}
	bag[np.Key] = nil
	return -1
}
