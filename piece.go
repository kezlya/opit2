package main

import (
	//"fmt"
	"log"
)

type Piece struct {
	Name     string
	Moves    string
	Key      int
	CurrentX int
	CurrentY int
	Rotation int

	Tspin        bool
	Tspin2       bool
	PerfectClear bool

	Space      map[string]Cell
	FieldAfter *Field
	Score      *Score
}

type Score struct {
	Burn   int
	Step   int
	BHoles int
	FHoles int
	CHoles int
	HighY  int
	Total  int
	NScore int
}

type Cell struct {
	X int
	Y int
}

func InitPiece(name string, x, y int) Piece {
	piece := Piece{
		Name:     name,
		CurrentX: x,
		CurrentY: y,
		Score:    &Score{},
	}
	piece.Space = make(map[string]Cell, 4)
	switch name {
	case I:
		piece.Space["m1"] = Cell{X: x, Y: y}
		piece.Space["m2"] = Cell{X: x + 1, Y: y}
		piece.Space["m3"] = Cell{X: x + 2, Y: y}
		piece.Space["m4"] = Cell{X: x + 3, Y: y}
	case J:
		piece.Space["m1"] = Cell{X: x, Y: y}
		piece.Space["m2"] = Cell{X: x + 1, Y: y}
		piece.Space["m3"] = Cell{X: x + 2, Y: y}
		piece.Space["t1"] = Cell{X: x, Y: y + 1}
	case L:
		piece.Space["m1"] = Cell{X: x, Y: y}
		piece.Space["m2"] = Cell{X: x + 1, Y: y}
		piece.Space["m3"] = Cell{X: x + 2, Y: y}
		piece.Space["t3"] = Cell{X: x + 2, Y: y + 1}
	case O:
		piece.Space["t1"] = Cell{X: x, Y: y + 1}
		piece.Space["t2"] = Cell{X: x + 1, Y: y + 1}
		piece.Space["m1"] = Cell{X: x, Y: y}
		piece.Space["m2"] = Cell{X: x + 1, Y: y}
	case S:
		piece.Space["t2"] = Cell{X: x + 1, Y: y + 1}
		piece.Space["t3"] = Cell{X: x + 2, Y: y + 1}
		piece.Space["m1"] = Cell{X: x, Y: y}
		piece.Space["m2"] = Cell{X: x + 1, Y: y}
	case T:
		piece.Space["m1"] = Cell{X: x, Y: y}
		piece.Space["m2"] = Cell{X: x + 1, Y: y}
		piece.Space["m3"] = Cell{X: x + 2, Y: y}
		piece.Space["t2"] = Cell{X: x + 1, Y: y + 1}
	case Z:
		piece.Space["t1"] = Cell{X: x, Y: y + 1}
		piece.Space["t2"] = Cell{X: x + 1, Y: y + 1}
		piece.Space["m2"] = Cell{X: x + 1, Y: y}
		piece.Space["m3"] = Cell{X: x + 2, Y: y}
	default:
		log.Fatalln(name, "piece is not supported")
	}
	piece.setKey()
	return piece
}

func (p *Piece) Left() *Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X - 1, Y: v.Y}
	}
	np := Piece{
		Name:     p.Name,
		Rotation: p.Rotation,
		CurrentX: p.CurrentX - 1,
		CurrentY: p.CurrentY,
		Space:    res,
		Score:    &Score{},
	}
	np.setKey()
	return &np
}

func (p *Piece) Right() *Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X + 1, Y: v.Y}
	}
	np := Piece{
		Name:     p.Name,
		Rotation: p.Rotation,
		CurrentX: p.CurrentX + 1,
		CurrentY: p.CurrentY,
		Space:    res,
		Score:    &Score{},
	}
	np.setKey()
	return &np
}

func (p *Piece) Down() *Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X, Y: v.Y - 1}
	}
	np := Piece{
		Name:     p.Name,
		Rotation: p.Rotation,
		CurrentX: p.CurrentX,
		CurrentY: p.CurrentY - 1,
		Space:    res,
		Score:    &Score{}}
	np.setKey()
	return &np
}

