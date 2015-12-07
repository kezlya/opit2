package main

import (
	"fmt"
	"strings"
)

type Field struct {
	Width   int
	Height  int
	Empty   int
	MaxPick int
	CountBH int
	CountFH int
	Burned  int

	Grid  Grid
	Picks []int
}

func (f Field) FindPositions(piece Piece) []Piece {
	positions := make([]Piece, 0)
	drop := piece.CurrentY - f.MaxPick - 2
	if drop > 0 {
		piece = piece.Drop(drop)
	}

	bag := make(map[int]*Piece)
	//queue := make(map[int]bool)
	//newKey := 0

	bag[piece.Key] = &piece
	//queue[piece.Key] = true

	/*buff := 5
	if piece.Name == "O" {
		buff = 3
	}*/

	ch := make(chan int)
	ch <- piece.Key
	for k := range ch {
		go f.Search(ch, bag, k, "down")
		go f.Search(ch, bag, k, "left")
		go f.Search(ch, bag, k, "right")
		if piece.Name != "O" {
			go f.Search(ch, bag, k, "turnleft")
			go f.Search(ch, bag, k, "turnright")
		}
	}
	/*
		for len(queue) > 0 {
			tmp := make(map[int]bool)
			for k, _ := range queue {
				go f.Search(ch, "down", k, bag)
				go f.Search(ch, "left", k, bag)
				go f.Search(ch, "right", k, bag)
				if piece.Name != "O" {
					go f.Search(ch, "turnleft", k, bag)
					go f.Search(ch, "turnright", k, bag)
				}
				for i := 0; i < buff; i++ {
					newKey = <-ch
					if newKey >= 0 {
						tmp[newKey] = false
					}
				}
			}
			fmt.Println("=========", len(tmp), "=========")
			queue = tmp
		}*/
	//fmt.Println("countSearchCalls", countSearchCalls)
	fmt.Println("bagLen", len(bag))
	fmt.Println("bagLen", len(bag))

	for k, p := range bag {
		_, ok := bag[k-1]
		if !ok && !f.Grid.IsCollision(p.Space, true) {
			newGrid := f.Grid.ApplyPiece(p.Space)
			newField := newGrid.ToField()
			p.FieldAfter = &newField
			p.Moves = strings.TrimPrefix(p.Moves, ",")
			positions = append(positions, *p)
		}
	}
	return positions
}

func (f Field) Search(keys chan int, bag map[int]*Piece, key int, dir string) {
	var ok bool
	var el *Piece
	var np Piece
	nMoves := bag[key].Moves + "," + dir

	switch dir {
	case "left":
		nextKey := key - 100
		if nextKey%10000/100 < 0 {
			keys <- -1
		}
		el, ok = bag[nextKey]
		if ok {
			el.shorterPath(nMoves)
			keys <- -1
		}
		np = bag[key].Left()
	case "right":
		nextKey := key + 100
		if nextKey%10000/100 > f.Width {
			keys <- -1
		}
		el, ok = bag[nextKey]
		if ok {
			el.shorterPath(nMoves)
			keys <- -1
		}
		np = bag[key].Right()
	case "down":
		nextKey := key - 1
		if nextKey%100 < 0 {
			keys <- -1
		}
		el, ok = bag[nextKey]
		if ok {
			el.shorterPath(nMoves)
			keys <- -1
		}
		np = bag[key].Down()
	case "turnleft":
		np = bag[key].Turnleft()
		el, ok = bag[np.Key]
		if ok {
			el.shorterPath(nMoves)
			keys <- -1
		}
	case "turnright":
		np = bag[key].Turnright()
		el, ok = bag[np.Key]
		if ok {
			el.shorterPath(nMoves)
			keys <- -1
		}
	}

	if np.Name == "I" || np.Name == "S" || np.Name == "Z" {
		_, ok1 := bag[np.Key-20000]
		_, ok2 := bag[np.Key+20000]
		if ok1 || ok2 {
			keys <- -1
		}
	}

	if f.Grid.IsCollision(np.Space, false) {
		keys <- -1
	}

	np.Moves = nMoves
	bag[np.Key] = &np
	keys <- np.Key
}
