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

		CalculateMoves()
	}
	fmt.Println("Done")
}

func Test_55c29f6435ec1d070e2b66e9_40(t *testing.T) {
	field := Field{{true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, false, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {false, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, false, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {false, true, true, false, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, false}, {false, true, true, false, false, true, true, true, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer = &Player{Field: field}
	CurrentPiece = "T"

	pos := CalculateMoves()

	if pos.X != 8 || pos.Rotation != 3 {
		t.Fail()
		fmt.Println(pos.X, pos.Rotation)
	}
}

func Test_55c2d43635ec1d070e2b697c_63(t *testing.T) {
	field := Field{{false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, true, true, true, true, true}, {true, true, true, true, true, false, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, false, true, true, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, true, false, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, false, true, true, false, true, true, true, true}, {true, true, true, true, true, false, true, true, true, false}, {true, false, false, false, false, false, false, true, true, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer = &Player{Field: field}
	CurrentPiece = "I"
	NextPiece = "L"

	pos := CalculateMoves()

	if pos.X != 1 || pos.Rotation != 0 {
		t.Fail()
		fmt.Println(pos.X, pos.Rotation)
	}
}
