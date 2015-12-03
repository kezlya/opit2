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

	field := Field{
		Height: len(g),
		Width:  len(g[0]),
		Grid:   g,
		Picks:  g.picks(),
	}
	field.MaxY = field.Picks.Max()
	field.Empty = field.Height - field.MaxY
	return field
}

func (g Grid) picks() Picks {
	if g == nil || len(g) == 0 || len(g[0]) == 0 {
		log.Fatalln("can't create Picks from malformed Grid")
	}

	picks := make([]int, len(g[0]))
	for i, row := range g {
		for j, col := range row {
			if col && i+1 > picks[j] {
				picks[j] = i + 1
			}
		}
	}
	return picks
}

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
