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
