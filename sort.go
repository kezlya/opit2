package main

import (
	"sort"
)

//http://golang.org/pkg/sort/#example_Interface

var SCORE = func(c1, c2 *Piece) bool { return c1.Score.Total < c2.Score.Total }

type lessFunc func(p1, p2 *Piece) bool

type multiSorter struct {
	positions []Piece
	less      []lessFunc
}

func (ms *multiSorter) Sort(positions []Piece) {
	ms.positions = positions
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.positions)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.positions[i], ms.positions[j] = ms.positions[j], ms.positions[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.positions[i], &ms.positions[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}

var MAXY = func(c1, c2 *Cell) bool { return c1.Y > c2.Y }

type cLessFunc func(p1, p2 *Cell) bool

type cMultiSorter struct {
	cells []Cell
	less  []cLessFunc
}

func (ms *cMultiSorter) Sort(cells []Cell) {
	ms.cells = cells
	sort.Sort(ms)
}

func CellOrderedBy(less ...cLessFunc) *cMultiSorter {
	return &cMultiSorter{
		less: less,
	}
}

func (ms *cMultiSorter) Len() int {
	return len(ms.cells)
}

func (ms *cMultiSorter) Swap(i, j int) {
	ms.cells[i], ms.cells[j] = ms.cells[j], ms.cells[i]
}

func (ms *cMultiSorter) Less(i, j int) bool {
	p, q := &ms.cells[i], &ms.cells[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}
