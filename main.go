package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Timebank, TimePerMove, Width, Height, Round, CurrentPieceX, CurrentPieceY int
var Players [2]Player
var MyPlayer *Player
var CurrentPiece, NextPiece string

type Player struct {
	Name    string
	Columns []int
	Points  int
	Combo   int
}

type Position struct {
	Rotation int
	X        int
}

func main() {
	consolereader := bufio.NewReader(os.Stdin)
	for {
		input, _ := consolereader.ReadString('\n')
		parts := strings.Split(strings.TrimSpace(input), " ")
		switch parts[0] {
		case "settings":
			_asignSettings(parts[1], parts[2])
		case "update":
			_asignUpdates(parts[1], parts[2], parts[3])
		case "action":
			time, _ := strconv.Atoi(parts[2])
			GetMoves(time)
		}
	}
}

func _asignSettings(action, value string) {
	switch action {
	case "timebank":
		Timebank, _ = strconv.Atoi(value)
	case "time_per_move":
		TimePerMove, _ = strconv.Atoi(value)
	case "player_names":
		names := strings.Split(value, ",")
		for i, name := range names {
			Players[i].Name = name
		}
	case "your_bot":
		for i, p := range Players {
			if p.Name == value {
				MyPlayer = &Players[i]
				break
			}
		}
	case "field_width":
		Width, _ = strconv.Atoi(value)
	case "field_height":
		Height, _ = strconv.Atoi(value)
	}
}

func _asignUpdates(who, action, value string) {
	switch who {
	case "game":
		switch action {
		case "round":
			Round, _ = strconv.Atoi(value)
		case "this_piece_type":
			CurrentPiece = value
		case "next_piece_type":
			NextPiece = value
		case "this_piece_position":
			cor := strings.Split(value, ",")
			CurrentPieceX, _ = strconv.Atoi(cor[0])
			CurrentPieceY, _ = strconv.Atoi(cor[1])
		}
	default:
		switch action {
		case "row_points":
			for i, p := range Players {
				if p.Name == who {
					Players[i].Points, _ = strconv.Atoi(value)
					break
				}
			}
		case "combo":
			for i, p := range Players {
				if p.Name == who {
					Players[i].Combo, _ = strconv.Atoi(value)
					break
				}
			}
		case "field":
			for i, p := range Players {
				if p.Name == who {
					piks := make([]int, Width, Width)
					rows := strings.Split(value, ";")
					for ii, row := range rows {
						y := Height - ii
						columns := strings.Split(row, ",")
						for iii, cell := range columns {
							if (cell == "2" || cell == "3") && y > piks[iii] {
								piks[iii] = y
							}
						}
					}
					Players[i].Columns = piks
					break
				}
			}
		}
	}
}

func GetMoves(time int) {
	if Round == 1 {
		fmt.Println("drop")
		return
	}

	var positions []Position
	switch CurrentPiece {
	case "I":
		positions = _fitsI()
	case "J":
		positions = _fitsJ()
	case "L":
		positions = _fitsL()
	case "O":
		positions = _fitsO()
	case "S":
		positions = _fitsS()
	case "T":
		positions = _fitsT()
	case "Z":
		positions = _fitsZ()
	}

	if len(positions) == 0 {

		finalPosition := _chooseMinimumDamage()
		_printMoves(finalPosition)
	} else {
		finalPositionIndex := _chooseLowestPosition(positions)
		_printMoves(positions[finalPositionIndex])
	}

}

