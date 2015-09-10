package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//if next move burn 2rows vs 1row current move preoritize next move up to 4

	//find datatype in golang like dictionary (map) but keys only

	//refactor init for Field

	// if two same shapes make sure first one land in better place test for
	// this to prove ut

	// to make it faster try to discard hiscore positions and position with big damage
	// damage >3 discard

	strategy := Strategy{
		Burn:   4,
		BHoles: 8,
		FHoles: 4,
		HighY:  2,
		Step:   2,
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
