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
	dK := 5
	hK := 1
	yK := 3
	bK := 3
	for n := 0; n < b.N; n++ {
		playGames(dK, hK, yK, bK, 100, true)
	}
}

func Benchmark_one(b *testing.B) {
	dK := 5
	hK := 1
	yK := 3
	bK := 3
	for n := 0; n < b.N; n++ {
		playGames(dK, hK, yK, bK, 1, false)
	}
}

func Benchmark_best_strategy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go playGames(7, 1, 2, 4, 10, false)
		go playGames(6, 1, 5, 3, 10, false)
		go playGames(7, 3, 4, 6, 10, false)
		go playGames(7, 1, 3, 4, 10, false)
		go playGames(4, 1, 2, 5, 10, false)
		go playGames(6, 2, 4, 5, 10, false)
		go playGames(6, 1, 3, 5, 10, false)
		go playGames(7, 2, 5, 8, 10, false)
		go playGames(7, 1, 4, 7, 10, false)
		go playGames(7, 2, 4, 6, 10, false)
		go playGames(7, 3, 6, 7, 10, false)
		go playGames(5, 1, 2, 6, 10, false)
		go playGames(6, 2, 3, 6, 10, false)
		go playGames(7, 2, 3, 3, 10, false)
		go playGames(7, 1, 6, 4, 10, false)
		time.Sleep(80000000000)
		go playGames(7, 2, 5, 5, 10, false)
		go playGames(6, 2, 4, 3, 10, false)
		go playGames(6, 2, 2, 5, 10, false)
		go playGames(7, 2, 5, 3, 10, false)
		go playGames(5, 2, 5, 6, 10, false)
		go playGames(7, 1, 5, 5, 10, false)
		go playGames(7, 2, 3, 5, 10, false)
		go playGames(6, 1, 2, 7, 10, false)
		go playGames(5, 1, 3, 4, 10, false)
		go playGames(6, 2, 4, 7, 10, false)
		go playGames(7, 3, 3, 8, 10, false)
		go playGames(5, 1, 2, 7, 10, false)
		go playGames(7, 2, 2, 7, 10, false)
		go playGames(5, 2, 3, 7, 10, false)
		go playGames(5, 2, 3, 6, 10, false)
		time.Sleep(80000000000)
		go playGames(6, 1, 3, 7, 10, false)
		go playGames(5, 1, 5, 4, 10, false)
		go playGames(6, 1, 5, 8, 10, false)
		go playGames(7, 2, 4, 7, 10, false)
		go playGames(5, 2, 3, 4, 10, false)
		go playGames(5, 2, 3, 5, 10, false)
		go playGames(6, 1, 4, 4, 10, false)
		go playGames(7, 2, 6, 8, 10, false)
		go playGames(7, 1, 5, 7, 10, false)
		go playGames(5, 1, 3, 8, 10, false)
		go playGames(7, 1, 3, 7, 10, false)
		go playGames(7, 2, 4, 8, 10, false)
		go playGames(5, 1, 4, 5, 10, false)
		go playGames(7, 1, 2, 5, 10, false)
		go playGames(7, 3, 3, 6, 10, false)
		time.Sleep(80000000000)
		go playGames(5, 1, 2, 4, 10, false)
		go playGames(5, 1, 4, 4, 10, false)
		go playGames(5, 1, 3, 5, 10, false)
		go playGames(6, 1, 5, 6, 10, false)
		go playGames(7, 3, 2, 5, 10, false)
		time.Sleep(30000000000)
	}
}

func Benchmark_strategy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		//strategies := [][]string{}
		for d := 4; d <= 7; d++ {
			for b := 3; b <= 8; b++ {
				for y := 2; y <= 6; y++ {
					for h := 1; h <= 3; h++ {
						//avrPoint, avrRound := playGames(d, h, y, b, 100, true)
						go playGames(d, h, y, b, 26, true)
						//strategy := []string{
						//	strconv.FormatFloat(avrPoint, 'f', 3, 64),
						//	strconv.FormatFloat(avrRound, 'f', 3, 64),
						//	"d" + strconv.Itoa(d) + " h" + strconv.Itoa(h) + " y" + strconv.Itoa(y) + " b" + strconv.Itoa(b)}
						//strategies = append(strategies, strategy)
					}
					//fmt.Println("start sleep")
					time.Sleep(40000000000)
					//fmt.Println("end sleep")
				}
				time.Sleep(20000000000)
			}
		}
		//save("strategies", strategies)
	}
}