func _fitsI() []Position {
	var pos []Position
	c := MyPlayer.Columns

	for i, v := range MyPlayer.Columns {
		if (_isRight(i, 1) && v+1 < c[i+1]) || (_isLeft(i, 1) && v+1 < c[i-1]) {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 3) && v == c[i+1] && v == c[i+2] && v == c[i+3] {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if (_isRight(i, 1) && v < c[i+1]) || (_isLeft(i, 1) && v < c[i-1]) {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}
	return pos
}

func _fitsJ() []Position {
	var pos []Position
	c := MyPlayer.Columns
	//TODO: replace MyPlayer.Columns to c
	for i, v := range MyPlayer.Columns {
		if _isRight(i, 1) && v+2 == c[i+1] {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 2) && v == c[i+1] && v == c[i+2]+1 {
			p := Position{Rotation: 2, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 2) && v == c[i+1] && v == c[i+2] {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 1) && v == c[i+1] {
			p := Position{Rotation: 3, X: i}
			pos = append(pos, p)
		}
	}

	return pos
}

func _fitsL() []Position {
	var pos []Position
	c := MyPlayer.Columns

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 1) && v == c[i+1]+2 {
			p := Position{Rotation: 3, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 2) && v+1 == c[i+1] && v+1 == c[i+2] {
			p := Position{Rotation: 2, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 2) && v == c[i+1] && v == c[i+2] {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 1) && v == c[i+1] {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}

	return pos
}

func _fitsO() []Position {
	var pos []Position
	c := MyPlayer.Columns

	for i, v := range MyPlayer.Columns {
		if (_isRight(i, 2) && v == c[i+1] && v < c[i+2]) || (_isLeft(i, 1) && v < c[i-1] && _isRight(i, 1) && v == c[i+1]) {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 1) && v == c[i+1] {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	return pos
}

func _fitsS() []Position {
	var pos []Position
	c := MyPlayer.Columns

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 2) && v == c[i+1] && v+1 == c[i+2] {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 1) && v == c[i+1]+1 {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}

	return pos
}

func _fitsT() []Position {
	var pos []Position
	c := MyPlayer.Columns

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 2) && v == c[i+1]+1 && v == c[i+2] {
			p := Position{Rotation: 2, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 2) && v == c[i+1] && v == c[i+2] {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 1) && v+1 == c[i+1] {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 1) && v == c[i+1]+1 {
			p := Position{Rotation: 3, X: i}
			pos = append(pos, p)
		}
	}

	return pos
}

func _fitsZ() []Position {
	var pos []Position
	c := MyPlayer.Columns

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 2) && v == c[i+1]+1 && v == c[i+2]+1 {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range MyPlayer.Columns {
		if _isRight(i, 1) && v+1 == c[i+1] {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}

	return pos
}

func _isRight(i, right int) bool {
	if i+right < Width {
		return true
	} else {
		return false
	}
}

func _isLeft(i, left int) bool {
	if i-left > 0 {
		return true
	} else {
		return false
	}
}

func _chooseLowestPosition(positions []Position) int {
	if len(positions) > 1 {
		cashX := positions[0].X
		indexSmallestPosition := 0
		for i := 1; i < len(positions); i++ {
			if MyPlayer.Columns[cashX] > MyPlayer.Columns[positions[i].X] {
				cashX = positions[i].X
				indexSmallestPosition = i
			}
		}
		return indexSmallestPosition
	}
	return 0
}

func _chooseMinimumDamage() Position {
	p := Position{}
	minDammage := 1000
	rotationMax := 1
	switch CurrentPiece {

	case "I", "Z", "S":
		rotationMax = 2
	case "J", "L", "T":
		rotationMax = 4
	}
	columnsSum := _sum(MyPlayer.Columns)
	for r := 0; r < rotationMax; r++ {
		for i := 0; i < Width; i++ {
			columsAfter := _getColumnsAfter(i, r, CurrentPiece)
			columsAfterSum := _sum(columsAfter)
			if columsAfterSum <= columnsSum {
				break
			}

			damage := columsAfterSum - columnsSum
			if damage < minDammage {
				minDammage = damage
				p.Rotation = r
				p.X = i
			}

			if damage == minDammage && MyPlayer.Columns[p.X] > MyPlayer.Columns[i] {
				p.X = i
			}

		}
	}

	return p
}

func _sum(c []int) int {
	sum := 0
	for i := 0; i < len(c); i++ {
		sum += c[i]
	}
	return sum
}

func _getColumnsAfter(i, r int, piece string) []int {
	c := MyPlayer.Columns
	a := MyPlayer.Columns
	switch piece {
	case "I":
		switch r {
		case 0:
			if _isRight(i, 3) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}
				if pick < c[i+2] {
					pick = c[i+2]
				}
				if pick < c[i+3] {
					pick = c[i+3]
				}

				a[i] = pick + 1
				a[i+1] = pick + 1
				a[i+2] = pick + 1
				a[i+3] = pick + 1
			}
		case 1:
			a[i] = c[i] + 4
		}
	case "J":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}
				if pick < c[i+2] {
					pick = c[i+2]
				}

				a[i] = pick + 2
				a[i+1] = pick + 1
				a[i+2] = pick + 1
			}
		case 1:
			if _isRight(i, 1) {

				if c[i+1] >= c[i]+2 {
					a[i] = c[i+1] + 1
					a[i+1] = c[i+1] + 1
				} else {
					a[i] = c[i] + 3
					a[i+1] = c[i] + 3
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}
				if pick < c[i+2] {
					pick = c[i+2]
				}

				if pick == c[i+2] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}

				a[i] = pick + 1
				a[i+1] = pick + 3
			}
		}
	case "L":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}
				if pick < c[i+2] {
					pick = c[i+2]
				}

				a[i] = pick + 1
				a[i+1] = pick + 1
				a[i+2] = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}

				a[i] = pick + 3
				a[i+1] = pick + 1
			}
		case 2:
			if _isRight(i, 2) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}
				if pick < c[i+2] {
					pick = c[i+2]
				}

				if pick == c[i] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				if c[i] >= c[i+1]+2 {
					a[i] = c[i] + 1
					a[i+1] = c[i] + 1
				} else {
					a[i] = c[i+1] + 3
					a[i+1] = c[i+1] + 3
				}
			}
		}
	case "O":
		if _isRight(i, 1) {
			pick := c[i]
			if pick < c[i+1] {
				pick = c[i+1]
			}
			a[i] = pick + 2
			a[i+1] = pick + 2
		}
	case "S":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}
				if pick < c[i+2] {
					pick = c[i+2]
				}

				if pick == c[i] || pick == c[i+1] {
					a[i] = pick + 1
					a[i+1] = pick + 2
					a[i+2] = pick + 2
				} else {
					a[i] = pick
					a[i+1] = pick + 1
					a[i+2] = pick + 1
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}

				if pick == c[i+1] {
					a[i] = pick + 3
					a[i+1] = pick + 2
				} else {
					a[i] = pick + 2
					a[i+1] = pick + 1
				}
			}
		}
	case "T":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}
				if pick < c[i+2] {
					pick = c[i+2]
				}

				a[i] = pick + 1
				a[i+1] = pick + 2
				a[i+2] = pick + 1
			}
		case 1:
			if _isRight(i, 1) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}

				if pick == c[i] {
					a[i] = pick + 3
					a[i+1] = pick + 2
				} else {
					a[i] = pick + 2
					a[i+1] = pick + 1
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}
				if pick < c[i+2] {
					pick = c[i+2]
				}

				if pick == c[i+1] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}

				if pick == c[i+1] {
					a[i] = pick + 2
					a[i+1] = pick + 3

				} else {
					a[i] = pick + 1
					a[i+1] = pick + 2
				}
			}
		}
	case "Z":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}
				if pick < c[i+2] {
					pick = c[i+2]
				}

				if pick == c[i+1] || pick == c[i+2] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 1
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := c[i]
				if pick < c[i+1] {
					pick = c[i+1]
				}

				if pick == c[i] {
					a[i] = pick + 2
					a[i+1] = pick + 3
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 2
				}
			}
		}
	}
	return a
}

func _printMoves(pos Position) {
	var buffer bytes.Buffer
	for i := 0; i < pos.Rotation; i++ {
		buffer.WriteString("turnright,")
	}
	if pos.Rotation == 1 {
		CurrentPieceX = CurrentPieceX + 1
		if CurrentPiece == "I" {
			CurrentPieceX = CurrentPieceX + 1
		}
	}
	if CurrentPieceX > pos.X {
		for i := 0; i < CurrentPieceX-pos.X; i++ {
			buffer.WriteString("left,")
		}
	}
	if CurrentPieceX < pos.X {
		for i := 0; i < pos.X-CurrentPieceX; i++ {
			buffer.WriteString("right,")
		}
	}
	buffer.WriteString("drop")
	fmt.Println(buffer.String())
}
