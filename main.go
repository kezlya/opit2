package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	game := Game{}
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

func _getBurned(positions []Position) []Position {
	var burnedPos []Position
	for _, pos := range positions {
		if pos.IsBurn > 0 {
			burnedPos = append(burnedPos, pos)
		}
	}
	return burnedPos
}

func roundOne() {
	fmt.Println("drop")
}
