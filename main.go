package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//best score so far d4_h1_y3_b3
	//best score so far d5_h1_y3_b3
	game := Game{
		DamageK: 5,
		HoleK:   1,
		PostyK:  3,
		BurnK:   3,
	}
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
