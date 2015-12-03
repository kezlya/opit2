package main

func (a Grid) isEqual(b Grid) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func (a Picks) isEqual(b Picks) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (g Grid) toField() Field {

	//TODO get rid of this
	picks := make([]int, len(g[0]))
	for i, row := range g {
		for j, col := range row {
			if col && i+1 > picks[j] {
				picks[j] = i + 1
			}
		}
	}

	field := Field{
		Height: len(g),
		Width:  len(g[0]),
		Grid:   g,
		Picks:  picks,
	}
	field.MaxY = field.Picks.Max()
	field.Empty = field.Height - field.MaxY
	return field
}