func (p *Piece) Drop(n int) *Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X, Y: v.Y - n}
	}
	np := Piece{
		Name:     p.Name,
		Rotation: p.Rotation,
		CurrentX: p.CurrentX,
		CurrentY: p.CurrentY - n,
		Space:    res,
		Score:    &Score{},
	}
	np.setKey()
	for i := 0; i < n; i++ {
		np.Moves += ",down"
	}
	return &np
}

func (p *Piece) Turnright() *Piece {
	np := Piece{
		Name:     p.Name,
		Score:    &Score{},
		Rotation: p.Rotation + 1,
	}

	if np.Rotation > 3 {
		np.Rotation = 0
	}

	sp := make(map[string]Cell, 4)
	for i, v := range p.Space {
		sp[i] = v
	}
	switch p.Name {
	case I:
		switch p.Rotation {
		case 0:
			t3 := Cell{X: sp["m1"].X + 2, Y: sp["m1"].Y + 1}
			b3 := Cell{X: sp["m2"].X + 1, Y: sp["m2"].Y - 1}
			x3 := Cell{X: sp["m4"].X - 1, Y: sp["m4"].Y - 2}
			delete(sp, "m1")
			delete(sp, "m2")
			delete(sp, "m4")
			sp["t3"] = t3
			sp["b3"] = b3
			sp["x3"] = x3
			np.CurrentX = x3.X
			np.CurrentY = x3.Y
		case 1:
			b1 := Cell{X: sp["t3"].X - 2, Y: sp["t3"].Y - 2}
			b2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y - 1}
			b4 := Cell{X: sp["x3"].X + 1, Y: sp["x3"].Y + 1}
			delete(sp, "t3")
			delete(sp, "m3")
			delete(sp, "x3")
			sp["b1"] = b1
			sp["b2"] = b2
			sp["b4"] = b4
			np.CurrentX = b1.X
			np.CurrentY = b1.Y
		case 2:
			t2 := Cell{X: sp["b1"].X + 1, Y: sp["b1"].Y + 2}
			m2 := Cell{X: sp["b3"].X - 1, Y: sp["b3"].Y + 1}
			x2 := Cell{X: sp["b4"].X - 2, Y: sp["b4"].Y - 1}
			delete(sp, "b1")
			delete(sp, "b3")
			delete(sp, "b4")
			sp["t2"] = t2
			sp["m2"] = m2
			sp["x2"] = x2
			np.CurrentX = x2.X
			np.CurrentY = x2.Y
		case 3:
			m1 := Cell{X: sp["x2"].X - 1, Y: sp["x2"].Y + 2}
			m3 := Cell{X: sp["b2"].X + 1, Y: sp["b2"].Y + 1}
			m4 := Cell{X: sp["t2"].X + 2, Y: sp["t2"].Y + -1}
			delete(sp, "x2")
			delete(sp, "b2")
			delete(sp, "t2")
			sp["m1"] = m1
			sp["m3"] = m3
			sp["m4"] = m4
			np.CurrentX = m1.X
			np.CurrentY = m1.Y
		}
	case J:
		switch p.Rotation {
		case 0:
			b2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y - 1}
			t2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y + 1}
			t3 := Cell{X: sp["t1"].X + 2, Y: sp["t1"].Y}
			delete(sp, "m3")
			delete(sp, "m1")
			delete(sp, "t1")
			sp["b2"] = b2
			sp["t2"] = t2
			sp["t3"] = t3
			np.CurrentX = b2.X
			np.CurrentY = b2.Y
		case 1:
			m1 := Cell{X: sp["b2"].X - 1, Y: sp["b2"].Y + 1}
			m3 := Cell{X: sp["t2"].X + 1, Y: sp["t2"].Y - 1}
			b3 := Cell{X: sp["t3"].X, Y: sp["t3"].Y - 2}
			delete(sp, "b2")
			delete(sp, "t2")
			delete(sp, "t3")
			sp["m1"] = m1
			sp["m3"] = m3
			sp["b3"] = b3
			np.CurrentX = m1.X
			np.CurrentY = b3.Y
		case 2:
			t2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y + 1}
			b2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y - 1}
			b1 := Cell{X: sp["b3"].X - 2, Y: sp["b3"].Y}
			delete(sp, "m1")
			delete(sp, "m3")
			delete(sp, "b3")
			sp["t2"] = t2
			sp["b2"] = b2
			sp["b1"] = b1
			np.CurrentX = b1.X
			np.CurrentY = b1.Y
		case 3:
			m1 := Cell{X: sp["b2"].X - 1, Y: sp["b2"].Y + 1}
			t1 := Cell{X: sp["b1"].X, Y: sp["b1"].Y + 2}
			m3 := Cell{X: sp["t2"].X + 1, Y: sp["t2"].Y - 1}
			delete(sp, "b2")
			delete(sp, "b1")
			delete(sp, "t2")
			sp["m1"] = m1
			sp["t1"] = t1
			sp["m3"] = m3
			np.CurrentX = m1.X
			np.CurrentY = m1.Y
		}
	case L:
		switch p.Rotation {
		case 0:
			t2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y + 1}
			b2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y - 1}
			b3 := Cell{X: sp["t3"].X, Y: sp["t3"].Y - 2}
			delete(sp, "m1")
			delete(sp, "m3")
			delete(sp, "t3")
			sp["t2"] = t2
			sp["b2"] = b2
			sp["b3"] = b3
			np.CurrentX = b2.X
			np.CurrentY = b2.Y
		case 1:
			m1 := Cell{X: sp["b2"].X - 1, Y: sp["b2"].Y + 1}
			m3 := Cell{X: sp["t2"].X + 1, Y: sp["t2"].Y - 1}
			b1 := Cell{X: sp["b3"].X - 2, Y: sp["b3"].Y}
			delete(sp, "b2")
			delete(sp, "t2")
			delete(sp, "b3")
			sp["m1"] = m1
			sp["m3"] = m3
			sp["b1"] = b1
			np.CurrentX = b1.X
			np.CurrentY = b1.Y
		case 2:
			t2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y + 1}
			b2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y - 1}
			t1 := Cell{X: sp["b1"].X, Y: sp["b1"].Y + 2}
			delete(sp, "m1")
			delete(sp, "m3")
			delete(sp, "b1")
			sp["t2"] = t2
			sp["b2"] = b2
			sp["t1"] = t1
			np.CurrentX = t1.X
			np.CurrentY = b2.Y
		case 3:
			m1 := Cell{X: sp["b2"].X - 1, Y: sp["b2"].Y + 1}
			m3 := Cell{X: sp["t2"].X + 1, Y: sp["t2"].Y - 1}
			t3 := Cell{X: sp["t1"].X + 2, Y: sp["t1"].Y}
			delete(sp, "b2")
			delete(sp, "t2")
			delete(sp, "t1")
			sp["m1"] = m1
			sp["m3"] = m3
			sp["t3"] = t3
			np.CurrentX = m1.X
			np.CurrentY = m1.Y
		}
	case T:
		switch p.Rotation {
		case 0:
			b2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y - 1}
			delete(sp, "m1")
			sp["b2"] = b2
			np.CurrentX = b2.X
			np.CurrentY = b2.Y
		case 1:
			m1 := Cell{X: sp["t2"].X - 1, Y: sp["t2"].Y - 1}
			delete(sp, "t2")
			sp["m1"] = m1
			np.CurrentX = m1.X
			np.CurrentY = p.CurrentY
		case 2:
			t2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y + 1}
			delete(sp, "m3")
			sp["t2"] = t2
			np.CurrentX = p.CurrentX
			np.CurrentY = p.CurrentY
		case 3:
			m3 := Cell{X: sp["b2"].X + 1, Y: sp["b2"].Y + 1}
			delete(sp, "b2")
			sp["m3"] = m3
			np.CurrentX = p.CurrentX
			np.CurrentY = m3.Y
		}
	case S:
		switch p.Rotation {
		case 0:
			b3 := Cell{X: sp["m1"].X + 2, Y: sp["m1"].Y - 1}
			m3 := Cell{X: sp["t3"].X, Y: sp["t3"].Y - 1}
			delete(sp, "m1")
			delete(sp, "t3")
			sp["b3"] = b3
			sp["m3"] = m3
			np.CurrentX = sp["m2"].X
			np.CurrentY = b3.Y
		case 1:
			b1 := Cell{X: sp["t2"].X - 1, Y: sp["t2"].Y - 2}
			b2 := Cell{X: sp["b3"].X - 1, Y: sp["b3"].Y}
			delete(sp, "t2")
			delete(sp, "b3")
			sp["b1"] = b1
			sp["b2"] = b2
			np.CurrentX = b1.X
			np.CurrentY = b1.Y
		case 2:
			t1 := Cell{X: sp["m3"].X - 2, Y: sp["m3"].Y + 1}
			m1 := Cell{X: sp["b1"].X, Y: sp["b1"].Y + 1}
			delete(sp, "m3")
			delete(sp, "b1")
			sp["t1"] = t1
			sp["m1"] = m1
			np.CurrentX = t1.X
			np.CurrentY = sp["b2"].Y
		case 3:
			t3 := Cell{X: sp["b2"].X + 1, Y: sp["b2"].Y + 2}
			t2 := Cell{X: sp["t1"].X + 1, Y: sp["t1"].Y}
			delete(sp, "b2")
			delete(sp, "t1")
			sp["t3"] = t3
			sp["t2"] = t2
			np.CurrentX = sp["m1"].X
			np.CurrentY = sp["m1"].Y
		}
	case Z:
		switch p.Rotation {
		case 0:
			b2 := Cell{X: sp["t1"].X + 1, Y: sp["t1"].Y - 2}
			t3 := Cell{X: sp["t2"].X + 1, Y: sp["t2"].Y}
			delete(sp, "t1")
			delete(sp, "t2")
			sp["b2"] = b2
			sp["t3"] = t3
			np.CurrentX = b2.X
			np.CurrentY = b2.Y
		case 1:
			m1 := Cell{X: sp["t3"].X - 2, Y: sp["t3"].Y - 1}
			b3 := Cell{X: sp["m3"].X, Y: sp["m3"].Y - 1}
			delete(sp, "t3")
			delete(sp, "m3")
			sp["m1"] = m1
			sp["b3"] = b3
			np.CurrentX = m1.X
			np.CurrentY = b3.Y
		case 2:
			t2 := Cell{X: sp["b3"].X - 1, Y: sp["b3"].Y + 2}
			b1 := Cell{X: sp["b2"].X - 1, Y: sp["b2"].Y}
			delete(sp, "b3")
			delete(sp, "b2")
			sp["t2"] = t2
			sp["b1"] = b1
			np.CurrentX = b1.X
			np.CurrentY = b1.Y
		case 3:
			m3 := Cell{X: sp["b1"].X + 2, Y: sp["b1"].Y + 1}
			t1 := Cell{X: sp["m1"].X, Y: sp["m1"].Y + 1}
			delete(sp, "b1")
			delete(sp, "m1")
			sp["t1"] = t1
			sp["m3"] = m3
			np.CurrentX = t1.X
			np.CurrentY = m3.Y
		}
	}
	np.Space = sp
	np.setKey()
	return &np
}

