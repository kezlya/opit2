package main

import (
	"fmt"
	//"strings"
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

	//bag := make(map[int]*Piece)
	//queue := make(map[int]bool)
	//newKey := 0

	//bag[piece.Key] = &piece
	//queue[piece.Key] = true

	/*buff := 5
	if piece.Name == "O" {
		buff = 3
	}*/

	stack := new(Stack)
	stack.Init()

	stack.Push(&piece)
	//fmt.Println("stack collection", stack.collection)
	for stack.Len() > 0 {
		k := stack.Pop()
		f.Search(stack, k, "down")
		f.Search(stack, k, "left")
		f.Search(stack, k, "right")
		if piece.Name != "O" {
			f.Search(stack, k, "turnleft")
			f.Search(stack, k, "turnright")
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
	fmt.Println("stack collection", len(stack.collection))

	bag := stack.collection
	for k, p := range bag {
		_, ok := bag[k-1]
		if !ok && !f.Grid.IsCollision(p.Space, true) {
			newGrid := f.Grid.ApplyPiece(p.Space)
			newField := newGrid.ToField()
			p.FieldAfter = &newField
			//p.Moves = strings.TrimPrefix(p.Moves, ",")
			positions = append(positions, *p)
		}
	}
	return positions
}

func (f Field) Search(stack *Stack, key int, dir string) {
	//fmt.Println(key)
	//var ok bool
	//var el *Piece

	//fmt.Println(bag[key])
	//nMoves := stack.collection[key].Moves + "," + dir
	var current = stack.Peek(key)
	if current == nil {
		return
	}
	var np Piece
	switch dir {
	case "left":
		nextKey := key - 100
		if nextKey%10000/100 < 0 {
			return
		}
		if stack.Exist(nextKey) {
			//el.shorterPath(nMoves)
			return
		}
		np = current.Left()
	case "right":
		nextKey := key + 100
		if nextKey%10000/100 > f.Width {
			return
		}
		if stack.Exist(nextKey) {
			//el.shorterPath(nMoves)
			return
		}
		np = current.Right()
	case "down":
		nextKey := key - 1
		if nextKey%100 < 0 {
			return
		}
		if stack.Exist(nextKey) {
			//el.shorterPath(nMoves)
			return
		}
		np = current.Down()
	case "turnleft":
		np = current.Turnleft()
		if stack.Exist(np.Key) {
			//el.shorterPath(nMoves)
			return
		}
	case "turnright":
		np = current.Turnright()
		if stack.Exist(np.Key) {
			//el.shorterPath(nMoves)
			return
		}
	}

	if np.Name == "I" || np.Name == "S" || np.Name == "Z" {
		if stack.Exist(np.Key-20000) || stack.Exist(np.Key+20000) {
			return
		}
	}

	if f.Grid.IsCollision(np.Space, false) {
		return
	}

	//np.Moves = nMoves
	//bag[np.Key] = &np
	//fmt.Println(np)
	stack.Push(&np)
}
