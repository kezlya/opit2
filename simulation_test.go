package main

import (
	//"encoding/csv"
	"fmt"
	"github.com/agonopol/gosplat"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	//"math/rand"
	"os"
	"sort"
	"strconv"
	//"strings"
	"testing"
	"time"
)

/*
func Test_generateGames(t *testing.T) {
	for j := 1; j <= 26; j++ {

		i := 1
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Print("var g", j, " =[400]string{\"", pieces[rand.Intn(len(pieces))], "\"")
		for i < 400 {
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
	for j := 1; j <= 26; j++ {
		i := 1
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Print("var gr", j, " =[60]int{", rand.Intn(size))
		for i < 60 {
			fmt.Print(",", rand.Intn(size))
			i++
		}
		fmt.Print("}")
		fmt.Println()
	}
}
*/

func Benchmark_one(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buff := 1
		ch_round := make(chan int, buff)
		ch_score := make(chan int, buff)
		go playGame(ch_round, ch_score, &Game{Strategy: strategy}, &g10, &gr7, true)
		scores := make([]int, buff)
		rounds := make([]int, buff)
		for k := 0; k < buff; k++ {
			scores[k] = <-ch_score
			rounds[k] = <-ch_round
		}

		fmt.Println("scores:", scores)
		fmt.Println("rounds:", rounds)
	}
}

func Benchmark_many(banch *testing.B) {
	for n := 0; n < banch.N; n++ {
		// strategy taking from main file current Strategy for the bot
		strategyName := strategy.name()
		fmt.Println()
		fmt.Println("================================")
		fmt.Println(strategyName)

		scores, rounds := playGames(strategy)
		Linechart(&oldScores, scores, &oldRounds, rounds, strategyName)
		fmt.Println("done")
	}
}

func Benchmark_strategy(banch *testing.B) {
	for n := 0; n < banch.N; n++ {
		for b := 1; b <= 2; b++ {
			for bh := 10; bh <= 12; bh++ {
				for fh := 8; fh <= 10; fh++ {
					for ch := 2; ch <= 3; ch++ {
						for hy := 2; hy <= 3; hy++ {
							for s := 2; s <= 3; s++ {
								st := Strategy{Burn: b, BHoles: bh, FHoles: fh, CHoles: ch, HighY: hy, Step: s}
								fmt.Println()
								fmt.Println("================================")
								strategyName := st.name()
								fmt.Println(strategyName)
								scores, rounds := playGames(st)
								if CheckIfStrategyIsBetter(&oldScores, scores, &oldRounds, rounds) {
									Linechart(&oldScores, scores, &oldRounds, rounds, strategyName)
								} else {
									fmt.Println("Bad")
								}
							}
						}
					}
				}
			}
		}
	}
}

func playGame(ch_round chan int, ch_score chan int, g *Game, input *[400]string, garbage *[60]int, visual bool) {
	g.asignSettings("player_names", "player1,player2")
	g.asignSettings("your_bot", "player1")
	g.asignSettings("field_width", "10")
	g.asignSettings("field_height", "20")
	field := EmptyGrig10x20.ToField()
	position := &Piece{FieldAfter: &field}
	round := 1
	keepGoing := true
	for keepGoing {
		// setup new round
		g.asignUpdates("game", "round", strconv.Itoa(round))
		g.asignUpdates("game", "this_piece_type", input[round-1])
		g.asignUpdates("game", "next_piece_type", input[round])
		g.asignUpdates("game", "this_piece_position", "3,-1")
		g.MyPlayer.Field = *position.FieldAfter
		g.initPieces()

		// play round
		position = g.calculateMoves()
		g.MyPlayer.Points += position.getPoints()
		if visual {
			fmt.Println()
			fmt.Println("===============================================================")
			fmt.Println()
			g.MyPlayer.Field.Grid.visual()
			fmt.Println(g.CurrentPiece.Name, "sore:", g.MyPlayer.Points, "round:", g.Round, "combo:", g.MyPlayer.Combo)
			fmt.Printf("%+v\n", position.Score)
			position.FieldAfter.Grid.visual()
			time.Sleep(1000000000)
		}

		// check if the game is over
		if position == nil ||
			g.MyPlayer.Field.Grid[g.MyPlayer.Field.Height-1][3] ||
			g.MyPlayer.Field.Grid[g.MyPlayer.Field.Height-1][4] ||
			g.MyPlayer.Field.Grid[g.MyPlayer.Field.Height-1][5] ||
			g.MyPlayer.Field.Grid[g.MyPlayer.Field.Height-1][6] {
			keepGoing = false
			break
		}
		if addSolidLines(g, position) {
			keepGoing = false
			break
		}
		if addGarbageLines(g, position, garbage) {
			keepGoing = false
			break
		}
		round++
	}
	ch_round <- g.Round
	ch_score <- g.MyPlayer.Points
}

