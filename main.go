package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

// 15	1	4	8	78.792	103.333
// 15	2	3	7	78.708	106.042

	// if two same shapes make sure first one land in better place test for
	// this to prove ut

	// do not cover garbage rows

	// shift left or right

	game := Game{
		DamageK: 9,
		HoleK:   1,
		PostyK:  3,
		BurnK:   8,
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
