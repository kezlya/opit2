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
