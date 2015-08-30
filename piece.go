package main

type Piece struct {
	Name     string
	Rotation int
	Space    map[string]Cell
	CurrentX int
	CurrentY int
	Key      int
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
	p.CurrentY = start.Y
	p.Key = p.key()
}

func (p *Piece) Left() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X - 1, Y: v.Y}
	}
	np := Piece{Name: p.Name, Rotation: p.Rotation, CurrentX: p.CurrentX - 1, CurrentY: p.CurrentY, Space: res}
	np.Key = np.key()
	return np
}

func (p *Piece) Right() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X + 1, Y: v.Y}
	}
	np := Piece{Name: p.Name, Rotation: p.Rotation, CurrentX: p.CurrentX + 1, CurrentY: p.CurrentY, Space: res}
	np.Key = np.key()
	return np
}

func (p *Piece) Down() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X, Y: v.Y - 1}
	}
	np := Piece{Name: p.Name, Rotation: p.Rotation, CurrentX: p.CurrentX, CurrentY: p.CurrentY - 1, Space: res}
	np.Key = np.key()
	return np
}

func (p *Piece) Turnright() Piece {
	np := Piece{Name: p.Name}

	np.Rotation = p.Rotation + 1
	if np.Rotation > 3 {
		np.Rotation = 0
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
	case "T":
		switch p.Rotation {
		case 0:
			b2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y - 1}
			delete(sp, "m1")
			sp["b2"] = b2
			np.CurrentX = b2.X
		case 1:
			m1 := Cell{X: sp["t2"].X - 1, Y: sp["t2"].Y - 1}
			delete(sp, "t2")
			sp["m1"] = m1
			np.CurrentX = m1.X
		case 2:
			t2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y + 1}
			delete(sp, "m3")
			sp["t2"] = t2
			np.CurrentX = p.CurrentX
		case 3:
			m3 := Cell{X: sp["b2"].X + 1, Y: sp["b2"].Y + 1}
			delete(sp, "b2")
			sp["m3"] = m3
			np.CurrentX = p.CurrentX
		}
	}
	np.Space = sp
	np.Key = np.key()
	return np
}

func (p *Piece) Turnleft() Piece {
	np := Piece{Name: p.Name}

	np.Rotation = p.Rotation - 1
	if np.Rotation < 0 {
		np.Rotation = 3
	}

	sp := make(map[string]Cell, 4)
	for i, v := range p.Space {
		sp[i] = v
	}

	switch p.Name {
	case "I":
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
	case "T":
		switch p.Rotation {
		case 0:
			b2 := Cell{X: sp["m3"].X - 1, Y: sp["m3"].Y - 1}
			delete(sp, "m3")
			sp["b2"] = b2
			np.CurrentX = p.CurrentX
		case 1:
			m1 := Cell{X: sp["b2"].X - 1, Y: sp["b2"].Y + 1}
			delete(sp, "b2")
			sp["m1"] = m1
			np.CurrentX = m1.X
		case 2:
			t2 := Cell{X: sp["m1"].X + 1, Y: sp["m1"].Y + 1}
			delete(sp, "m1")
			sp["t2"] = t2
			np.CurrentX = t2.X
		case 3:
			m3 := Cell{X: sp["t2"].X + 1, Y: sp["t2"].Y - 1}
			delete(sp, "t2")
			sp["m3"] = m3
			np.CurrentX = p.CurrentX
		}
	}
	np.Space = sp
	np.Key = np.key()
	return np
}

func (p *Piece) key() int {
	return p.Rotation*10000 + p.CurrentX*100 + p.CurrentY
}