func (p *Piece) Turnleft() *Piece {
	np := Piece{
		Name:     p.Name,
		Score:    &Score{},
		Rotation: p.Rotation - 1,
	}

	if np.Rotation < 0 {
		np.Rotation = 3
	}

	sp := make(map[string]Cell, 4)
	for i, v := range p.Space {
		sp[i] = v
	}

	switch p.Name {
	case I:
		switch p.Rotation {
		case 0:
			t2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y + 1}
			b2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y - 1}
			x2 := Cell{X: sp["m4"].X - 2, Y: sp["m4"].Y - 2}
			delete(sp, "m1")
			delete(sp, "m3")
			delete(sp, "m4")
			sp["t2"] = t2
			sp["b2"] = b2
			sp["x2"] = x2
			np.CurrentX = x2.X
			np.CurrentY = x2.Y
		case 1:
			m1 := Cell{X: sp["t3"].X - 2, Y: sp["t3"].Y - 1}
			m2 := Cell{X: sp["b3"].X - 1, Y: sp["b3"].Y + 1}
			m4 := Cell{X: sp["x3"].X + 1, Y: sp["x3"].Y + 2}
			delete(sp, "t3")
			delete(sp, "b3")
			delete(sp, "x3")
			sp["m1"] = m1
			sp["m2"] = m2
			sp["m4"] = m4
			np.CurrentX = m1.X
			np.CurrentY = m1.Y
		case 2:
			t3 := Cell{X: sp["b4"].X - 1, Y: sp["b4"].Y + 2}
			m3 := Cell{X: sp["b2"].X + 1, Y: sp["b2"].Y + 1}
			x3 := Cell{X: sp["b1"].X + 2, Y: sp["b1"].Y - 1}
			delete(sp, "b4")
			delete(sp, "b2")
			delete(sp, "b1")
			sp["t3"] = t3
			sp["m3"] = m3
			sp["x3"] = x3
			np.CurrentX = x3.X
			np.CurrentY = x3.Y
		case 3:
			b1 := Cell{X: sp["t2"].X - 1, Y: sp["t2"].Y - 2}
			b3 := Cell{X: sp["m2"].X + 1, Y: sp["m2"].Y - 1}
			b4 := Cell{X: sp["x2"].X + 2, Y: sp["x2"].Y + 1}
			delete(sp, "t2")
			delete(sp, "m2")
			delete(sp, "x2")
			sp["b1"] = b1
			sp["b3"] = b3
			sp["b4"] = b4
			np.CurrentX = b1.X
			np.CurrentY = b1.Y
		}
	case J:
		switch p.Rotation {
		case 0:
			b1 := Cell{X: sp["t1"].X, Y: sp["t1"].Y - 2}
			b2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y - 1}
			t2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y + 1}
			delete(sp, "t1")
			delete(sp, "m1")
			delete(sp, "m3")
			sp["b1"] = b1
			sp["b2"] = b2
			sp["t2"] = t2
			np.CurrentX = b1.X
			np.CurrentY = b1.Y
		case 1:
			m1 := Cell{X: sp["t2"].X - 1, Y: sp["t2"].Y - 1}
			t1 := Cell{X: sp["t3"].X - 2, Y: sp["t3"].Y}
			m3 := Cell{X: sp["b2"].X + 1, Y: sp["b2"].Y + 1}
			delete(sp, "t2")
			delete(sp, "t3")
			delete(sp, "b2")
			sp["m1"] = m1
			sp["t1"] = t1
			sp["m3"] = m3
			np.CurrentX = m1.X
			np.CurrentY = m1.Y
		case 2:
			t2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y + 1}
			t3 := Cell{X: sp["b3"].X, Y: sp["b3"].Y + 2}
			b2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y - 1}
			delete(sp, "m3")
			delete(sp, "b3")
			delete(sp, "m1")
			sp["t2"] = t2
			sp["t3"] = t3
			sp["b2"] = b2
			np.CurrentX = b2.X
			np.CurrentY = b2.Y
		case 3:
			m1 := Cell{X: sp["t2"].X - 1, Y: sp["t2"].Y - 1}
			m3 := Cell{X: sp["b2"].X + 1, Y: sp["b2"].Y + 1}
			b3 := Cell{X: sp["b1"].X + 2, Y: sp["b1"].Y}
			delete(sp, "t2")
			delete(sp, "b2")
			delete(sp, "b1")
			sp["m1"] = m1
			sp["m3"] = m3
			sp["b3"] = b3
			np.CurrentX = m1.X
			np.CurrentY = b3.Y
		}
	case L:
		switch p.Rotation {
		case 0:
			t2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y + 1}
			b2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y - 1}
			t1 := Cell{X: sp["t3"].X - 2, Y: sp["t3"].Y}
			delete(sp, "m1")
			delete(sp, "m3")
			delete(sp, "t3")
			sp["t2"] = t2
			sp["b2"] = b2
			sp["t1"] = t1
			np.CurrentX = t1.X
			np.CurrentY = b2.Y
		case 1:
			m1 := Cell{X: sp["t2"].X - 1, Y: sp["t2"].Y - 1}
			m3 := Cell{X: sp["b2"].X + 1, Y: sp["b2"].Y + 1}
			t3 := Cell{X: sp["b3"].X, Y: sp["b3"].Y + 2}
			delete(sp, "b2")
			delete(sp, "t2")
			delete(sp, "b3")
			sp["m1"] = m1
			sp["m3"] = m3
			sp["t3"] = t3
			np.CurrentX = m1.X
			np.CurrentY = m1.Y
		case 2:
			t2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y + 1}
			b2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y - 1}
			b3 := Cell{X: sp["b1"].X + 2, Y: sp["b1"].Y}
			delete(sp, "m1")
			delete(sp, "m3")
			delete(sp, "b1")
			sp["t2"] = t2
			sp["b2"] = b2
			sp["b3"] = b3
			np.CurrentX = b2.X
			np.CurrentY = b2.Y
		case 3:
			m1 := Cell{X: sp["t2"].X - 1, Y: sp["t2"].Y - 1}
			m3 := Cell{X: sp["b2"].X + 1, Y: sp["b2"].Y + 1}
			b1 := Cell{X: sp["t1"].X, Y: sp["t1"].Y - 2}
			delete(sp, "b2")
			delete(sp, "t2")
			delete(sp, "t1")
			sp["m1"] = m1
			sp["m3"] = m3
			sp["b1"] = b1
			np.CurrentX = b1.X
			np.CurrentY = b1.Y
		}
	case T:
		switch p.Rotation {
		case 0:
			b2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y - 1}
			delete(sp, "m3")
			sp["b2"] = b2
			np.CurrentX = p.CurrentX
			np.CurrentY = b2.Y
		case 1:
			m1 := Cell{X: sp["b2"].X - 1, Y: sp["b2"].Y + 1}
			delete(sp, "b2")
			sp["m1"] = m1
			np.CurrentX = m1.X
			np.CurrentY = m1.Y
		case 2:
			t2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y + 1}
			delete(sp, "m1")
			sp["t2"] = t2
			np.CurrentX = t2.X
			np.CurrentY = p.CurrentY
		case 3:
			m3 := Cell{X: sp["t2"].X + 1, Y: sp["t2"].Y - 1}
			delete(sp, "t2")
			sp["m3"] = m3
			np.CurrentX = p.CurrentX
			np.CurrentY = p.CurrentY
		}
	case S:
		switch p.Rotation {
		case 0:
			t1 := Cell{X: sp["t2"].X - 1, Y: sp["t2"].Y}
			b2 := Cell{X: sp["t3"].X - 1, Y: sp["t3"].Y - 2}
			delete(sp, "t2")
			delete(sp, "t3")
			sp["t1"] = t1
			sp["b2"] = b2
			np.CurrentX = t1.X
			np.CurrentY = b2.Y
		case 1:
			t3 := Cell{X: sp["m3"].X, Y: sp["m3"].Y + 1}
			m1 := Cell{X: sp["b3"].X - 2, Y: sp["b3"].Y + 1}
			delete(sp, "m3")
			delete(sp, "b3")
			sp["t3"] = t3
			sp["m1"] = m1
			np.CurrentX = m1.X
			np.CurrentY = m1.Y
		case 2:
			b3 := Cell{X: sp["b2"].X + 1, Y: sp["b2"].Y}
			t2 := Cell{X: sp["b1"].X + 1, Y: sp["b1"].Y + 2}
			delete(sp, "b2")
			delete(sp, "b1")
			sp["b3"] = b3
			sp["t2"] = t2
			np.CurrentX = t2.X
			np.CurrentY = b3.Y
		case 3:
			b1 := Cell{X: sp["m1"].X, Y: sp["m1"].Y - 1}
			m3 := Cell{X: sp["t1"].X + 2, Y: sp["t1"].Y - 1}
			delete(sp, "m1")
			delete(sp, "t1")
			sp["b1"] = b1
			sp["m3"] = m3
			np.CurrentX = b1.X
			np.CurrentY = b1.Y
		}
	case Z:
		switch p.Rotation {
		case 0:
			b1 := Cell{X: sp["m3"].X - 2, Y: sp["m3"].Y - 1}
			m1 := Cell{X: sp["t1"].X, Y: sp["t1"].Y - 1}
			delete(sp, "m3")
			delete(sp, "t1")
			sp["b1"] = b1
			sp["m1"] = m1
			np.CurrentX = b1.X
			np.CurrentY = b1.Y
		case 1:
			t1 := Cell{X: sp["b2"].X - 1, Y: sp["b2"].Y + 2}
			t2 := Cell{X: sp["t3"].X - 1, Y: sp["t3"].Y}
			delete(sp, "b2")
			delete(sp, "t3")
			sp["t1"] = t1
			sp["t2"] = t2
			np.CurrentX = t1.X
			np.CurrentY = sp["m2"].Y
		case 2:
			t3 := Cell{X: sp["m1"].X + 2, Y: sp["m1"].Y + 1}
			m3 := Cell{X: sp["b3"].X, Y: sp["b3"].Y + 1}
			delete(sp, "m1")
			delete(sp, "b3")
			sp["t3"] = t3
			sp["m3"] = m3
			np.CurrentX = sp["b2"].X
			np.CurrentY = sp["b2"].Y
		case 3:
			b3 := Cell{X: sp["t2"].X + 1, Y: sp["t2"].Y - 2}
			b2 := Cell{X: sp["b1"].X + 1, Y: sp["b1"].Y}
			delete(sp, "t2")
			delete(sp, "b1")
			sp["b3"] = b3
			sp["b2"] = b2
			np.CurrentX = sp["m1"].X
			np.CurrentY = b2.Y
		}
	}
	np.Space = sp
	np.setKey()
	return &np
}

