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

func (f Field) FindPositions(piece Piece) []*Piece {
	p := &piece
	positions := make([]*Piece, 0)
	drop := p.CurrentY - f.MaxPick - 2
	if drop > 0 {
		p = p.Drop(drop)
	}

	stack := InitStack()
	stack.Push(p)
	for p != nil {
		f.Search(stack, p, down)
		f.Search(stack, p, left)
		f.Search(stack, p, right)
		if p.Name != O {
			f.Search(stack, p, turnleft)
			f.Search(stack, p, turnright)
		}
		p = stack.Pop()
	}

	for _, p := range stack.collection {
		if p.IsDown(stack) && !f.Grid.IsCollision(p.Space, true) {
			newGrid := f.Grid.ApplyPiece(p.Space)
			newField := newGrid.ToField()
			p.FieldAfter = &newField
			p.Moves = strings.TrimPrefix(p.Moves, ",")
			positions = append(positions, p)
		}
	}
	return positions
}

func (f Field) Search(stack *Stack, p *Piece, dir string) {
	nmCount := p.MovesCount + 1
	var ex, np *Piece
	switch dir {
	case left:
		nextKey := p.Key - 100
		if nextKey%10000/100 < 0 {
			return
		}
		ex = stack.Peek(nextKey)
		if ex != nil {
			ex.shorterPath(nmCount, dir)
			return
		}
		np = p.Left()
	case right:
		nextKey := p.Key + 100
		if nextKey%10000/100 > f.Width {
			return
		}
		ex = stack.Peek(nextKey)
		if ex != nil {
			ex.shorterPath(nmCount, dir)
			return
		}
		np = p.Right()
	case down:
		nextKey := p.Key - 1
		if nextKey%100 < 0 {
			return
		}
		ex = stack.Peek(nextKey)
		if ex != nil {
			ex.shorterPath(nmCount, dir)
			return
		}
		np = p.Down()
	case turnleft:
		np = p.Turnleft()
		ex = stack.Peek(np.Key)
		if ex != nil {
			ex.shorterPath(nmCount, dir)
			return
		}
	case turnright:
		np = p.Turnright()
		ex = stack.Peek(np.Key)
		if ex != nil {
			ex.shorterPath(nmCount, dir)
			return
		}
	}

	if f.Grid.IsCollision(np.Space, false) {
		return
	}
	stack.Push(np)
}
