package main

import (
	"sort"
)

//http://golang.org/pkg/sort/#example_Interface

var DAMAGE = func(c1, c2 *Position) bool { return c1.Damage < c2.Damage }
var LOWY = func(c1, c2 *Position) bool { return c1.LowY < c2.LowY }
var HIGHY = func(c1, c2 *Position) bool { return c1.HighY < c2.HighY }
var SCORE = func(c1, c2 *Position) bool { return c1.Score < c2.Score }

type lessFunc func(p1, p2 *Position) bool

type multiSorter struct {
	positions []Position
	less      []lessFunc
}

func (ms *multiSorter) Sort(positions []Position) {
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
