package main

import (
//"fmt"
)

func _getAllPossiblePositions() ([]Position, int) {
	var positions []Position
	bestScore := 1000
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
			columsAfter, maxY := _getColumnsAfter(MyPlayer.Columns, i, r, CurrentPiece)
			damage := _sum(columsAfter) - columnsSum
			if damage > 0 {
				p := Position{
					Rotation:     r,
					X:            i,
					MaxY:         maxY,
					Damadge:      damage,
					Score:        maxY + damage,
					ColumnsAfter: columsAfter}
				positions = append(positions, p)
				if maxY+damage < bestScore {
					bestScore = maxY + damage
				}
			}
		}
	}
	return positions, bestScore
}

func _calculateMoves(time int) Position {
	roofIsnear := false
	savePlay := false
	for _, pick := range MyPlayer.Columns {
		if Height-pick <= 5 {
			//fmt.Println(Height,pick)
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
				//fmt.Println(_isHole(pos.ColumnsAfter))
			}
		}
		return allPositins[goldenIndex]
	}

	if savePlay {
		//old code here
		noDamadgePositions := _getNoDamadgePositions(allPositins)
		if len(noDamadgePositions) > 0 {
			tempMaxY := 1000
			for i, pos := range noDamadgePositions {

				if (!_isHole(pos.ColumnsAfter)) && pos.MaxY < tempMaxY {
					tempMaxY = pos.MaxY
					goldenIndex = i
					//fmt.Println(_isHole(pos.ColumnsAfter))
				}
			}
			return noDamadgePositions[goldenIndex]
		}
	}

	bestPositions := _getBestScorePositions(allPositins, bestScore)
	tempDamadge := 1000
	for i, pos := range bestPositions {
		//check if it burns rows

		if (!_isHole(pos.ColumnsAfter)) && pos.Damadge < tempDamadge {
			tempDamadge = pos.Damadge
			goldenIndex = i
			//fmt.Println(_isHole(pos.ColumnsAfter))
		}
	}
	return bestPositions[goldenIndex]

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

func _isHole(cols []int) bool {
	for i, c := range cols {
		if _isRight(i, 1) && (c-cols[i+1] < -2 || c-cols[i+1] > 2) && CurrentPiece != "I" && NextPiece != "I" {
			return true
		}
	}
	return false
}

func _isRight(i, right int) bool {
	if i+right < Width {
		return true
	}
	return false
}

func _isUp(i, up int) bool {
	if i+up < Height {
		return true
	}
	return false
}

func _isLeft(i, left int) bool {
	if i-left > 0 {
		return true
	}
	return false
}

func _getColumnsAfter(c []int, i, r int, piece string) ([]int, int) {
	a := make([]int, len(c))
	copy(a, c[:])
	y := 0
	switch piece {
	case "I":
		switch r {
		case 0:
			if _isRight(i, 3) {
				pick := _getPick(i, 3)
				a[i] = pick + 1
				a[i+1] = pick + 1
				a[i+2] = pick + 1
				a[i+3] = pick + 1
				y = pick + 1
			}
		case 1:
			a[i] = c[i] + 4
			y = c[i] + 4
		}
	case "J":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				a[i] = pick + 2
				a[i+1] = pick + 1
				a[i+2] = pick + 1
				y = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				if c[i+1] >= c[i]+2 {
					a[i] = c[i+1] + 1
					a[i+1] = c[i+1] + 1
					y = c[i+1] + 1
				} else {
					a[i] = c[i] + 3
					a[i+1] = c[i] + 3
					y = c[i] + 3
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i+2] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				a[i] = pick + 1
				a[i+1] = pick + 3
				y = pick + 3
			}
		}
	case "L":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				a[i] = pick + 1
				a[i+1] = pick + 1
				a[i+2] = pick + 2
				y = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				a[i] = pick + 3
				a[i+1] = pick + 1
				y = pick + 3
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				if c[i] >= c[i+1]+2 {
					a[i] = c[i] + 1
					a[i+1] = c[i] + 1
					y = c[i] + 1
				} else {
					a[i] = c[i+1] + 3
					a[i+1] = c[i+1] + 3
					y = c[i+1] + 3
				}
			}
		}
	case "O":
		if _isRight(i, 1) {
			pick := _getPick(i, 1)
			a[i] = pick + 2
			a[i+1] = pick + 2
			y = pick + 2
		}
	case "S":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i] || pick == c[i+1] {
					a[i] = pick + 1
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i+1] {
					a[i] = pick + 3
					a[i+1] = pick + 2
					y = pick + 3
				} else {
					a[i] = pick + 2
					a[i+1] = pick + 1
					y = pick + 2
				}
			}
		}
	case "T":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				a[i] = pick + 1
				a[i+1] = pick + 2
				a[i+2] = pick + 1
				y = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i] {
					a[i] = pick + 3
					a[i+1] = pick + 2
					y = pick + 3
				} else {
					a[i] = pick + 2
					a[i+1] = pick + 1
					y = pick + 2
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i+1] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i+1] {
					a[i] = pick + 2
					a[i+1] = pick + 3
					y = pick + 3

				} else {
					a[i] = pick + 1
					a[i+1] = pick + 2
					y = pick + 2
				}
			}
		}
	case "Z":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i+1] || pick == c[i+2] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 1
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick
					y = pick + 1
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i] {
					a[i] = pick + 2
					a[i+1] = pick + 3
					y = pick + 3
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 2
					y = pick + 2
				}
			}
		}
	}
	return a, y
}

