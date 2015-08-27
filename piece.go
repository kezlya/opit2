package main

type Piece struct {
	Name  string
	Space [4]Cell
}

func (p Piece) Left() Piece {
	res := [4]Cell{}
	res[0] = Cell{X: p.Space[0].X - 1, Y: p.Space[0].Y}
	res[1] = Cell{X: p.Space[1].X - 1, Y: p.Space[1].Y}
	res[2] = Cell{X: p.Space[2].X - 1, Y: p.Space[2].Y}
	res[3] = Cell{X: p.Space[3].X - 1, Y: p.Space[3].Y}
	return Piece{Name: p.Name, Space: res}
}

func (p Piece) Right() Piece {
	res := [4]Cell{}
	res[0] = Cell{X: p.Space[0].X + 1, Y: p.Space[0].Y}
	res[1] = Cell{X: p.Space[1].X + 1, Y: p.Space[1].Y}
	res[2] = Cell{X: p.Space[2].X + 1, Y: p.Space[2].Y}
	res[3] = Cell{X: p.Space[3].X + 1, Y: p.Space[3].Y}
	return Piece{Name: p.Name, Space: res}
}

func (p Piece) Down() Piece {
	res := [4]Cell{}
	res[0] = Cell{X: p.Space[0].X, Y: p.Space[0].Y - 1}
	res[1] = Cell{X: p.Space[1].X, Y: p.Space[1].Y - 1}
	res[2] = Cell{X: p.Space[2].X, Y: p.Space[2].Y - 1}
	res[3] = Cell{X: p.Space[3].X, Y: p.Space[3].Y - 1}
	return Piece{Name: p.Name, Space: res}
}
