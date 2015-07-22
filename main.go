package main

import (
	"bufio"
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

func _getMoves(time int) {
	fmt.Println("left,left,left,drop")
	
	FindFitPositions(CurrentPiece,MyPlayer.Columns)
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
