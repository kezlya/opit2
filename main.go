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
			_calculateMoves(time)
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

func _calculateMoves(time int) {
	if Round == 1 {
		fmt.Println("drop")
		return
	}
	var goldenIndex int
	allPositins, minDamadge := _getAllPossiblePositions()



	/*
		if minDamadge == 4{
			//we have best fit
		}
	*/

	//lowestY
	lowestY := 1000
	for i, pos := range allPositins {
		if pos.Damadge == minDamadge && pos.MaxY < lowestY {
			lowestY = pos.MaxY
			goldenIndex = i
		}
	}
	
	//TODO absolute lowest when close to the roof or tool bildings
	
	//TODO look into the next piece

	_printMoves(allPositins[goldenIndex])
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
