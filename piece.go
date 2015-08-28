package main

type Piece struct {
	Name     string
	Rotation int
	Space    map[string]Cell
	CurrentX int
}

func (p *Piece) InitSpace(start Cell) {
	space := make(map[string]Cell, 4)
	switch p.Name {
	case "I":
		space["m1"] = Cell{X: start.X, Y: start.Y}
		space["m2"] = Cell{X: start.X + 1, Y: start.Y}
		space["m3"] = Cell{X: start.X + 2, Y: start.Y}
		space["m4"] = Cell{X: start.X + 3, Y: start.Y}
	case "J":
		space["m1"] = Cell{X: start.X, Y: start.Y}
		space["m2"] = Cell{X: start.X + 1, Y: start.Y}
		space["m3"] = Cell{X: start.X + 2, Y: start.Y}
		space["t1"] = Cell{X: start.X, Y: start.Y + 1}
	case "L":
		space["m1"] = Cell{X: start.X, Y: start.Y}
		space["m2"] = Cell{X: start.X + 1, Y: start.Y}
		space["m3"] = Cell{X: start.X + 2, Y: start.Y}
		space["t3"] = Cell{X: start.X + 2, Y: start.Y + 1}
	case "O":
		space["t1"] = Cell{X: start.X, Y: start.Y + 1}
		space["t2"] = Cell{X: start.X + 1, Y: start.Y + 1}
		space["m1"] = Cell{X: start.X, Y: start.Y}
		space["m2"] = Cell{X: start.X + 1, Y: start.Y}
	case "S":
		space["t2"] = Cell{X: start.X + 1, Y: start.Y + 1}
		space["t3"] = Cell{X: start.X + 2, Y: start.Y + 1}
		space["m1"] = Cell{X: start.X, Y: start.Y}
		space["m2"] = Cell{X: start.X + 1, Y: start.Y}
	case "T":
		space["m1"] = Cell{X: start.X, Y: start.Y}
		space["m2"] = Cell{X: start.X + 1, Y: start.Y}
		space["m3"] = Cell{X: start.X + 2, Y: start.Y}
		space["t2"] = Cell{X: start.X + 1, Y: start.Y + 1}
	case "Z":
		space["t1"] = Cell{X: start.X, Y: start.Y + 1}
		space["t2"] = Cell{X: start.X + 1, Y: start.Y + 1}
		space["m2"] = Cell{X: start.X + 1, Y: start.Y}
		space["m3"] = Cell{X: start.X + 2, Y: start.Y}
	}
	p.Space = space
	p.CurrentX = start.X
}

func (p *Piece) Left() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X - 1, Y: v.Y}
	}
	return Piece{Name: p.Name, Rotation: p.Rotation, CurrentX: p.CurrentX - 1, Space: res}
}

func (p *Piece) Right() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X + 1, Y: v.Y}
	}
	return Piece{Name: p.Name, Rotation: p.Rotation, CurrentX: p.CurrentX + 1, Space: res}
}

func (p *Piece) Down() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X, Y: v.Y - 1}
	}
	return Piece{Name: p.Name, Rotation: p.Rotation, CurrentX: p.CurrentX, Space: res}
}

func (p *Piece) Turnright() Piece {
	piece := Piece{Name: p.Name}

	piece.Rotation = p.Rotation + 1
	if piece.Rotation > 3 {
		piece.Rotation = 0
	}

	sp := make(map[string]Cell, 4)
	for i, v := range p.Space {
		sp[i] = v
	}
	switch p.Name {
	case "I":
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
			piece.CurrentX = t3.X
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
			piece.CurrentX = b1.X
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
			piece.CurrentX = t2.X
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
			piece.CurrentX = m1.X
		}
	case "T":
		switch p.Rotation {
		case 0:
			nX := sp["m1"].X + 1
			nY := sp["m1"].Y - 1
			delete(sp, "m1")
			sp["b2"] = Cell{X: nX, Y: nY}
			piece.CurrentX = nX
		case 1:
			nX := sp["t2"].X - 1
			nY := sp["t2"].Y - 1
			delete(sp, "t2")
			sp["m1"] = Cell{X: nX, Y: nY}
			piece.CurrentX = nX
		case 2:
			nX := sp["m3"].X - 1
			nY := sp["m3"].Y + 1
			delete(sp, "m3")
			sp["t2"] = Cell{X: nX, Y: nY}
			piece.CurrentX = p.CurrentX
		case 3:
			nX := sp["b2"].X + 1
			nY := sp["b2"].Y + 1
			delete(sp, "b2")
			sp["m3"] = Cell{X: nX, Y: nY}
			piece.CurrentX = p.CurrentX
		}
	}
	piece.Space = sp
	return piece
}

func (p *Piece) Turnleft() Piece {
	piece := Piece{Name: p.Name}

	piece.Rotation = p.Rotation - 1
	if piece.Rotation < 0 {
		piece.Rotation = 3
	}

	sp := make(map[string]Cell, 4)
	for i, v := range p.Space {
		sp[i] = v
	}

	switch p.Name {
	case "T":
		switch p.Rotation {
		case 0:
			nX := sp["m3"].X - 1
			nY := sp["m3"].Y - 1
			delete(sp, "m3")
			sp["b2"] = Cell{X: nX, Y: nY}
			piece.CurrentX = p.CurrentX
		case 1:
			nX := sp["b2"].X - 1
			nY := sp["b2"].Y + 1
			delete(sp, "b2")
			sp["m1"] = Cell{X: nX, Y: nY}
			piece.CurrentX = nX
		case 2:
			nX := sp["m1"].X + 1
			nY := sp["m1"].Y + 1
			delete(sp, "m1")
			sp["t2"] = Cell{X: nX, Y: nY}
			piece.CurrentX = nX
		case 3:
			nX := sp["t2"].X + 1
			nY := sp["t2"].Y - 1
			delete(sp, "t2")
			sp["m3"] = Cell{X: nX, Y: nY}
			piece.CurrentX = p.CurrentX
		}
	}
	piece.Space = sp
	return piece
}
