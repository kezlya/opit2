package main

import (
//"fmt"
//"strings"
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
	Picks []int
}

func (f Field) FindPositions(piece Piece) []Piece {
	positions := make([]Piece, 0)
	drop := piece.CurrentY - f.MaxPick - 2
	if drop > 0 {
		piece = piece.Drop(drop)
	}

	stack := InitStack()
	stack.Push(&piece)
	//fmt.Println("stack collection", stack.collection)
	for stack.Len() > 0 {
		p := stack.Pop()
		f.Search(stack, p, "down")
		f.Search(stack, p, "left")
		f.Search(stack, p, "right")
		if piece.Name != "O" {
			f.Search(stack, p, "turnleft")
			f.Search(stack, p, "turnright")
		}
	}

	bag := stack.collection
	for k, p := range bag {
		_, ok := bag[k-1]
		if !ok && !f.Grid.IsCollision(p.Space, true) {
			newGrid := f.Grid.ApplyPiece(p.Space)
			newField := newGrid.ToField()
			p.FieldAfter = &newField
			//p.Moves = strings.TrimPrefix(p.Moves, ",")
			positions = append(positions, *p)
		}
	}
	return positions
}

func (f Field) Search(stack *Stack, p *Piece, dir string) {
	//p.Moves + "," + dir
	var np Piece
	switch dir {
	case "left":
		nextKey := p.Key - 100
		if nextKey%10000/100 < 0 {
			return
		}
		if stack.Exist(nextKey) {
			//el.shorterPath(nMoves)
			return
		}
		np = p.Left()
	case "right":
		nextKey := p.Key + 100
		if nextKey%10000/100 > f.Width {
			return
		}
		if stack.Exist(nextKey) {
			//el.shorterPath(nMoves)
			return
		}
		np = p.Right()
	case "down":
		nextKey := p.Key - 1
		if nextKey%100 < 0 {
			return
		}
		if stack.Exist(nextKey) {
			//el.shorterPath(nMoves)
			return
		}
		np = p.Down()
	case "turnleft":
		np = p.Turnleft()
		if stack.Exist(np.Key) {
			//el.shorterPath(nMoves)
			return
		}
	case "turnright":
		np = p.Turnright()
		if stack.Exist(np.Key) {
			//el.shorterPath(nMoves)
			return
		}
	}

	if np.Name == "I" || np.Name == "S" || np.Name == "Z" {
		if stack.Exist(np.Key-20000) || stack.Exist(np.Key+20000) {
			return
		}
	}

	if f.Grid.IsCollision(np.Space, false) {
		return
	}

	//np.Moves = nMoves
	//bag[np.Key] = &np
	//fmt.Println(np)
	stack.Push(&np)
}
