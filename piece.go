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
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = v
	}

	switch p.Name {
	case "T":
		switch p.Rotation {
		case 0:
			nX := res["m1"].X + 1
			nY := res["m1"].Y - 1
			delete(res, "m1")
			res["b2"] = Cell{X: nX, Y: nY}
		case 1:
			nX := res["t2"].X - 1
			nY := res["t2"].Y - 1
			delete(res, "t2")
			res["m1"] = Cell{X: nX, Y: nY}
		case 2:
			nX := res["m3"].X - 1
			nY := res["m3"].Y + 1
			delete(res, "m3")
			res["t2"] = Cell{X: nX, Y: nY}
		case 3:
			nX := res["b2"].X + 1
			nY := res["b2"].Y + 1
			delete(res, "b2")
			res["m3"] = Cell{X: nX, Y: nY}
		}
	}

	rot := p.Rotation + 1
	if rot > 3 {
		rot = 0
	}
	return Piece{Name: p.Name, Rotation: rot, Space: res}
}

func (p *Piece) Turnleft() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = v
	}

	switch p.Name {
	case "T":
		switch p.Rotation {
		case 0:
			nX := res["m3"].X - 1
			nY := res["m3"].Y - 1
			delete(res, "m3")
			res["b2"] = Cell{X: nX, Y: nY}
		case 1:
			nX := res["b2"].X - 1
			nY := res["b2"].Y + 1
			delete(res, "b2")
			res["m1"] = Cell{X: nX, Y: nY}
		case 2:
			nX := res["m1"].X + 1
			nY := res["m1"].Y + 1
			delete(res, "m1")
			res["t2"] = Cell{X: nX, Y: nY}
		case 3:
			nX := res["t2"].X + 1
			nY := res["t2"].Y - 1
			delete(res, "t2")
			res["m3"] = Cell{X: nX, Y: nY}
		}
	}

	rot := p.Rotation - 1
	if rot < 0 {
		rot = 3
	}
	return Piece{Name: p.Name, Rotation: rot, Space: res}
}