func (p *Piece) IsDown(stack *Stack) bool {
	if p.Name == I || p.Name == S || p.Name == Z {
		if stack.Exist(p.Key - 20000) {
			return false
		}
	}
	if stack.Exist(p.Key - 1) {
		return false
	}
	return true
}

func (p *Piece) shorterPath(nMoves string) {
	if len(nMoves) < len(p.Moves) {
		p.Moves = nMoves
	}
}

func (p *Piece) setKey() {
	if p.CurrentX >= 0 && p.CurrentY >= 0 {
		p.Key = p.Rotation*10000 + p.CurrentX*100 + p.CurrentY
	} else {
		p.Key = -1
	}
}

func (p *Piece) setHighY() {
	p.Score.HighY = p.CurrentY
}

func (p *Piece) setStep() {
	if p.Tspin || p.Tspin2 || p.PerfectClear {
		return
	}
	pp := p.FieldAfter.Picks
	maxX, leftY, rightY := 0, 0, 0

	for _, c := range p.Space {
		if c.X > maxX {
			maxX = c.X
		}
	}
	for _, c := range p.Space {
		if c.X == p.CurrentX && c.Y > leftY {
			leftY = c.Y
		}
		if c.X == maxX && c.Y > rightY {
			rightY = c.Y
		}
	}

	if p.CurrentX > 0 {
		x := p.CurrentX - 1
		pick := pp[x] - 1 + p.FieldAfter.Burned
		if leftY > pick {
			p.Score.Step += leftY - pick
		}
	}

	if maxX < p.FieldAfter.Width-1 {
		x := maxX + 1
		pick := pp[x] - 1 + p.FieldAfter.Burned
		if rightY > pick {
			p.Score.Step += rightY - pick
		}
	}
}

