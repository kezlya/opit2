package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// if two same shapes make sure first one land in better place test for
	// this to prove ut

	// shift left or right

	//validate path to move piece to the place

	//investigate performance by returning fieldAfter nil

	//init position

	strategy := Strategy{
		DamageK: 9,
		StepK:   1,
		PostyK:  3,
		BurnK:   8,
	}
	game := Game{Strategy: strategy}
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
