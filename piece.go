package main

//import "fmt"

type Piece struct {
	Key        int
	Name       string
	CurrentX   int
	CurrentY   int
	Rotation   int
	Space      map[string]Cell
	FieldAfter Field
	Moves      string
	Score      Score
	IsHole     bool
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
	p.setKey()
}

func (p *Piece) Left() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X - 1, Y: v.Y}
	}
	np := Piece{Name: p.Name, Rotation: p.Rotation, CurrentX: p.CurrentX - 1, CurrentY: p.CurrentY, Space: res}
	np.setKey()
	return np
}

func (p *Piece) Right() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X + 1, Y: v.Y}
	}
	np := Piece{Name: p.Name, Rotation: p.Rotation, CurrentX: p.CurrentX + 1, CurrentY: p.CurrentY, Space: res}
	np.setKey()
	return np
}

func (p *Piece) Down() Piece {
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X, Y: v.Y - 1}
	}
	np := Piece{Name: p.Name, Rotation: p.Rotation, CurrentX: p.CurrentX, CurrentY: p.CurrentY - 1, Space: res}
	np.setKey()
	return np
}

func (p *Piece) DropTo(y int) Piece {
	drop := p.CurrentY - y
	res := make(map[string]Cell, 4)
	for i, v := range p.Space {
		res[i] = Cell{X: v.X, Y: v.Y - drop}
	}
	np := Piece{Name: p.Name, Rotation: p.Rotation, CurrentX: p.CurrentX, CurrentY: p.CurrentY - drop, Space: res}
	np.setKey()
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
	case "J":
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
	case "L":
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
	case "T":
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
	case "S":
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
	case "Z":
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
	case "J":
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
	case "L":
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
	case "T":
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
	case "S":
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
	case "Z":
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
	return np
}

func (p *Piece) setKey() {
	if p.CurrentX >= 0 && p.CurrentY >= 0 {
		p.Key = p.Rotation*10000 + p.CurrentX*100 + p.CurrentY
	}
}

func (p *Piece) setBurn() {
	p.Score.Burn = p.FieldAfter.WillBurn()
}

func (p *Piece) setHighY() {
	switch p.Name {
	case "I":
		switch p.Rotation {
		case 0, 2:
			p.Score.HighY = p.CurrentY
		case 1, 3:
			p.Score.HighY = p.CurrentY + 3
		}
	case "J", "L", "S", "T", "Z":
		switch p.Rotation {
		case 0, 2:
			p.Score.HighY = p.CurrentY + 1
		case 1, 3:
			p.Score.HighY = p.CurrentY + 2
		}
	case "O":
		p.Score.HighY = p.CurrentY + 1
	}
}

func (p *Piece) setStep(pp Picks) {
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
		pick := pp[x] - 1
		if leftY > pick {
			p.Score.Step += leftY - pick
		}
	}

	if maxX < p.FieldAfter.Width()-1 {
		x := maxX + 1
		pick := pp[x] - 1
		if rightY > pick {
			p.Score.Step += rightY - pick
		}
	}
}

func (p *Piece) setCHoles(hBlocked []Cell) {
	var effective []Cell
	lowEffectiveY := 0
	if len(hBlocked) > 5 {
		CellOrderedBy(MAXY).Sort(hBlocked)
		effective = hBlocked[:5]
		lowEffectiveY = effective[4].Y
	} else {
		effective = hBlocked
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
	/*
		for _, h := range effective {
			stucking := 0
			if h.X == ps[0] {
				stucking++
			}
			if h.X == ps[1] {
				stucking++
			}
			if h.X == ps[2] {
				stucking++

			}
			if h.X == ps[3] {
				stucking++
			}

			if stucking > 0 {
				deep := h.Y - lowEffectiveY + 1

				p.Score.CHoles += deep + stucking
			}
		}
	*/
}

func (p *Piece) setTotalScore(st Strategy, empty int) {

	kY := st.HighY
	kS := st.Step
	kC := st.CHoles
	kB := st.Burn

	if empty < 5 {
		kS = kS + 3
	}

	if empty < 5 {
		kB = kB + 3
	}

	points := p.getPoints()
	p.Score.Total = p.Score.BHoles*st.BHoles +
		p.Score.FHoles*st.FHoles +
		p.Score.HighY*kY +
		p.Score.Step*kS +
		p.Score.NScore +
		p.Score.CHoles*kC -
		p.Score.Burn*kB -
		points*4

	if p.isSingleTSpin() && empty > 4 {
		p.Score.Total = p.Score.Total - 10
	}

	if p.isDoubleTSpin() && empty > 4 {
		p.Score.Total = p.Score.Total - 20
	}
	/*
		if combo > 3 && p.Score.Burn > 0 {
			p.Score.Total = p.Score.Total - 10
		}

		delta := p.FieldAfter.Picks().SumStep()
		if delta > 17 {
			p.Score.Total = p.Score.Total + 7
		}
	*/

	if empty == 0 {
		p.Score.Total = p.Score.Total + 100
		//fmt.Println("YESYEYEWYSYEYEYSYEYSYYEYSYEYEYYSYSYEYSYSYYEYSYSYEYEY")
	}
}

func (p *Piece) getPoints() int {
	burn := 0
	switch p.Score.Burn {
	case 2:
		burn = 3
	case 3:
		burn = 6
	case 4:
		burn = 10
	}

	if p.isSingleTSpin() {
		burn = 5
	}

	if p.isDoubleTSpin() {
		burn = 10
	}

	if p.isPerfectClear() {
		burn = 18
	}
	return burn
}

func (p *Piece) isSingleTSpin() bool {
	if p.Name != "T" {
		return false
	}
	if !p.IsHole {
		return false
	}
	if p.Rotation != 2 {
		return false
	}
	if p.Score.Burn != 1 {
		return false
	}
	if p.Space["m1"].Y-1 < 0 {
		return false
	}
	if p.FieldAfter[p.Space["m1"].Y][p.Space["m1"].X] || p.FieldAfter[p.Space["m3"].Y][p.Space["m3"].X] {
		return true
	}
	return false
}

func (p *Piece) isDoubleTSpin() bool {
	if p.Name != "T" {
		return false
	}
	if !p.IsHole {
		return false
	}
	if p.Rotation != 2 {
		return false
	}
	if p.Score.Burn != 2 {
		return false
	}
	if p.Space["m1"].Y-2 < 0 {
		return false
	}
	if p.FieldAfter[p.Space["m1"].Y-1][p.Space["m1"].X] || p.FieldAfter[p.Space["m3"].Y-1][p.Space["m3"].X] {
		return true
	}
	return false
}

func (p *Piece) isPerfectClear() bool {
	for _, row := range p.FieldAfter {
		for _, col := range row {
			if col {
				return false
			}
		}
	}
	return true
}
