package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// f.After(validPiece.CurrentX, validPiece.Rotation, piece.Name) pass just Piece

	// to dramaticly improve performance do a drop on empty rows when fixing the holes

	// when testing strategy output range of all games

	//see afect on performance if you pass Picks to this function
	//func (f Field) After(x, r int, piece string) Field {

	//find datatype in golang like dictionary (map) but keys only

	//refactor init for Field

	// if two same shapes make sure first one land in better place test for
	// this to prove ut

	// when testing strategy output range of all games

	// to make it faster try to discard hiscore positions and position with big damage

	strategy := Strategy{
		DamageK: 10,
		StepK:   2,
		PostyK:  3,
		BurnK:   5,
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
			game.CurrentPiece.InitSpace(Cell{X: game.X, Y: game.Height + game.Y})
			game.NextPiece.InitSpace(Cell{X: 3, Y: game.Height + game.Y})
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
