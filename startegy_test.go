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
		game := Game{Strategy: gameSt}
		game.asignSettings("timebank", "10000")
		game.asignSettings("time_per_move", "500")
		game.asignSettings("player_names", "player1,player2")
		game.asignSettings("your_bot", "player1")
		game.asignSettings("field_width", "10")
		game.asignSettings("field_height", "20")

		rand.Seed(time.Now().UTC().UnixNano())

		row1 := make([]string, 10, 10)
		for i := 0; i < 10; i++ {
			row1[i] = strconv.Itoa(rand.Intn(3))
		}
		row2 := make([]string, 10, 10)
		for i := 0; i < 10; i++ {
			row2[i] = strconv.Itoa(rand.Intn(3))
		}
		row3 := make([]string, 10, 10)
		for i := 0; i < 10; i++ {
			row3[i] = strconv.Itoa(rand.Intn(3))
		}

		game.asignUpdates("game", "round", "4")
		game.asignUpdates("game", "this_piece_type", pieces[rand.Intn(len(pieces))])
		game.asignUpdates("game", "next_piece_type", pieces[rand.Intn(len(pieces))])
		game.asignUpdates("game", "this_piece_position", "3,-1")
		game.asignUpdates("player1", "field", "0,0,0,1,1,1,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;"+strings.Join(row1, ",")+";"+strings.Join(row2, ",")+";"+strings.Join(row3, ","))
		game.asignUpdates("player1", "row_points", "0")
		game.asignUpdates("player1", "combo", "0")
		game.calculateMoves()
	}
}
func Benchmark_fixholes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testField := Field{{false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
		piece := Piece{Name: "T", Rotation: 0}
		piece.InitSpace(Cell{X: 3, Y: 19})
		hole := Cell{X: 5, Y: 0}
		testField.FixHoles(piece, []Cell{hole}, testField.Picks().Max())
	}
}

func Benchmark_many(b *testing.B) {
	for n := 0; n < b.N; n++ {
		playGames(gameSt, 100, false, false)
	}
}

func Benchmark_one(b *testing.B) {
	for n := 0; n < b.N; n++ {
		playGames(gameSt, 1, false, true)
	}
}

/*func Benchmark_best_strategy(b *testing.B) {
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
}*/

// when testing strategy output range of all games
// when testing strategy output range of all games
// when testing strategy output range of all games
// when testing strategy output range of all games
func Benchmark_strategy(banch *testing.B) {
	fmt.Println("")
	fmt.Println("Burn	BHoles	FHoles	HighY	Step	Score	Round")
	for n := 0; n < banch.N; n++ {
		for b := 1; b <= 5; b++ {
			for bh := 1; bh <= 5; bh++ {
				for fh := 1; fh <= 5; fh++ {
					for hy := 1; hy <= 5; hy++ {
						for s := 1; s <= 5; s++ {
							st := Strategy{Burn: b, BHoles: bh, FHoles: fh, HighY: hy, Step: s}
							go playGames(st, 16, false, false)
						}
						//fmt.Println("start sleep")
						time.Sleep(60000000000)
						//fmt.Println("end sleep")
					}
					time.Sleep(10000000000)
				}
			}
		}
		//save("strategies", strategies)
	}
}

func playGames(st Strategy, amount int, saveReport bool, visual bool) {
	records := [][]string{}
	scores := []int{}
	rounds := []int{}
	for i := 0; i < amount; i++ {

		g := Game{Strategy: st}
		roud, score := playGame(&g, visual)
		records = append(records, []string{strconv.Itoa(roud), strconv.Itoa(score)})
		scores = append(scores, score)
		rounds = append(rounds, roud)
		//fmt.Println(roud, score)
	}
	avrPoint := average(scores)
	avrRound := average(rounds)

	/*
		if saveReport {
			filename := "d" + strconv.Itoa(st.DamageK) + "_s" + strconv.Itoa(st.StepK) + "_y" + strconv.Itoa(st.PostyK) + "_b" + strconv.Itoa(st.BurnK) +
				"_s" + strconv.FormatFloat(avrPoint, 'f', 3, 64) +
				"_r" + strconv.FormatFloat(avrRound, 'f', 3, 64) +
				"_" + strconv.FormatInt(int64(time.Now().UTC().UnixNano()), 10)
			save(filename, records)
		}*/
	//fmt.Printf("%+v\n", st)
	//fmt.Println("points:", avrPoint)
	//fmt.Println("rounds", avrRound)
	fmt.Println()
	fmt.Print(st.Burn)
	fmt.Print("	")
	fmt.Print(st.BHoles)
	fmt.Print("	")
	fmt.Print(st.FHoles)
	fmt.Print("	")
	fmt.Print(st.HighY)
	fmt.Print("	")
	fmt.Print(st.Step)
	fmt.Print("	")
	fmt.Print(avrPoint)
	fmt.Print("	")
	fmt.Print(avrRound)

	//return avrPoint, avrRound
}

func playGame(g *Game, visual bool) (int, int) {
	g.asignSettings("player_names", "player1,player2")
	g.asignSettings("your_bot", "player1")
	g.Round = 0
	g.MyPlayer.Points = 0
	g.MyPlayer.Field = initialField
	g.MyPlayer.Picks = initialField.Picks()
	position := &Piece{}
	position.FieldAfter = initialField
	assignPieces(g)
	keepGoing := true

	for keepGoing {
		applyPoints(g, position)
		position.FieldAfter.Burn()
		g.MyPlayer.Field = position.FieldAfter
		addSolidLines(g)
		addGarbageLines(g)
		g.MyPlayer.Picks = g.MyPlayer.Field.Picks()
		assignPieces(g)
		g.Round++

		if visual {
			//fmt.Println("D", position.Damage, "S", position.Score)
			fmt.Println(g.CurrentPiece.Name)
			if position.Moves == "" {
				fmt.Println("drop")
			} else {
				fmt.Println(position.Moves + ",drop")
			}
			PrintVisual(g.MyPlayer.Field)
			//time.Sleep(1000000000)
		}

		position = g.calculateMoves()

		if position == nil ||
			g.MyPlayer.Field[g.MyPlayer.Field.Height()-1][3] ||
			g.MyPlayer.Field[g.MyPlayer.Field.Height()-1][4] ||
			g.MyPlayer.Field[g.MyPlayer.Field.Height()-1][5] ||
			g.MyPlayer.Field[g.MyPlayer.Field.Height()-1][6] {
			keepGoing = false
		}
	}

	return g.Round, g.MyPlayer.Points
}

func assignPieces(g *Game) {
	rand.Seed(time.Now().UTC().UnixNano())
	g.CurrentPiece = g.NextPiece
	x := 3
	piece := pieces[rand.Intn(len(pieces))]
	if piece == "O" {
		x = 4
	}
	g.NextPiece = Piece{Name: piece, Rotation: 0}
	g.NextPiece.InitSpace(Cell{x, g.MyPlayer.Field.Height() - 1})
}

func applyPoints(g *Game, p *Piece) {
	if p.Score.Burn > 0 {
		g.MyPlayer.Combo++
		g.MyPlayer.Points += g.MyPlayer.Combo - 1
		switch p.Score.Burn {
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
	for _, col := range g.MyPlayer.Field[g.MyPlayer.Field.Height()-1] {
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
		g.Height = g.Height - 1
	}
}

func addGarbageLines(g *Game) {
	r := g.Round % 6
	if r == 0 && g.Round != 0 {
		size := g.MyPlayer.Field.Width()
		row := make([]bool, size)
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
