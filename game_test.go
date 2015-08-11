package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
	"encoding/csv"
	"os"
)

var initialField = Field{{false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
var pieces = []string{"I", "J", "L", "O", "S", "T", "Z"}

func Benchmark_Play100(b *testing.B) {
	// try to get everadge score to hight
	// try to get rounds to low
	// game lost count vs games won count

	//var rounds []int
	//var scores []int

	for n := 0; n < b.N; n++ {
		
		records := [][]string{}
		for raz := 0; raz < 100; raz++ {
			roud, score := playGame()
			records = append(records,[]string{strconv.Itoa(roud),strconv.Itoa(score)})
			fmt.Println(roud, score)
		}
		filename := "100_"+strconv.FormatInt(int64(time.Now().UTC().UnixNano()),10)
		saveReport(filename,records)
	}
}

func Benchmark_Play1(b *testing.B) {
	// try to get everadge score to hight
	// try to get rounds to low
	// game lost count vs games won count

	//var rounds []int
	//var scores []int

	for n := 0; n < b.N; n++ {	
		//records := [][]string{}
		for raz := 0; raz < 1; raz++ {
			roud, score := playGame()
			//records = append(records,[]string{strconv.Itoa(roud),strconv.Itoa(score)})
			fmt.Println("game:", roud, score)
		}
		//filename := "1_"+strconv.FormatInt(int64(time.Now().UTC().UnixNano()),10)
		//saveReport(filename,records)
	}
}

func playGame() (int, int) {
	rand.Seed(time.Now().UTC().UnixNano())

	Round = 0
	NextPiece = pieces[rand.Intn(len(pieces))]
	MyPlayer.Points = 0
	position := &Position{}
	position.FieldAfter = initialField

	for position != nil {
		rand.Seed(time.Now().UTC().UnixNano())
		CurrentPiece = NextPiece
		NextPiece = pieces[rand.Intn(len(pieces))]
		if position.IsBurn>0 {
			MyPlayer.Combo++
			MyPlayer.Points += MyPlayer.Combo-1
			switch position.IsBurn {
				case 1:
					MyPlayer.Points += 1
				case 2:
					MyPlayer.Points += 3
				case 3:
					MyPlayer.Points += 6
				case 4:
					MyPlayer.Points += 12
			}
		} else {
			MyPlayer.Combo = 0
		}
		//fmt.Println(Round,position.IsBurn, MyPlayer.Combo, MyPlayer.Points)
		

		position.FieldAfter.Burn()
		MyPlayer.Field = position.FieldAfter

		position = _calculateMoves()
		Round++
		//add garbage raws
	}

	return Round, MyPlayer.Points
}

func saveReport(fileName string, records [][]string) {
	csvfile, err := os.Create("output/"+fileName+".csv")
          if err != nil {
                  fmt.Println("Error:", err)
                  return
          }
          defer csvfile.Close()

          
          writer := csv.NewWriter(csvfile)
          for _, record := range records {
                  err := writer.Write(record)
                  if err != nil {
                          fmt.Println("Error:", err)
                          return
                  }
          }
          writer.Flush()
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

		_calculateMoves()
	}
	fmt.Println("Done")
}

func Test_55c29f6435ec1d070e2b66e9_40(t *testing.T) {
	field := Field{{true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, false, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {false, true, true, false, true, true, true, true, true, true}, {true, true, true, true, true, true, true, true, true, false}, {true, true, true, true, true, true, false, true, false, true}, {true, true, true, true, true, true, false, true, true, true}, {true, true, true, true, true, true, true, false, true, true}, {true, true, true, true, true, true, true, false, true, true}, {false, true, true, false, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, true}, {false, true, true, true, true, true, true, true, true, false}, {false, true, true, false, false, true, true, true, false, false}, {false, true, true, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
	MyPlayer = &Player{Field: field}
	CurrentPiece = "T"

	pos := _calculateMoves()

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

	pos := _calculateMoves()

	if pos.X != 1 || pos.Rotation != 0 {
		t.Fail()
		fmt.Println(pos.X, pos.Rotation)
	}
}
