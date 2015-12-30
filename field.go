package main

import (
	"fmt"
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
	Holes []*Cell
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
		f.Search(stack, p, left)
		f.Search(stack, p, right)
		if p.Name != O {
			f.Search(stack, p, turnleft)
			f.Search(stack, p, turnright)
		}
		f.Search(stack, p, down)
		p = stack.Pop()
	}

	for _, p := range stack.collection {
		if p.IsDown(stack) && !f.Grid.IsCollision(p.Space, true) {
			newGrid := f.Grid.ApplyPiece(p.Space)
			newField := newGrid.ToField()
			p.FieldAfter = &newField
			positions = append(positions, p)
		}
	}
	return positions
}

func (f Field) Search(stack *Stack, p *Piece, dir string) {
	nMoves := p.Moves + "," + dir
	var ex, np *Piece
	switch dir {
	case left:
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
	case right:
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
	case down:
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
	case turnleft:
		np = p.Turnleft()
		ex = stack.Peek(np.Key)
		if ex != nil {
			ex.shorterPath(nMoves)
			return
		}
	case turnright:
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
	np.Moves = nMoves
	stack.Push(np)
}

func (f Field) HideTspace() {
	for y, row := range f.Grid {
		x := -1
		isOneHole := false
		for i, col := range row {
			if !col {
				if x < 0 {
					x = i
					isOneHole = true
				} else {
					isOneHole = false
				}
			}
		}
		if isOneHole && f.Grid.isTshapeSpace(&Cell{X: x, Y: y}) {
			valid := true
			left := x - 2
			for left >= 0 {
				if !f.Grid[y+1][left] {
					valid = false
				}
				left--
			}
			right := x + 2
			for right < f.Width {
				if !f.Grid[y+1][right] {
					valid = false
				}
				right++
			}

			//cut the field
			if valid {
				fmt.Println("hide field")
				ng := f.Grid[y+2:]
				nf := ng.ToField()
				f = nf
			}
			/*
				if valid {
					if f.Grid[y+2][x-1] && !f[y+2][x] && !f[y+2][x+1] {
						return x
					}
					if !f[y+2][x-1] && !f[y+2][x] && f[y+2][x+1] {
						return x - 1
					}
				}*/
		}
	}
}
