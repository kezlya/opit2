package main

import (
	"log"
	"strings"
)

type Grid [][]bool

func GridFromString(raw string) Grid {
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
	f.HolesBlocked, f.HolesFixable = g.findHoles(f.Picks)
	f.CountBH = len(f.HolesBlocked)
	f.CountFH = len(f.HolesFixable)
	return f
}

//TODO return only count of holes
func (g Grid) findHoles(picks Picks) ([]Cell, []Cell) {
	w := len(g[0])
	blocked := make([]Cell, 0)
	fixable := make([]Cell, 0)
	for i, pick := range picks {
		for j := 0; j < pick; j++ {
			if !g[j][i] {
				hole := Cell{X: i, Y: j}
				if (i-2 > -1 && !g[j][i-1] && !g[j][i-2] && !g[j+1][i-1] && !g[j+1][i-2]) ||
					(i+2 < w && !g[j][i+1] && !g[j][i+2] && !g[j+1][i+1] && !g[j+1][i+2]) {
					fixable = append(fixable, hole)
				} else {
					blocked = append(blocked, hole)
				}
			}
		}
	}
	return blocked, fixable
}
