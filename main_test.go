package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func _setup() {
	_asignSettings("field_width", "10")
	_asignSettings("field_height", "20")
	_asignSettings("player_names", "player1,player2")
	_asignSettings("your_bot", "player1")
}

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


/*
func Test_getMoves(t *testing.T) {
	_setup()
	
	_asignUpdates("game", "this_piece_type", "J")
	_asignUpdates("player1", "field", "0,0,0,1,1,1,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,2,0,2,0,0,2,0,0;0,0,2,2,2,0,2,2,0,0;0,0,2,2,2,2,2,2,2,0;0,2,2,2,2,2,2,2,2,0;0,2,2,2,2,2,2,2,2,0;0,2,0,2,2,2,2,2,2,0;0,2,2,2,2,2,2,2,2,0;0,2,2,2,2,2,2,2,2,0;0,2,2,2,2,2,2,2,2,0;0,2,2,2,2,2,2,2,2,0;2,2,2,2,2,2,2,2,2,0;2,2,2,2,2,0,2,2,2,2;3,3,3,3,3,3,3,3,3,3;3,3,3,3,3,3,3,3,3,3;3,3,3,3,3,3,3,3,3,3")
	_getMoves(500)
		fmt.Println("should: ", "turnright,left,left,drop")
		t.Fail()

}
*/



//http://theaigames.com/competitions/ai-block-battle/games/55b270a535ec1d487cd5d5b5 round 38
func Test_chooseMinimumDamageJ0(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "J")
	MyPlayer.Columns = []int{2, 9, 12, 11, 12, 10, 11, 12, 10, 1}
	result := _chooseMinimumDamage()
	should := Position{Rotation: 1, X: 1}
	if result.Rotation != should.Rotation || result.X != should.X {
		fmt.Println("result: ", result)
		fmt.Println("should: ", should)
		t.Fail()
	}
}

/*
func Test_chooseMinimumDamageJ1(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "J")
	MyPlayer.Columns = []int{7, 6, 4, 7, 6, 7, 8, 6, 3, 0}
	result := _chooseMinimumDamage()
	should := Position{Rotation: 0, X: 2}
	if result.Rotation != should.Rotation || result.X != should.X {
		fmt.Println("result: ", result)
		fmt.Println("should: ", should)
		t.Fail()
	}
}

func Test_chooseMinimumDamageJ2(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "J")
	MyPlayer.Columns = []int{9, 10, 2, 3, 2, 5, 4, 2, 1, 0}
	result := _chooseMinimumDamage()
	should := Position{Rotation: 3, X: 8}
	if result.Rotation != should.Rotation || result.X != should.X {
		fmt.Println("result: ", result)
		fmt.Println("should: ", should)
		t.Fail()
	}
}

func Test_chooseMinimumDamageZ1(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "Z")
	MyPlayer.Columns = []int{0, 0, 5, 4, 3, 5, 4, 3, 1, 1}
	result := _chooseMinimumDamage()
	should := Position{Rotation: 1, X: 0}
	if result.Rotation != should.Rotation || result.X != should.X {
		fmt.Println("result: ", result)
		fmt.Println("should: ", should)
		t.Fail()
	}
}

func Test_chooseMinimumDamageZ2(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "Z")
	MyPlayer.Columns = []int{2, 0, 3, 3, 3, 2, 4, 3, 2, 1}
	result := _chooseMinimumDamage()
	should := Position{Rotation: 0, X: 7}
	if result.Rotation != should.Rotation || result.X != should.X {
		fmt.Println("result: ", result)
		fmt.Println("should: ", should)
		t.Fail()
	}
}

func Test_chooseMinimumDamageZ3(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "Z")
	MyPlayer.Columns = []int{3, 1, 3, 5, 4, 3, 6, 6, 2, 0}
	result := _chooseMinimumDamage()
	should := Position{Rotation: 1, X: 1}
	if result.Rotation != should.Rotation || result.X != should.X {
		fmt.Println("result: ", result)
		fmt.Println("should: ", should)
		t.Fail()
	}
}

func Test_chooseMinimumDamageS1(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "S")
	MyPlayer.Columns = []int{2, 3, 4, 4, 2, 0, 5, 5, 1, 2}
	result := _chooseMinimumDamage()
	should := Position{Rotation: 0, X: 0}
	if result.Rotation != should.Rotation || result.X != should.X {
		fmt.Println("result: ", result)
		fmt.Println("should: ", should)
		t.Fail()
	}
}

func Test_chooseMinimumDamageS2(t *testing.T) {
	_setup()
	_asignUpdates("game", "this_piece_type", "S")
	MyPlayer.Columns = []int{2, 0, 3, 1, 3, 4, 0, 3, 3, 3}
	result := _chooseMinimumDamage()
	should := Position{Rotation: 0, X: 7}
	if result.Rotation != should.Rotation || result.X != should.X {
		fmt.Println("result: ", result)
		fmt.Println("should: ", should)
		t.Fail()
	}
}*/
