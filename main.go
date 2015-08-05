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
	height := len(rows)
	var field = make([][]bool, height)
	for rowIndex, row := range rows {
		y := height - rowIndex
		cells := strings.Split(row, ",")
		var colums = make([]bool, len(cells))
		for columIndex, value := range cells {
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

func _calculateMoves() Position {
	//TODO: choose plasements clother to the wall

	if MyPlayer.Combo > 0 || MyPlayer.State == "dangerous" {
		pos, isFound := _keepUpBurn()
		if isFound {
			return pos
		}
	}

	if MyPlayer.State == "safe" {
		shortField := MyPlayer.Field.Trim(2)
		shortPositions := shortField.Positions(CurrentPiece)
		OrderedBy(SCORE, DAMAGE, LOWY).Sort(shortPositions)
		return shortPositions[0] //TODO: predict next piece
	}

	positions := MyPlayer.Field.Positions(CurrentPiece)

	if MyPlayer.State == "normal" {
		OrderedBy(SCORE, DAMAGE, LOWY).Sort(positions)
		//TODO check if burn and check if no damage
		//TODO: predict next piece
	}

	// play save try to burn rows and get lowest Y
	if MyPlayer.State == "dangerous" {
		OrderedBy(LOWY, SCORE, DAMAGE).Sort(positions)
		//TODO: predict next piece
	}

	return positions[0]
}

func _keepUpBurn() (Position, bool) {
	var emptyPos Position
	var burnedPos []Position
	positions := MyPlayer.Field.Positions(CurrentPiece)

	for _, pos := range positions {
		if pos.IsBurn > 0 {
			burnedPos = append(burnedPos, pos)
		}
	}
	burnedPosTotal := len(burnedPos)

	if burnedPosTotal == 1 {
		return burnedPos[0], true
	}

	if burnedPosTotal > 1 {
		OrderedBy(SCORE, DAMAGE).Sort(positions)
		bIndex := 0
		for current_i, pos := range burnedPos {
			pos.FieldAfter.Burn()
			nextPiecePositions := pos.FieldAfter.Positions(NextPiece)
			for _, nextPos := range nextPiecePositions {
				if nextPos.IsBurn > 0 {
					bIndex = current_i
					break
				}
			}
		}
		return burnedPos[bIndex], true
	}
	return emptyPos, false
}

func _roundOne() {
	fmt.Println("drop")
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