func _fieldAfter(f [][]bool, i, r int, piece string) ([][]bool, int) {
	a := make([][]bool, len(f))
	for i, row := range f {
		a[i] = make([]bool, len(row))
		copy(a[i], row[:])
	}

	switch piece {
	case "I":
		switch r {
		case 0:
			if _isRight(i, 3) && _isUp() {
				pick := _getPick(i, 3)

				a[pick][i] = true
				a[pick][i+1] = true
				a[pick][i+2] = true
				a[pick][i+3] = true
			}
		case 1:
			a[pick][i] = true
			a[pick+1][i] = true
			a[pick+2][i] = true
			a[pick+3][i] = true
		}
	case "J":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				a[i] = pick + 2
				a[i+1] = pick + 1
				a[i+2] = pick + 1
				y = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				if c[i+1] >= c[i]+2 {
					a[i] = c[i+1] + 1
					a[i+1] = c[i+1] + 1
					y = c[i+1] + 1
				} else {
					a[i] = c[i] + 3
					a[i+1] = c[i] + 3
					y = c[i] + 3
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i+2] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				a[i] = pick + 1
				a[i+1] = pick + 3
				y = pick + 3
			}
		}
	case "L":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				a[i] = pick + 1
				a[i+1] = pick + 1
				a[i+2] = pick + 2
				y = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				a[i] = pick + 3
				a[i+1] = pick + 1
				y = pick + 3
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				if c[i] >= c[i+1]+2 {
					a[i] = c[i] + 1
					a[i+1] = c[i] + 1
					y = c[i] + 1
				} else {
					a[i] = c[i+1] + 3
					a[i+1] = c[i+1] + 3
					y = c[i+1] + 3
				}
			}
		}
	case "O":
		if _isRight(i, 1) {
			pick := _getPick(i, 1)
			a[i] = pick + 2
			a[i+1] = pick + 2
			y = pick + 2
		}
	case "S":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i] || pick == c[i+1] {
					a[i] = pick + 1
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i+1] {
					a[i] = pick + 3
					a[i+1] = pick + 2
					y = pick + 3
				} else {
					a[i] = pick + 2
					a[i+1] = pick + 1
					y = pick + 2
				}
			}
		}
	case "T":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				a[i] = pick + 1
				a[i+1] = pick + 2
				a[i+2] = pick + 1
				y = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i] {
					a[i] = pick + 3
					a[i+1] = pick + 2
					y = pick + 3
				} else {
					a[i] = pick + 2
					a[i+1] = pick + 1
					y = pick + 2
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i+1] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i+1] {
					a[i] = pick + 2
					a[i+1] = pick + 3
					y = pick + 3

				} else {
					a[i] = pick + 1
					a[i+1] = pick + 2
					y = pick + 2
				}
			}
		}
	case "Z":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i+1] || pick == c[i+2] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 1
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick
					y = pick + 1
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i] {
					a[i] = pick + 2
					a[i+1] = pick + 3
					y = pick + 3
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 2
					y = pick + 2
				}
			}
		}
	}
	return a, y
}

func _getPick(i, v int) int {
	pick := MyPlayer.Columns[i]
	for n := 1; n <= v; n++ {
		if pick < MyPlayer.Columns[i+n] {
			pick = MyPlayer.Columns[i+n]
		}
	}
	return pick
}

func _sum(c []int) int {
	sum := 0
	for i := 0; i < len(c); i++ {
		sum += c[i]
	}
	return sum
}
