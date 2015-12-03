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
	}

	field.Picks = make([]int, field.Width)
	for i, row := range g {
		for j, col := range row {
			if col && i+1 > field.Picks[j] {
				field.Picks[j] = i + 1
			}
		}
	}

	field.MaxY = field.Picks.Max()
	field.Empty = field.Height - field.MaxY
	return field
}
