package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var strategy = Strategy{
	Burn:   2,
	BHoles: 12,
	FHoles: 10,
	CHoles: 2,
	HighY:  2,
	Step:   3,
}

func main() {
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
			game.initPieces()
			pos := game.calculateMoves()
			if pos.Moves == "" {
				fmt.Println("drop")
			} else {
				fmt.Println(pos.Moves + ",drop")
			}
		}
	}
}
