package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// be caful when final print moves Need to check for "" and the do not add comma pos.Moves + "."+"drop"

	// rebuild gametest into visual and remove rotation and X from position Object

	// f.After(validPiece.CurrentX, validPiece.Rotation, piece.Name) pass just Piece

	// to dramaticly improve performance do a drop on empty rows when fixing the holes

	// when testing strategy output range of all games

	//see afect on performance if you pass Picks to this function
	//func (f Field) After(x, r int, piece string) Field {

	//find datatype in golang like dictionary (map) but keys only

	//refactor init for Field

	// add coordinats of 3, -1  and 4, -1 => O when testing the game
	// validate top positions shoud work already

	// if two same shapes make sure first one land in better place test for
	// this to prove ut

	// shift left or right

	// validate path to move piece to the place

	// when testing strategy output range of all games

	//10	2	3	5	76.283	103.609

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
				fmt.Print(pos.Moves + ",drop")
			}
		}
	}
}
