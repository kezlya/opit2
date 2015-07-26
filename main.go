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
			if Round == 1 {
				fmt.Println("drop")
			} else {
				_printMoves(_calculateMoves(time))
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
					Players[i].Field, Players[i].Columns = _convertField(cleanSource)
					break
				}
			}
		}
	}
}

func _convertField(rawField string) ([][]bool, []int) {
	rows := strings.Split(rawField, ";")
	Height = len(rows)
	var piks = make([]int, Width)
	var field = make([][]bool, Height)
	for rowIndex, row := range rows {
		y := Height - rowIndex
		var colums = make([]bool, Width)
		for columIndex, colum := range strings.Split(row, ",") {
			if colum == "2" {
				colums[columIndex] = true
				if y > piks[columIndex] {
					piks[columIndex] = y
				}
			} else {
				colums[columIndex] = false
			}
		}
		field[rowIndex] = colums
	}
	return field, piks
}

func _calculateMoves(time int) Position {
	roofIsnear := false
	savePlay := false
	for _, pick := range MyPlayer.Columns {
		if Height-pick <= 5 {
			roofIsnear = true
		}
		if Height-pick >= 10 {
			savePlay = true
		}
	}

	//TODO very bad game http://theaigames.com/competitions/ai-block-battle/games/55b40c2e35ec1d487cd5e908

	//TODO fix round 41 http://theaigames.com/competitions/ai-block-battle/games/55b403af35ec1d487cd5e8aa

	//TODO fix round 51 http://theaigames.com/competitions/ai-block-battle/games/55b3ea7b35ec1d487cd5e77a

	//TODO fix round 76 http://theaigames.com/competitions/ai-block-battle/games/55b3eea335ec1d487cd5e7a5

	//TODO: I should not behave as minimum damadge need to use best fit from before

	//TODO: choose plasements clother to the wall

	var goldenIndex int
	allPositins, bestScore := _getAllPossiblePositions()
	// perfect fit when damadge 4

	if roofIsnear {
		lowestY := 1000
		for i, pos := range allPositins {
			//fmt.Println(pos.Rotation, pos.X,pos.Damadge,pos.MaxY,lowestY)
			if pos.MaxY < lowestY {
				lowestY = pos.MaxY
				goldenIndex = i

				if pos.Damadge < allPositins[goldenIndex].Damadge {
					goldenIndex = i
				}
				//fmt.Println("************")
				//fmt.Println(pos.Rotation, pos.X,pos.Damadge,pos.MaxY,lowestY)
				//fmt.Println(pos.ColumnsAfter)
			}
		}
		return allPositins[goldenIndex]
	} else if savePlay {
		//old code here
		noDamadgePositions := _getNoDamadgePositions(allPositins)
		if len(noDamadgePositions) > 0 {
			tempMaxY := 1000
			for i, pos := range noDamadgePositions {
				if pos.MaxY < tempMaxY {
					tempMaxY = pos.MaxY
					goldenIndex = i
				}
			}
			return noDamadgePositions[goldenIndex]
		} else {
			//redundand code below
			bestPositions := _getBestScorePositions(allPositins, bestScore)
			tempDamadge := 1000
			for i, pos := range bestPositions {
				//check if it burns rows

				if pos.Damadge < tempDamadge {
					tempDamadge = pos.Damadge
					goldenIndex = i
				}
			}
			return bestPositions[goldenIndex]
			//////////end redundand code

		}

	} else {
		bestPositions := _getBestScorePositions(allPositins, bestScore)
		tempDamadge := 1000
		for i, pos := range bestPositions {
			//check if it burns rows

			if pos.Damadge < tempDamadge {
				tempDamadge = pos.Damadge
				goldenIndex = i
			}
		}
		return bestPositions[goldenIndex]
	}

	//lowestY
	/*lowestY := 1000
	for i, pos := range allPositins {
		if roofIsnear {
			if pos.MaxY < lowestY {
				lowestY = pos.MaxY
				goldenIndex = i
			}
			if pos.MaxY == lowestY {

			}
		} else {
			if pos.Damadge == minDamadge && pos.MaxY < lowestY {
				goldenIndex = i
				lowestY = pos.MaxY
			}
		}
	}*/

	//TODO absolute lowest when close to the roof or tool bildings

	//TODO look into the next piece

}

func _getBestScorePositions(positions []Position, bestScore int) []Position {
	var result []Position
	for _, pos := range positions {
		if pos.Score == bestScore {
			result = append(result, pos)
		}
		//TODO: predict next move
	}
	return result
}

func _getNoDamadgePositions(positions []Position) []Position {
	var result []Position
	for _, pos := range positions {
		if pos.Damadge == 4 {
			result = append(result, pos)
		}
		//TODO: predict next move
	}
	return result
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
