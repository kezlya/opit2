package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

var initialField = Field{{false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
var pieces = []string{"I", "J", "L", "O", "S", "T", "Z"}

func Benchmark_moves(b *testing.B) {

	for n := 0; n < b.N; n++ {

		game := Game{}
		game.asignSettings("field_width", "10")
		game.asignSettings("field_height", "20")
		game.asignSettings("player_names", "player1,player2")
		game.asignSettings("your_bot", "player1")

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

		game.asignUpdates("game", "this_piece_type", pieces[rand.Intn(len(pieces))])
		game.asignUpdates("game", "this_piece_position", "3,1")
		game.asignUpdates("player1", "field", "0,0,0,1,1,1,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;"+strings.Join(row1, ",")+";"+strings.Join(row2, ",")+";"+strings.Join(row1, ","))

		game.calculateMoves()
	}
}

func Benchmark_many(b *testing.B) {
	for n := 0; n < b.N; n++ {
		playGames(100, true)
	}
}

func Benchmark_one(b *testing.B) {
	for n := 0; n < b.N; n++ {
		playGames(1, false)
	}
}

func Benchmark_strategy(b *testing.B) {
	game := Game{}
	for n := 0; n < b.N; n++ {
		game.Round = n + 1
		//go fake(game.Round, "hello"+strconv.Itoa(n))
		//Round = (n + 1) * 2
		//fake(game.Round, "privet"+strconv.Itoa(n))
	}
}

func playGames(amount int, saveReport bool) {
	records := [][]string{}
	for i := 0; i < amount; i++ {
		roud, score := playGame()
		records = append(records, []string{strconv.Itoa(roud), strconv.Itoa(score)})
		fmt.Println(roud, score)
	}
	if saveReport {
		filename := strconv.Itoa(amount) + "_" +
			"_" + "avaradge_round" + "_" + "avaradge_score" + "_" +
			strconv.FormatInt(int64(time.Now().UTC().UnixNano()), 10)
		save(filename, records)
		fmt.Println(filename)
	}
}

func playGame() (int, int) {
	rand.Seed(time.Now().UTC().UnixNano())
	g := Game{}
	g.asignSettings("player_names", "player1,player2")
	g.asignSettings("your_bot", "player1")
	g.Round = 0
	g.NextPiece = pieces[rand.Intn(len(pieces))]
	g.MyPlayer.Points = 0
	position := &Position{}
	position.FieldAfter = initialField

	for position != nil {
		assignPieces(&g)
		applyPoints(&g, position)
		position.FieldAfter.Burn()
		g.MyPlayer.Field = position.FieldAfter
		addSolidLines(&g)
		//add garbage raws
		g.Round++

		position = g.calculateMoves()
	}

	return g.Round, g.MyPlayer.Points
}

func assignPieces(g *Game) {
	rand.Seed(time.Now().UTC().UnixNano())
	g.CurrentPiece = g.NextPiece
	g.NextPiece = pieces[rand.Intn(len(pieces))]
}

func applyPoints(g *Game, pos *Position) {
	if pos.IsBurn > 0 {
		g.MyPlayer.Combo++
		g.MyPlayer.Points += g.MyPlayer.Combo - 1
		switch pos.IsBurn {
		case 1:
			g.MyPlayer.Points += 1
		case 2:
			g.MyPlayer.Points += 3
		case 3:
			g.MyPlayer.Points += 6
		case 4:
			g.MyPlayer.Points += 12
		}
	} else {
		g.MyPlayer.Combo = 0
	}
}

func addSolidLines(g *Game) {
	r := g.Round % 20
	if r == 0 {
		g.MyPlayer.Field = g.MyPlayer.Field[1:]
	}
}

func save(fileName string, records [][]string) {
	csvfile, err := os.Create("output/" + fileName + ".csv")
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
