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

// game one Score: 105 Round 145
// Score: 108 Round 146
// Score: 115 Round 158
// Score: 139 Round 181
// Score: 121 Round 162
var defaultStrategy = Strategy{
	Burn:   2,
	BHoles: 7,
	FHoles: 4,
	CHoles: 1,
	HighY:  2,
	Step:   2,
}

/*
func Test_generateGames(t *testing.T) {
	for j := 1; j <= 20; j++ {

		i := 1
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Print("var g", j, " =[300]string{\"", pieces[rand.Intn(len(pieces))], "\"")
		for i < 300 {
			fmt.Print(",\"", pieces[rand.Intn(len(pieces))], "\"")
			i++
		}
		fmt.Print("}")
		fmt.Println()

	}

}
*/
/*
func Test_generateGarbageRows(t *testing.T) {
	size := 10
	for j := 1; j <= 20; j++ {
		i := 1
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Print("var gr", j, " =[300]int{", rand.Intn(size))
		for i < 300 {
			fmt.Print(",", rand.Intn(size))
			i++
		}
		fmt.Print("}")
		fmt.Println()
	}
}
*/

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
	fmt.Println()
	for n := 0; n < b.N; n++ {
		round, score := playGame(&Game{Strategy: defaultStrategy}, g4, gr4, false)
		fmt.Println("Score:", score, "Round", round)
	}
}

func Benchmark_one(b *testing.B) {
	for n := 0; n < b.N; n++ {
		playGame(&Game{Strategy: defaultStrategy}, g4, gr4, true)
	}
}

func Benchmark_strategy(banch *testing.B) {
	for n := 0; n < banch.N; n++ {
		for b := 1; b <= 5; b++ {
			for bh := 5; bh <= 15; bh++ {
				for fh := 1; fh <= 5; fh++ {
					for ch := 2; ch <= 3; ch++ {
						for hy := 1; hy <= 3; hy++ {
							for s := 1; s <= 3; s++ {
								st := Strategy{Burn: b, BHoles: bh, FHoles: fh, CHoles: 1, HighY: hy, Step: s}
								go playGames(st, "")
							}
							//fmt.Println("start sleep")
							time.Sleep(50000000000)
							//fmt.Println("end sleep")
						}
						time.Sleep(10000000000)
					}
					time.Sleep(10000000000)
				}
				time.Sleep(10000000000)
			}
		}
	}
}

func Benchmark_investigate(banch *testing.B) {
	for n := 0; n < banch.N; n++ {
		fmt.Println()
		strategy := Strategy{
			Burn:   2,
			BHoles: 7,
			FHoles: 4,
			CHoles: 1,
			HighY:  1,
			Step:   2,
		}
		go playGames(strategy, "investigation_")
		time.Sleep(45000000000)
	}
}

func playGame(g *Game, input [300]string, garbage [300]int, visual bool) (int, int) {
	g.asignSettings("player_names", "player1,player2")
	g.asignSettings("your_bot", "player1")
	g.Round = 0
	g.MyPlayer.Points = 0
	g.MyPlayer.Field = initialField
	g.MyPlayer.Picks = initialField.Picks()
	position := &Piece{}
	position.FieldAfter = initialField
	assignPieces(g, input[0])
	keepGoing := true

	i := 0
	for keepGoing {
		applyPoints(g, position)
		position.FieldAfter.Burn()
		g.MyPlayer.Field = position.FieldAfter
		addSolidLines(g)
		addGarbageLines(g, garbage)
		g.MyPlayer.Picks = g.MyPlayer.Field.Picks()
		g.MyPlayer.Empty = g.MyPlayer.Field.Height() - g.MyPlayer.Picks.Max()
		assignPieces(g, input[i])
		g.Round++

		if visual {
			//fmt.Println("D", position.Damage, "S", position.Score)
			fmt.Println(g.CurrentPiece.Name, "sore:", g.MyPlayer.Points, "round:", g.Round, "combo:", g.MyPlayer.Combo)
			fmt.Printf("%+v %+v\n", position.Score, position.Name)
			PrintVisual(g.MyPlayer.Field)
			time.Sleep(1000000000)
		}

		position = g.calculateMoves()

		if position == nil ||
			g.MyPlayer.Field[g.MyPlayer.Field.Height()-1][3] ||
			g.MyPlayer.Field[g.MyPlayer.Field.Height()-1][4] ||
			g.MyPlayer.Field[g.MyPlayer.Field.Height()-1][5] ||
			g.MyPlayer.Field[g.MyPlayer.Field.Height()-1][6] {
			keepGoing = false
		}
		i++
	}

	return g.Round, g.MyPlayer.Points
}

