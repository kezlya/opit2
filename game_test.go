package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

var pieces = []string{"I", "J", "L", "O", "S", "T", "Z"}

func Benchmark_Score(b *testing.B) {
	// try to get everadge score to hight
	// try to get rounds to low
	// game lost count vs games won count

	//var rounds []int
	//var scores []int
	//var	victories []bool

	for n := 0; n < b.N; n++ {
		_asignSettings("field_width", "10")
		_asignSettings("field_height", "20")
		_asignSettings("player_names", "player1,player2")
		_asignSettings("your_bot", "player1")

		playGame()
		//rounds = append(rounds,round)
		//scores = append(scores,score)
		//victories = append(victories,isWinner)

		fmt.Println(Round, MyPlayer.Points)
	}
}

func playGame() {
	rand.Seed(time.Now().UTC().UnixNano())
	MyPlayer.Points = 0

	Round = 0

	fmt.Println(Height, Width)

	for Height > 17 && Pick <= Height {

		Round++

		Height = rand.Intn(50)
		Pick = rand.Intn(20)

		MyPlayer.Points++
	}
}

func Benchmark_getMoves(b *testing.B) {

	for n := 0; n < b.N; n++ {

		rand.Seed(time.Now().UTC().UnixNano())

		row1 := make([]string, 10, 10)
		for i := 0; i < 10; i++ {
			row1[i] = strconv.Itoa(rand.Intn(4))
		}
		row2 := make([]string, 10, 10)
		for i := 0; i < 10; i++ {
			row2[i] = strconv.Itoa(rand.Intn(4))
		}
		row3 := make([]string, 10, 10)
		for i := 0; i < 10; i++ {
			row3[i] = strconv.Itoa(rand.Intn(4))
		}

		_asignSettings("field_width", "10")
		_asignSettings("field_height", "20")
		_asignSettings("player_names", "player1,player2")
		_asignSettings("your_bot", "player1")
		_asignUpdates("game", "this_piece_type", pieces[rand.Intn(len(pieces))])
		_asignUpdates("game", "this_piece_position", "3,1")
		_asignUpdates("player1", "field", "0,0,0,1,1,1,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;"+strings.Join(row1, ",")+";"+strings.Join(row2, ",")+";"+strings.Join(row1, ","))

		_calculateMoves(500)
	}
	fmt.Println("Done")
}
