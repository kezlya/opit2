package main

import (
	//"fmt"
	"log"
	"strings"
)

type Grid [][]bool

func InitGrid(raw string) Grid {
	if raw == "" {
		log.Fatalln("can't create Grid from empty string")
	}

	width := 0
	rows := strings.Split(raw, ";")
	height := len(rows)
	grid := make([][]bool, height)
	for rowIndex, row := range rows {
		y := height - rowIndex
		cells := strings.Split(row, ",")
		w := len(cells)
		if width > 0 && w != width {
			log.Fatalln("can't create Grid from malformed string")
		}
		width = w
		var colums = make([]bool, width)
		for columIndex, value := range cells {
			if value == "2" {
				colums[columIndex] = true
			}
		}
		grid[y-1] = colums
	}
	return grid
}

func (g Grid) ToField() Field {
	if g == nil || len(g) == 0 || len(g[0]) == 0 {
		log.Fatalln("can't create Field from malformed Grid")
	}

	f := Field{
		Height: len(g),
		Width:  len(g[0]),
		Grid:   g,
		Burned: g.burn(),
	}
	f.Picks = make([]int, f.Width)
	for i, row := range g {
		for j, col := range row {
			if !col {
				continue
			}
			y := i + 1
			if y > f.Picks[j] {
				f.Picks[j] = y
			}
			if y > f.MaxPick {
				f.MaxPick = y
			}
		}
	}
	f.Empty = f.Height - f.MaxPick
	f.Holes, f.CountFH = g.findHoles(f.Picks)
	f.CountBH = len(f.Holes)
	return f
}

func (g Grid) Copy() Grid {
	if g == nil || len(g) == 0 || len(g[0]) == 0 {
		log.Fatalln("can't create Field from malformed Grid")
	}
	newGrid := make([][]bool, len(g))
	for i, row := range g {
		newGrid[i] = make([]bool, len(g[0]))
		copy(newGrid[i], row[:])
	}
	return newGrid
}

func (g Grid) ApplyPiece(cells map[string]Cell) Grid {
	newGrid := g.Copy()
	for _, cell := range cells {
		newGrid[cell.Y][cell.X] = true
	}
	return newGrid
}

func (g Grid) IsCollision(cells map[string]Cell, checkTop bool) bool {
	h := len(g)
	if h <= 0 {
		return true
	}
	w := len(g[0])
	if w <= 0 {
		return true
	}
	for _, c := range cells {
		if c.X < 0 || c.X >= w || c.Y < 0 {
			return true
		}
		if c.Y >= h {
			if checkTop {
				return true
			}
			continue
		}
		if g[c.Y][c.X] {
			return true
		}
	}
	return false
}

func (g Grid) findHoles(picks []int) ([]*Cell, int) {
	w := len(g[0])
	blocked := make([]*Cell, 0)
	fixable := 0
	for i, pick := range picks {
		for j := 0; j < pick; j++ {
			if !g[j][i] {
				if (i-2 > -1 && !g[j][i-1] && !g[j][i-2] && !g[j+1][i-1] && !g[j+1][i-2]) ||
					(i+2 < w && !g[j][i+1] && !g[j][i+2] && !g[j+1][i+1] && !g[j+1][i+2]) {
					fixable++
				} else {
					cell := Cell{X: i, Y: j}
					blocked = append(blocked, &cell)
				}
			}
		}
	}
	return blocked, fixable
}

//TODO need to depricate eventially move to the loop of ToField method
func (g Grid) burn() int {
	burned := 0
	for i := 0; i < len(g); i++ {
		check := true
		for _, col := range g[i] {
			if !col {
				check = false
				break
			}
		}
		if check {
			g = append(g[:i], g[i+1:]...)
			burned++
			i--
		}
	}
	return burned
}

func (g Grid) tSpinLevels(maxH int) (int, int, int, int) {
	var tlevel1, tlevel2, tlevel3, tlevel4 int
	for row := 0; row < maxH; row++ {
		check, cell := g.isHoleInTheRow(row, maxH)
		if check {
			if g.tSpinLevel2(&cell) {
				if g.tSpinLevel3(&cell) {
					if g.tSpinLevel4(&cell) {
						tlevel4++
					} else {
						tlevel3++
					}
				} else {
					tlevel2++
				}
			} else {
				tlevel1++
			}
		}
	}
	return tlevel1, tlevel2, tlevel3, tlevel4
}

func (g Grid) isHoleInTheRow(row, maxH int) (bool, Cell) {
	var isHole bool
	var cell Cell
	for i, c := range g[row] {
		if !c {
			if isHole {
				isHole = false
				break
			} else {
				isHole = true
				cell = Cell{X: i, Y: row}
			}
		}
	}
	if cell.X < 2 || cell.X > len(g[row])-3 {
		isHole = false
	}
	/*for col := 1; col < maxW; col++ {
		if !g[row][col] {
			if isHole {
				isHole = false
				break
			} else {
				isHole = true
				cell = Cell{X: col, Y: row}
			}
		}
	}*/
	if isHole {
		for y := cell.Y; y < maxH; y++ {
			if g[y][cell.X] {
				isHole = false
				break
			}
		}
	}
	return isHole, cell
}

func (g Grid) tSpinLevel2(c *Cell) bool {
	if c.Y+1 < len(g)-1 &&
		!g[c.Y+1][c.X] &&
		!g[c.Y+1][c.X-1] &&
		!g[c.Y+1][c.X+1] {
		return true
	}
	return false
}

func (g Grid) tSpinLevel3(c *Cell) bool {
	check := true
	for i, col := range g[c.Y+1] {
		if !col && (i < c.X-1 || i > c.X+1) {
			check = false
		}
	}
	return check
}

func (g Grid) tSpinLevel4(c *Cell) bool {
	var check bool
	if (!g[c.Y+2][c.X-1] && !g[c.Y+2][c.X] && g[c.Y+2][c.X+1]) ||
		(g[c.Y+2][c.X-1] && !g[c.Y+2][c.X] && !g[c.Y+2][c.X+1]) {
		check = true
	}
	return check
}