func (p *Piece) setCHoles() {
	var effective []*Cell
	lowEffectiveY := 0
	if p.FieldAfter.CountBH > 5 {
		cellsOrderedBy(MAXY).Sort(p.FieldAfter.Holes)
		effective = p.FieldAfter.Holes[:5]
		lowEffectiveY = effective[4].Y
	} else {
		effective = p.FieldAfter.Holes
	}

	var ps [4]int
	i := 0
	for _, v := range p.Space {
		ps[i] = v.X
		i++
	}

	for _, h := range effective {
		if h.X == ps[0] || h.X == ps[1] || h.X == ps[2] || h.X == ps[3] {
			p.Score.CHoles += h.Y - lowEffectiveY
		}
	}
}

func (p *Piece) setTotalScore(st Strategy) {
	p.Score.Burn = p.FieldAfter.Burned
	kS := st.Step
	if p.FieldAfter.Empty < 5 {
		kS = st.Step + 1
	}

	points := p.getPoints()
	p.Score.Total = p.Score.BHoles*st.BHoles +
		p.Score.FHoles*st.FHoles +
		p.Score.HighY*st.HighY +
		p.Score.Step*kS +
		p.Score.NScore +
		p.Score.CHoles*st.CHoles -
		p.Score.Burn*st.Burn -
		points*4

	if p.Score.Burn == 4 {
		p.Score.Total = p.Score.Total - 60
	}

	if p.Tspin && p.FieldAfter.Empty > 4 {
		p.Score.Total = p.Score.Total - 10
	}

	if p.Tspin2 && p.FieldAfter.Empty > 2 {
		p.Score.Total = p.Score.Total - 60
	}

	if p.FieldAfter.Empty == 0 {
		p.Score.Total = p.Score.Total + 100
	}

	if p.FieldAfter.Empty == 1 {
		p.Score.Total = p.Score.Total + 10
	}
}

