package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

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
			if game.Round == 1 {
				roundOne()
			} else {
				game.printMoves()
			}
		}
	}
}

func roundOne() {
	fmt.Println("drop")
}
