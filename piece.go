package main

type Piece struct {
	Name     string
	Rotation int8
	Space    map[string]Cell
}

func (p Piece) Left() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X - 1, Y: v.Y}
	}
	return Piece{Name: p.Name, Rotation: p.Rotation, Space: res}
}

func (p Piece) Right() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X + 1, Y: v.Y}
	}
	return Piece{Name: p.Name, Space: res}
}

func (p Piece) Down() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X, Y: v.Y - 1}
	}
	return Piece{Name: p.Name, Rotation: p.Rotation, Space: res}
}

func (p Piece) Turnright() Piece {
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
		case 2:
		case 3:

		}
	}

	rot := p.Rotation + 1
	if rot > 3 {
		rot = 0
	}
	return Piece{Name: p.Name, Rotation: rot, Space: res}
}