func playGames(dK, hK, yK, bK, amount int, saveReport bool) {
	records := [][]string{}
	scores := []int{}
	rounds := []int{}
	for i := 0; i < amount; i++ {
		g := Game{
			DamageK: dK,
			HoleK:   hK,
			PostyK:  yK,
			BurnK:   bK,
		}
		roud, score := playGame(&g)
		records = append(records, []string{strconv.Itoa(roud), strconv.Itoa(score)})
		scores = append(scores, score)
		rounds = append(rounds, roud)
		//fmt.Println(roud, score)
	}
	avrPoint := average(scores)
	avrRound := average(rounds)

	if saveReport {
		filename := "d" + strconv.Itoa(dK) + "_h" + strconv.Itoa(hK) + "_y" + strconv.Itoa(yK) + "_b" + strconv.Itoa(bK) +
			"_s" + strconv.FormatFloat(avrPoint, 'f', 3, 64) +
			"_r" + strconv.FormatFloat(avrRound, 'f', 3, 64) +
			"_" + strconv.FormatInt(int64(time.Now().UTC().UnixNano()), 10)
		save(filename, records)
	}
	result := strconv.Itoa(dK) +
		"	" + strconv.Itoa(hK) +
		"	" + strconv.Itoa(yK) +
		"	" + strconv.Itoa(bK) +
		"	" + strconv.FormatFloat(avrPoint, 'f', 3, 64) +
		"	" + strconv.FormatFloat(avrRound, 'f', 3, 64)
	fmt.Println(result)

	//return avrPoint, avrRound
}

func playGame(g *Game) (int, int) {
	rand.Seed(time.Now().UTC().UnixNano())
	g.asignSettings("player_names", "player1,player2")
	g.asignSettings("your_bot", "player1")
	g.Round = 0
	g.NextPiece = pieces[rand.Intn(len(pieces))]
	g.MyPlayer.Points = 0
	position := &Position{}
	position.FieldAfter = initialField
	keepGoing := true

	for keepGoing {
		assignPieces(g)
		applyPoints(g, position)
		position.FieldAfter.Burn()
		g.MyPlayer.Field = position.FieldAfter
		addSolidLines(g)
		addGarbageLines(g)
		g.Round++

		position = g.calculateMoves()
		
		if position == nil || isRoof(g) {
			keepGoing = false
		}
	}

	return g.Round, g.MyPlayer.Points
}

func assignPieces(g *Game) {
	rand.Seed(time.Now().UTC().UnixNano())
	g.CurrentPiece = g.NextPiece
	g.NextPiece = pieces[rand.Intn(len(pieces))]
}

func applyPoints(g *Game, pos *Position) {
	if pos.Burn > 0 {
		g.MyPlayer.Combo++
		g.MyPlayer.Points += g.MyPlayer.Combo - 1
		switch pos.Burn {
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

func isRoof(g *Game) bool {
	for _, col := range g.MyPlayer.Field[g.MyPlayer.Field.Height()-1]{
		if col {
			//fmt.Println("roof", g.MyPlayer.Field.Height())
			return true
		}
	}
	return false
}

func addSolidLines(g *Game) {
	r := g.Round % 20
	if r == 0 {
		g.MyPlayer.Field = g.MyPlayer.Field[:g.MyPlayer.Field.Height()-1]
	}
}

func addGarbageLines(g *Game) {
	r := g.Round % 6
	if r == 0 {
		size := g.MyPlayer.Field.Width()
		row := make([]bool, size, size)
		for i := range row {
			row[i] = true
		}
		hole := rand.Intn(size)
		row[hole] = false
		g.MyPlayer.Field = append([][]bool{row}, [][]bool(g.MyPlayer.Field[:g.MyPlayer.Field.Height()-1])...)
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

func average(a []int) float64 {
	total := 0.0
	min := 10000000000.0
	max := 0.0
	for _, v := range a {
		vv := float64(v)
		total += vv
		if vv > max {
			max = vv
		}
		if vv < min {
			min = vv
		}
	}
	total = total - max - min

	return total / float64(len(a)-2)
}
