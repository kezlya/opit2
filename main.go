package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
			_getMoves(time)
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

func _getMoves(time int) {
	if Round == 1 {
		fmt.Println("drop")
		return
	}

	var positions []Position
	switch CurrentPiece {
	case "I":
		positions = _fitsI(MyPlayer.Columns)
	case "J":
		positions = _fitsJ(MyPlayer.Columns)
	case "L":
		positions = _fitsL(MyPlayer.Columns)
	case "O":
		positions = _fitsO(MyPlayer.Columns)
	case "S":
		positions = _fitsS(MyPlayer.Columns)
	case "T":
		positions = _fitsT(MyPlayer.Columns)
	case "Z":
		positions = _fitsZ(MyPlayer.Columns)
	}

	if len(positions) == 0 {

		finalPosition := _chooseMinimumDamage()
		_printMoves(finalPosition)
	} else {
		finalPositionIndex := _chooseLowestPosition(positions)
		_printMoves(positions[finalPositionIndex])
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
			columsAfter := _getColumnsAfter(MyPlayer.Columns, i, r, CurrentPiece)
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
