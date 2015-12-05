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

func (f Field) FindPositions(piece Piece) []Piece {
	positions := make([]Piece, 0)
	//countSearchCalls := 0

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
			//countSearchCalls++
			nkey = f.Search("down", k, bag)
			if nkey >= 0 {
				tmp[nkey] = false
			}
			//countSearchCalls++
			nkey = f.Search("left", k, bag)
			if nkey >= 0 {
				tmp[nkey] = false
			}
			//countSearchCalls++
			nkey = f.Search("right", k, bag)
			if nkey > 0 {
				tmp[nkey] = false
			}
			if piece.Name != "O" {
				//countSearchCalls++
				nkey = f.Search("turnleft", k, bag)
				if nkey >= 0 {
					tmp[nkey] = false
				}
				//countSearchCalls++
				nkey = f.Search("turnright", k, bag)
				if nkey >= 0 {
					tmp[nkey] = false
				}
			}
		}
		queue = tmp
	}
	//fmt.Println("countSearchCalls", countSearchCalls)
	//fmt.Println("bagLen", len(bag))
	for k, p := range bag {
		_, ok := bag[k-1]
		if !ok {
			if !f.Grid.isCollision(&p.Space, true) {
				p.FieldAfter = f.AfterHole(p.Space)
				p.Moves = strings.TrimPrefix(p.Moves, ",")
				positions = append(positions, *p)
			}
		}
	}
	return positions
}

func (f Field) AfterHole(space map[string]Cell) *Field {
	if len(space) != 4 {
		return nil
	}
	newGrid := f.Grid.Copy()
	for _, cell := range space {
		if newGrid[cell.Y][cell.X] {
			return nil
		} else {
			newGrid[cell.Y][cell.X] = true
		}
	}
	newField := newGrid.ToField()
	return &newField
}

func (f Field) Search(dir string, key int, bag map[int]*Piece) int {
	var ok bool
	var el *Piece
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

	if np.Name == "I" || np.Name == "S" || np.Name == "Z" {
		_, ok1 := bag[np.Key-20000]
		_, ok2 := bag[np.Key+20000]
		if ok1 || ok2 {
			return -1
		}
	}

	if f.Grid.isCollision(&np.Space, false) {
		return -1
	}

	np.Moves = nMoves
	bag[np.Key] = &np
	return np.Key
}
