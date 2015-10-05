package main

import (
	"encoding/csv"
	"fmt"
	"github.com/agonopol/gosplat"
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

/*
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
*/
/*
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
*/
func Benchmark_investigate(banch *testing.B) {
	oldScores := []int{38, 48, 58, 63, 69, 69, 71, 79, 81, 81, 82, 90, 94, 102, 108, 110, 128, 169, 180, 180}
	oldRounds := []int{76, 81, 93, 101, 101, 107, 111, 111, 113, 116, 119, 121, 126, 143, 150, 161, 176, 201, 213, 225}
	for n := 0; n < banch.N; n++ {
		fmt.Println()
		st := Strategy{
			Burn:   2,
			BHoles: 7,
			FHoles: 4,
			CHoles: 1,
			HighY:  1,
			Step:   2,
		}
		strategyName := "b" + strconv.Itoa(st.Burn) + " bh" + strconv.Itoa(st.BHoles) + " fh" + strconv.Itoa(st.FHoles) + " ch" + strconv.Itoa(st.CHoles) + " y" + strconv.Itoa(st.HighY) + " s" + strconv.Itoa(st.Step)
		scores, rounds := playGames(st)
		//Linechart(scores, scores, rounds, rounds, strategyName)
		Linechart(&oldScores, scores, &oldRounds, rounds, strategyName)
		fmt.Println("done")
	}
}

func playGame(ch_round chan int, ch_score chan int, g *Game, input *[300]string, garbage *[300]int, visual bool) {
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
	ch_round <- g.Round
	ch_score <- g.MyPlayer.Points
}

func playGames(st Strategy) (*[]int, *[]int) {
	buff := 20
	ch_round := make(chan int, buff)
	ch_score := make(chan int, buff)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g1, &gr1, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g2, &gr2, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g3, &gr3, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g4, &gr4, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g5, &gr5, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g6, &gr6, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g7, &gr7, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g8, &gr8, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g9, &gr9, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g10, &gr10, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g11, &gr11, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g12, &gr12, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g13, &gr13, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g14, &gr14, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g15, &gr15, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g16, &gr16, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g17, &gr17, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g18, &gr18, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g19, &gr19, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g20, &gr20, false)

	scores := make([]int, 20)
	rounds := make([]int, 20)
	for k := 0; k < buff; k++ {
		scores[k] = <-ch_score
		rounds[k] = <-ch_round
	}
	sort.Ints(scores)
	sort.Ints(rounds)
	fmt.Println("scores:", scores)
	fmt.Println("rounds:", rounds)
	return &scores, &rounds
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

func addGarbageLines(g *Game, garbage *[300]int) {
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

func Linechart(scores, new_scores, rounds, new_rounds *[]int, strategy string) {
	cScores := gosplat.NewChart()
	cRounds := gosplat.NewChart()
	for i := 0; i < len(*scores); i++ {
		cScores.Append(map[string]interface{}{"game": i, "old": (*scores)[i], "new": (*new_scores)[i]})
		cRounds.Append(map[string]interface{}{"game": i, "old": (*rounds)[i], "new": (*new_rounds)[i]})
	}

	f := gosplat.NewFrame(strategy)
	f.Append("Score", cScores.Linechart(map[string]interface{}{
		"series": map[string]interface{}{
			"0": map[string]interface{}{
				"color": "red"},
			"1": map[string]interface{}{
				"color": "black"}}}))
	f.Append("Round", cRounds.Linechart(map[string]interface{}{
		"series": map[string]interface{}{
			"0": map[string]interface{}{
				"color": "red"},
			"1": map[string]interface{}{
				"color": "black"}}}))

	//Preview generates a tmp html file and opens it with the default browser
	err := f.Preview()
	if err != nil {
		panic(err)
	}

	//Html returns bytes.Buffer of the html
	//buffer, err := f.Html()
	//if err != nil {
	//	panic(err)
	//}
	//println(buffer.String())
}
