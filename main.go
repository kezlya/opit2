package main

import (
	"bufio"
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
			//time, _ := strconv.Atoi(parts[2])
			if Round == 1 {
				_roundOne()
			} else {
				_printMoves(_calculateMoves())
			}
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
		OriginalHeight, _ = strconv.Atoi(value)
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
			cleanSource := strings.Replace(value, ";3,3,3,3,3,3,3,3,3,3", "", OriginalHeight)
			for i, p := range Players {
				if p.Name == who {
					Players[i].Field = _convertField(cleanSource)
					picks := Players[i].Field.Picks()
					y := picks.Max()
					rowsLeft := Players[i].Field.Height() - y
					if rowsLeft >= 10 {
						Players[i].State = "safe"
					} else {
						if rowsLeft < 5 {
							Players[i].State = "dangerous"
						} else {
							Players[i].State = "normal"
						}
					}
					break
				}
			}
		}
	}
}

func _convertField(rawField string) Field {
	rows := strings.Split(rawField, ";")
	Height = len(rows)
	var field = make([][]bool, Height)
	for rowIndex, row := range rows {
		y := Height - rowIndex
		var colums = make([]bool, Width)
		for columIndex, value := range strings.Split(row, ",") {
			if value == "2" {
				colums[columIndex] = true
			} else {
				colums[columIndex] = false
			}
		}
		field[y-1] = colums
	}
	return field
}
