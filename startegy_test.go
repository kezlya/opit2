package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
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

func Benchmark_best_strategy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go playGames(Strategy{Burn: 3, Step: 1, BHoles: 7, FHoles: 3, HighY: 2}, 22, false, false)
		go playGames(Strategy{Burn: 5, Step: 2, BHoles: 7, FHoles: 4, HighY: 1}, 22, false, false)
		go playGames(Strategy{Burn: 1, Step: 1, BHoles: 6, FHoles: 4, HighY: 2}, 22, false, false)
		go playGames(Strategy{Burn: 2, Step: 2, BHoles: 5, FHoles: 3, HighY: 2}, 22, false, false)
		time.Sleep(50000000000)
		go playGames(Strategy{Burn: 4, Step: 1, BHoles: 6, FHoles: 4, HighY: 3}, 22, false, false)
		go playGames(Strategy{Burn: 1, Step: 1, BHoles: 4, FHoles: 3, HighY: 1}, 22, false, false)
		go playGames(Strategy{Burn: 2, Step: 1, BHoles: 4, FHoles: 3, HighY: 1}, 22, false, false)
		go playGames(Strategy{Burn: 1, Step: 2, BHoles: 6, FHoles: 4, HighY: 1}, 22, false, false)
		time.Sleep(50000000000)
		go playGames(Strategy{Burn: 3, Step: 2, BHoles: 6, FHoles: 3, HighY: 1}, 22, false, false)
		go playGames(Strategy{Burn: 3, Step: 2, BHoles: 7, FHoles: 4, HighY: 3}, 22, false, false)
		go playGames(Strategy{Burn: 5, Step: 1, BHoles: 4, FHoles: 2, HighY: 1}, 22, false, false)
		go playGames(Strategy{Burn: 1, Step: 1, BHoles: 6, FHoles: 4, HighY: 3}, 22, false, false)
		time.Sleep(50000000000)
		go playGames(Strategy{Burn: 2, Step: 1, BHoles: 6, FHoles: 4, HighY: 1}, 22, false, false)
		go playGames(Strategy{Burn: 2, Step: 2, BHoles: 7, FHoles: 2, HighY: 1}, 22, false, false)
		go playGames(Strategy{Burn: 3, Step: 2, BHoles: 7, FHoles: 4, HighY: 1}, 22, false, false)
		go playGames(Strategy{Burn: 4, Step: 2, BHoles: 7, FHoles: 3, HighY: 4}, 22, false, false)
		time.Sleep(50000000000)
		go playGames(Strategy{Burn: 1, Step: 1, BHoles: 5, FHoles: 2, HighY: 2}, 22, false, false)
		go playGames(Strategy{Burn: 1, Step: 2, BHoles: 7, FHoles: 4, HighY: 2}, 22, false, false)
		go playGames(Strategy{Burn: 2, Step: 2, BHoles: 7, FHoles: 4, HighY: 2}, 22, false, false)
		go playGames(Strategy{Burn: 5, Step: 2, BHoles: 4, FHoles: 4, HighY: 1}, 22, false, false)
		time.Sleep(50000000000)
		go playGames(Strategy{Burn: 4, Step: 1, BHoles: 6, FHoles: 3, HighY: 1}, 22, false, false)
		go playGames(Strategy{Burn: 3, Step: 1, BHoles: 7, FHoles: 2, HighY: 4}, 22, false, false)
		go playGames(Strategy{Burn: 3, Step: 2, BHoles: 7, FHoles: 2, HighY: 4}, 22, false, false)
		go playGames(Strategy{Burn: 4, Step: 2, BHoles: 7, FHoles: 4, HighY: 2}, 22, false, false)
		time.Sleep(50000000000)
		go playGames(Strategy{Burn: 5, Step: 2, BHoles: 7, FHoles: 3, HighY: 2}, 22, false, false)
		go playGames(Strategy{Burn: 2, Step: 2, BHoles: 5, FHoles: 2, HighY: 1}, 22, false, false)
		go playGames(Strategy{Burn: 2, Step: 1, BHoles: 5, FHoles: 4, HighY: 2}, 22, false, false)
		go playGames(Strategy{Burn: 4, Step: 1, BHoles: 5, FHoles: 2, HighY: 1}, 22, false, false)
		time.Sleep(50000000000)
		go playGames(Strategy{Burn: 1, Step: 2, BHoles: 7, FHoles: 4, HighY: 3}, 22, false, false)
		go playGames(Strategy{Burn: 3, Step: 2, BHoles: 7, FHoles: 3, HighY: 1}, 22, false, false)
		go playGames(Strategy{Burn: 4, Step: 1, BHoles: 7, FHoles: 2, HighY: 3}, 22, false, false)
		time.Sleep(50000000000)
		//fmt.Println("end sleep")
	}
}

func Benchmark_strategy(banch *testing.B) {
	fmt.Println("")
	fmt.Println("Burn	BHoles	FHoles	HighY	Step	Score	minS	maxS	Round	minR	maxR")
	for n := 0; n < banch.N; n++ {
		for b := 3; b <= 6; b++ {
			for bh := 6; bh <= 9; bh++ {
				for fh := 2; fh <= 5; fh++ {
					for hy := 1; hy <= 2; hy++ {
						for s := 1; s <= 2; s++ {
							st := Strategy{Burn: b, BHoles: bh, FHoles: fh, HighY: hy, Step: s}
							go playGames(st, 22, false, false)
						}
					}
					//fmt.Println("start sleep")
					time.Sleep(60000000000)
					//fmt.Println("end sleep")
				}
				time.Sleep(60000000000)
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
	avrScore, minScore, maxScore := statistic(scores)
	avrRound, minRound, maxRound := statistic(rounds)

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
	fmt.Print(avrScore)
	fmt.Print("	")
	fmt.Print(minScore)
	fmt.Print("	")
	fmt.Print(maxScore)
	fmt.Print("	")
	fmt.Print(avrRound)
	fmt.Print("	")
	fmt.Print(minRound)
	fmt.Print("	")
	fmt.Print(maxRound)

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

func statistic(a []int) (int, int, int) {
	sort.Ints(a)
	var total int
	for i := 1; i < len(a)-1; i++ {
		total += a[i]
	}
	avr := total/len(a) - 2
	return avr, a[1], a[len(a)-2]
}