func playGames(st Strategy) (*[]int, *[]int) {
	buff := 26
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

	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g21, &gr21, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g22, &gr22, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g23, &gr23, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g24, &gr24, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g25, &gr25, false)
	go playGame(ch_round, ch_score, &Game{Strategy: st}, &g26, &gr26, false)

	scores := make([]int, buff)
	rounds := make([]int, buff)
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

func addSolidLines(g *Game, p *Piece) bool {
	r := g.Round % 20
	if r == 0 && g.Round != 0 {
		if g.MyPlayer.Field.Empty == 0 {
			return true
		}
		newGrid := p.FieldAfter.Grid[:p.FieldAfter.Height-1]
		newField := newGrid.ToField()
		p.FieldAfter = &newField
	}
	return false
}

func addGarbageLines(g *Game, p *Piece, garbage *[60]int) bool {
	speed := 7
	r := g.Round % speed
	if r == 0 && g.Round != 0 {
		if g.MyPlayer.Field.Empty == 0 {
			return true
		}
		size := g.Width
		row := make([]bool, size)
		for i := range row {
			row[i] = true
		}
		row[garbage[g.Round/speed]] = false
		row[garbage[len(garbage)-g.Round/speed]] = false
		newGrid := Grid(append([][]bool{row}, [][]bool(p.FieldAfter.Grid[:p.FieldAfter.Height-1])...))
		newField := newGrid.ToField()
		p.FieldAfter = &newField
	}
	return false
}

/*
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
*/
func Linechart(scores, new_scores, rounds, new_rounds *[]int, strategy string) {
	cScores := gosplat.NewChart()
	cRounds := gosplat.NewChart()
	for i := 0; i < len(*new_scores); i++ {
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
	html, err := f.Html()
	if err != nil {
		panic(err)
	}
	p, err := ioutil.TempFile("", "go2splat.preview.")
	if err != nil {
		panic(err)
	}
	_, err = p.Write(html.Bytes())
	if err != nil {
		panic(err)
	}
	p.Close()
	name := fmt.Sprintf("%s.html", p.Name())
	os.Rename(p.Name(), name)
	open.Run(name)
	fmt.Println(name)
}

func CheckIfStrategyIsBetter(scores, new_scores, rounds, new_rounds *[]int) bool {
	counterS := 0
	counterR := 0
	half := len(*new_scores) / 2
	for i := 0; i < len(*new_scores); i++ {
		if (*scores)[i]-(*new_scores)[i] <= 0 {
			counterS++
		}
		if (*rounds)[i]-(*new_rounds)[i] <= 0 {
			counterR++
		}
	}
	fmt.Println("Better Scores:", counterS)
	fmt.Println("Better Rounds:", counterR)
	return (counterS > half || counterR > half)
}

func (s *Strategy) name() string {
	return "b" + strconv.Itoa(s.Burn) +
		" bh" + strconv.Itoa(s.BHoles) +
		" fh" + strconv.Itoa(s.FHoles) +
		" ch" + strconv.Itoa(s.CHoles) +
		" y" + strconv.Itoa(s.HighY) +
		" s" + strconv.Itoa(s.Step)
}