func (p *Piece) SetScore(st Strategy, oldBH, oldFH, nextScore int) {
	p.Score.BHoles = p.FieldAfter.CountBH - oldBH
	p.Score.FHoles = p.FieldAfter.CountFH - oldFH
	p.Score.NScore = nextScore
	p.Tspin = p.isSingleTSpin()
	p.Tspin2 = p.isDoubleTSpin()
	p.PerfectClear = p.isPerfectClear()
	p.setHighY()
	p.setStep()
	p.setCHoles()
	p.setTotalScore(st)
}

func (p *Piece) getPoints() int {
	points := 0
	if p == nil {
		return points
	}
	switch p.Score.Burn {
	case 2:
		points = 3
	case 3:
		points = 6
	case 4:
		points = 10
	}
	if p.Tspin {
		points = 5
	}
	if p.Tspin2 {
		points = 10
	}
	if p.PerfectClear {
		points = 18
	}
	return points
}

func (p *Piece) shouldSkip(skips int) bool {
	if p.Score.Total > 20 && skips > 0 {
		return true
	}
	return false
}

func (p *Piece) isSingleTSpin() bool {
	if p.Name != T {
		return false
	}
	if p.Rotation != 2 {
		return false
	}
	if p.FieldAfter.Burned != 1 {
		return false
	}
	if p.Space["m1"].Y-1 < 0 {
		return false
	}
	if p.FieldAfter.Grid[p.Space["m1"].Y][p.Space["m1"].X] || p.FieldAfter.Grid[p.Space["m3"].Y][p.Space["m3"].X] {
		return true
	}
	return false
}

func (p *Piece) isDoubleTSpin() bool {
	if p.Name != T {
		return false
	}
	if p.Rotation != 2 {
		return false
	}
	if p.FieldAfter.Burned != 2 {
		return false
	}
	if p.Space["m1"].Y-2 < 0 {
		return false
	}
	if p.FieldAfter.Grid[p.Space["m1"].Y-1][p.Space["m1"].X] || p.FieldAfter.Grid[p.Space["m3"].Y-1][p.Space["m3"].X] {
		return true
	}
	return false
}

func (p *Piece) isPerfectClear() bool {
	for _, row := range p.FieldAfter.Grid {
		for _, col := range row {
			if col {
				return false
			}
		}
	}
	return true
}
