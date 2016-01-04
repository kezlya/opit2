package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Happy New 2016 Year

// External API constants
const I, J, L, O, S, T, Z = "I", "J", "L", "O", "S", "T", "Z"
const settings, update, action = "settings", "update", "action"
const down = "down"
const left = "left"
const right = "right"
const turnleft = "turnleft"
const turnright = "turnright"
const drop = "drop"
const skip = "skip"

var strategy = Strategy{
	Burn:   1,
	BHoles: 12,
	FHoles: 12,
	CHoles: 2,
	HighY:  1,
	Step:   3,
}

func main() {
	game := Game{Strategy: strategy}
	consolereader := bufio.NewReader(os.Stdin)
	for {
		input, _ := consolereader.ReadString('\n')
		parts := strings.Split(strings.TrimSpace(input), " ")
		switch parts[0] {
		case settings:
			game.asignSettings(parts[1], parts[2])
		case update:
			game.asignUpdates(parts[1], parts[2], parts[3])
		case action:
			//time, _ := strconv.Atoi(parts[2])
			game.initPieces()
			pos := game.calculateMoves()
			if pos != nil && pos.Score != nil && pos.Moves != "" {
				if pos.Score.Total > 50 && game.MyPlayer.Skips > 0 {
					fmt.Println(skip)
				} else {
					fmt.Println(strings.TrimPrefix(pos.Moves, ",") + "," + drop)
				}
			} else {
				fmt.Println(drop)
			}
		}
	}
}