func playGames(st Strategy, filename string) {
	r1, s1 := playGame(&Game{Strategy: st}, g1, gr1, false)
	r2, s2 := playGame(&Game{Strategy: st}, g2, gr2, false)
	r3, s3 := playGame(&Game{Strategy: st}, g3, gr3, false)
	r4, s4 := playGame(&Game{Strategy: st}, g4, gr4, false)
	r5, s5 := playGame(&Game{Strategy: st}, g5, gr5, false)
	r6, s6 := playGame(&Game{Strategy: st}, g6, gr6, false)
	r7, s7 := playGame(&Game{Strategy: st}, g7, gr7, false)
	r8, s8 := playGame(&Game{Strategy: st}, g8, gr8, false)
	r9, s9 := playGame(&Game{Strategy: st}, g9, gr9, false)
	r10, s10 := playGame(&Game{Strategy: st}, g10, gr10, false)
	r11, s11 := playGame(&Game{Strategy: st}, g11, gr11, false)
	r12, s12 := playGame(&Game{Strategy: st}, g12, gr12, false)
	r13, s13 := playGame(&Game{Strategy: st}, g13, gr13, false)
	r14, s14 := playGame(&Game{Strategy: st}, g14, gr14, false)
	r15, s15 := playGame(&Game{Strategy: st}, g15, gr15, false)
	r16, s16 := playGame(&Game{Strategy: st}, g16, gr16, false)
	r17, s17 := playGame(&Game{Strategy: st}, g17, gr17, false)
	r18, s18 := playGame(&Game{Strategy: st}, g18, gr18, false)
	r19, s19 := playGame(&Game{Strategy: st}, g19, gr19, false)
	r20, s20 := playGame(&Game{Strategy: st}, g20, gr20, false)

	strategyName := "b" + strconv.Itoa(st.Burn) + " bh" + strconv.Itoa(st.BHoles) + " fh" + strconv.Itoa(st.FHoles) + " ch" + strconv.Itoa(st.CHoles) + " y" + strconv.Itoa(st.HighY) + " s" + strconv.Itoa(st.Step)

	scores := []string{strategyName, strconv.Itoa(s1), strconv.Itoa(s2), strconv.Itoa(s3), strconv.Itoa(s4), strconv.Itoa(s5), strconv.Itoa(s6), strconv.Itoa(s7), strconv.Itoa(s8), strconv.Itoa(s9), strconv.Itoa(s10), strconv.Itoa(s11), strconv.Itoa(s12), strconv.Itoa(s13), strconv.Itoa(s14), strconv.Itoa(s15), strconv.Itoa(s16), strconv.Itoa(s17), strconv.Itoa(s18), strconv.Itoa(s19), strconv.Itoa(s20)}
	rounds := []string{strategyName, strconv.Itoa(r1), strconv.Itoa(r2), strconv.Itoa(r3), strconv.Itoa(r4), strconv.Itoa(r5), strconv.Itoa(r6), strconv.Itoa(r7), strconv.Itoa(r8), strconv.Itoa(r9), strconv.Itoa(r10), strconv.Itoa(r11), strconv.Itoa(r12), strconv.Itoa(r13), strconv.Itoa(r14), strconv.Itoa(r15), strconv.Itoa(r16), strconv.Itoa(r17), strconv.Itoa(r18), strconv.Itoa(r19), strconv.Itoa(r20)}

	fmt.Println("scores:", scores)
	fmt.Println("rounds:", rounds)

	save(filename+"score", scores)
	save(filename+"round", rounds)
}

func assignPieces(g *Game, piece string) {
	g.CurrentPiece = g.NextPiece
	x := 3
	if piece == "O" {
		x = 4
	}
	g.NextPiece = Piece{Name: piece, Rotation: 0}
	g.NextPiece.InitSpace(Cell{x, g.MyPlayer.Field.Height() - 1})
}

func applyPoints(g *Game, p *Piece) {
	if g.Round > 1 {
		points := p.getPoints(g.MyPlayer.Combo)
		if points > 0 {
			g.MyPlayer.Combo++
		} else {
			g.MyPlayer.Combo = 0
		}
		g.MyPlayer.Points += points
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
	if r == 0 && g.Round != 0 {
		g.MyPlayer.Field = g.MyPlayer.Field[:g.MyPlayer.Field.Height()-1]
		g.Height = g.Height - 1
	}
}

func addGarbageLines(g *Game, garbage [300]int) {
	r := g.Round % 5
	if r == 0 && g.Round != 0 {
		size := g.MyPlayer.Field.Width()
		row := make([]bool, size)
		for i := range row {
			row[i] = true
		}
		row[garbage[g.Round/5]] = false
		g.MyPlayer.Field = append([][]bool{row}, [][]bool(g.MyPlayer.Field[:g.MyPlayer.Field.Height()-1])...)
	}
}

func save(fileName string, record []string) {
	csvfile, err := os.OpenFile("output/"+fileName+".csv", os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	err = writer.Write(record)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	writer.Flush()
}

func statistic(a []int) (int, int, int) {
	if len(a) < 4 {
		return a[0], a[0], a[0]
	}
	sort.Ints(a)
	var total int
	for i := 1; i < len(a)-1; i++ {
		total += a[i]
	}
	avr := total/len(a) - 2
	return avr, a[1], a[len(a)-2]
}
