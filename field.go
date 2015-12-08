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
	//fmt.Println(bag)
	for _, p := range bag {
		if p.IsDown(stack) && !f.Grid.IsCollision(p.Space, true) {
			//fmt.Println(k)
			newGrid := f.Grid.ApplyPiece(p.Space)
			newField := newGrid.ToField()
			p.FieldAfter = &newField
			p.Moves = strings.TrimPrefix(p.Moves, ",")
			positions = append(positions, *p)
		}
	}
	return positions
}

func (f Field) Search(stack *Stack, p *Piece, dir string) {
	nMoves := p.Moves + "," + dir
	var np Piece
	var ex *Piece

	switch dir {
	case "left":
		nextKey := p.Key - 100
		if nextKey%10000/100 < 0 {
			return
		}
		ex = stack.Peek(nextKey)
		if ex != nil {
			ex.shorterPath(nMoves)
			return
		}
		np = p.Left()
	case "right":
		nextKey := p.Key + 100
		if nextKey%10000/100 > f.Width {
			return
		}
		ex = stack.Peek(nextKey)
		if ex != nil {
			ex.shorterPath(nMoves)
			return
		}
		np = p.Right()
	case "down":
		nextKey := p.Key - 1
		if nextKey%100 < 0 {
			return
		}
		ex = stack.Peek(nextKey)
		if ex != nil {
			ex.shorterPath(nMoves)
			return
		}
		np = p.Down()
	case "turnleft":
		np = p.Turnleft()
		ex = stack.Peek(np.Key)
		if ex != nil {
			ex.shorterPath(nMoves)
			return
		}
	case "turnright":
		np = p.Turnright()
		ex = stack.Peek(np.Key)
		if ex != nil {
			ex.shorterPath(nMoves)
			return
		}
	}

	if f.Grid.IsCollision(np.Space, false) {
		return
	}

	stack.Push(&np)
}
