package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// after testing one found out that same game plays differently due no order in slice of positions.
	// planing to impliment micro coofecent into Piece.Totalscore() that adding score to determent if pice in on a side of the field then 0 if more to center then 5 points
	// idea was to push piecec closer to the sides of the field but it will break covered hols or maybe not
	//

	//find datatype in golang like dictionary (map) but keys only

	//refactor init for Field

	// to make it faster try to discard hiscore positions and position with big damage
	// damage >3 discard

	//new statistic showed that this is the best strategy after new inprovements
	//4	14	5	1	2	3	110	73	188	139	107	203
	strategy := Strategy{
		Burn:   4,
		BHoles: 14,
		FHoles: 5,
		CHoles: 1,
		HighY:  2,
		Step:   3,
	}
	game := Game{
		Strategy:     strategy,
		CurrentPiece: Piece{},
		NextPiece:    Piece{}}

	consolereader := bufio.NewReader(os.Stdin)
	for {
		input, _ := consolereader.ReadString('\n')
		parts := strings.Split(strings.TrimSpace(input), " ")
		switch parts[0] {
		case "settings":
			game.asignSettings(parts[1], parts[2])
		case "update":
			game.asignUpdates(parts[1], parts[2], parts[3])
		case "action":
			//time, _ := strconv.Atoi(parts[2])
			game.CurrentPiece.InitSpace(Cell{X: game.X, Y: game.MyPlayer.Field.Height() + game.Y})
			game.NextPiece.InitSpace(Cell{X: 3, Y: game.MyPlayer.Field.Height() + game.Y})
			if game.Round == 1 {
				fmt.Println("drop")
			} else {
				pos := game.calculateMoves()
				if pos.Moves == "" {
					fmt.Println("drop")
				} else {
					fmt.Println(pos.Moves + ",drop")
				}
			}
		}
	}
}
