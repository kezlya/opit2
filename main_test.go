package main

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Benchmark_getMoves(b *testing.B) {

	rand.Seed(time.Now().UTC().UnixNano())
	pieces := []string{"I", "J", "L", "O", "S", "T", "Z"}

	for n := 0; n < b.N; n++ {

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

		_getMoves(500)
	}

}
