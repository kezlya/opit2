package main

import (
	//"fmt"
	"strconv"
	"strings"
)

type Game struct {
	X            int
	Y            int
	Timebank     int
	TimePerMove  int
	Width        int
	Height       int
	Round        int
	Players      []Player
	MyPlayer     *Player
	CurrentPiece Piece
	NextPiece    Piece
	Strategy     Strategy
}

type Player struct {
	Name   string
	Field  Field
	Picks  Picks
	Empty  int
	Points int
	Combo  int
}

type Strategy struct {
	Burn   int
	Step   int
	BHoles int
	FHoles int
	CHoles int
	HighY  int
}

func (g *Game) asignSettings(action, value string) {
	switch action {
	case "timebank":
		g.Timebank, _ = strconv.Atoi(value)
	case "time_per_move":
		g.TimePerMove, _ = strconv.Atoi(value)
	case "player_names":
		names := strings.Split(value, ",")
		g.Players = make([]Player, len(names))
		for i, name := range names {
			g.Players[i].Name = name
		}
	case "your_bot":
		for i, p := range g.Players {
			if p.Name == value {
				g.MyPlayer = &g.Players[i]
				break
			}
		}
	case "field_width":
		g.Width, _ = strconv.Atoi(value)
	case "field_height":
		g.Height, _ = strconv.Atoi(value)
	}
}

func (g *Game) asignUpdates(who, action, value string) {
	switch who {
	case "game":
		switch action {
		case "round":
			g.Round, _ = strconv.Atoi(value)
		case "this_piece_type":
			g.CurrentPiece.Name = value
			g.CurrentPiece.Rotation = 0
		case "next_piece_type":
			g.NextPiece.Name = value
			g.NextPiece.Rotation = 0
		case "this_piece_position":
			cor := strings.Split(value, ",")
			g.X, _ = strconv.Atoi(cor[0])
			g.Y, _ = strconv.Atoi(cor[1])
		}
	default:
		switch action {
		case "row_points":
			for i, p := range g.Players {
				if p.Name == who {
					g.Players[i].Points, _ = strconv.Atoi(value)
					break
				}
			}
		case "combo":
			for i, p := range g.Players {
				if p.Name == who {
					g.Players[i].Combo, _ = strconv.Atoi(value)
					break
				}
			}
		case "field":
			cleanSource := strings.Replace(value, ";3,3,3,3,3,3,3,3,3,3", "", g.Height)
			for i, p := range g.Players {
				if p.Name == who {
					var pf Field
					pf = pf.init(cleanSource)
					pfp := pf.Picks()

					g.Players[i].Field = pf
					g.Players[i].Picks = pfp
					g.Players[i].Empty = pf.Height() - pfp.Max()
					break
				}
			}
		}
	}
}

func (g *Game) calculateMoves() *Piece {
	trim := g.trimStrategy()
	//dsr_x := g.MyPlayer.Field.IsDSR()

	mf := g.MyPlayer.Field
	mfp := mf.Picks()

	positions := mf.ValidPosition(g.CurrentPiece, mfp, trim)
	hBlocked, hFixable := mf.FindHoles()
	countBh := len(hBlocked)
	countFh := len(hFixable)
	if len(hFixable) > 0 {
		fixes := mf.FixHoles(g.CurrentPiece, hFixable, mfp.Max())
		positions = append(positions, fixes...)
	}

	for i, p := range positions {
		/*if !p.isDSRfriendly(g.MyPlayer.Field.Height(), g.MyPlayer.Empty) {
			fmt.Println(p.Name, p.Space, "skiped")
			continue
		}*/

		if p.Score.Burn > 0 {
			p.FieldAfter.Burn()
		}
		pp := p.FieldAfter.Picks()

		nPositions := p.FieldAfter.ValidPosition(g.NextPiece, pp, trim)
		//ndsr_x := p.FieldAfter.IsDSR()
		nhBlocked, nhFixable := p.FieldAfter.FindHoles()
		ncountBh := len(nhBlocked)
		ncountFh := len(nhFixable)
		if len(nhFixable) > 0 {
			nfixes := p.FieldAfter.FixHoles(g.NextPiece, nhFixable, pp.Max())
			nPositions = append(nPositions, nfixes...)
		}

		positions[i].Score.BHoles = ncountBh - countBh
		positions[i].Score.FHoles = ncountFh - countFh
		positions[i].setHighY()
		//positions[i].setDSR(dsr_x, ndsr_x)
		positions[i].setStep(g.MyPlayer.Picks)
		positions[i].setCHoles(nhBlocked)

		for j, np := range nPositions {
			/*if !np.isDSRfriendly(p.FieldAfter.Height(), p.FieldAfter.Height()-pp.Max()) {
				continue
			}*/

			if np.Score.Burn > 0 {
				np.FieldAfter.Burn()
			}
			if ((g.Round + 1) % 20) == 0 {
				np.FieldAfter = np.FieldAfter[:np.FieldAfter.Height()-1]
			}

			npp := np.FieldAfter.Picks()
			nEmpty := np.FieldAfter.Height() - npp.Max()

			nnhBlocked, nnhFixable := np.FieldAfter.FindHoles()
			//nndsr_x := np.FieldAfter.IsDSR()

			nPositions[j].Score.BHoles = len(nnhBlocked) - ncountBh
			nPositions[j].Score.FHoles = len(nnhFixable) - ncountFh
			nPositions[j].setHighY()
			//nPositions[j].setDSR(ndsr_x, nndsr_x)
			nPositions[j].setStep(pp)
			nPositions[j].setCHoles(nnhBlocked)
			nPositions[j].setTotalScore(g.Strategy, nEmpty, ncountBh)
		}

		if len(nPositions) > 0 {
			OrderedBy(SCORE).Sort(nPositions)
			positions[i].Score.NScore = nPositions[0].Score.Total
		} else {
			positions[i].Score.NScore = 10000000000000
		}
		positions[i].setTotalScore(g.Strategy, g.MyPlayer.Empty, countBh)
		//fmt.Printf("%+v\n", p.sco)
	}

	if len(positions) > 0 {

		OrderedBy(SCORE).Sort(positions)
		/*for _, tempP := range positions {
			fmt.Printf("%+v\n", tempP.Score)
			fmt.Println(tempP.CurrentX, tempP.Name, tempP.Rotation)
			fmt.Println("")
		}*/

		return &positions[0]
	}
	return nil
}

func (g *Game) trimStrategy() int {
	trim := 0
	//trim 1 doesn't work
	/*if g.MyPlayer.Empty > 9 && g.Round < 100 {
		trim = 1
		if g.CurrentPiece.Name == "I" && (g.MyPlayer.Picks[len(g.MyPlayer.Picks)-2]-g.MyPlayer.Picks[len(g.MyPlayer.Picks)-1]) > 2 {
			trim = 0
		}
		if g.CurrentPiece.Name == "L" &&
			(g.MyPlayer.Picks[len(g.MyPlayer.Picks)-3]-g.MyPlayer.Picks[len(g.MyPlayer.Picks)-2]) > 0 &&
			(g.MyPlayer.Picks[len(g.MyPlayer.Picks)-2]-g.MyPlayer.Picks[len(g.MyPlayer.Picks)-1]) > 1 {
			trim = 0
		}
	}*/
	/*if g.MyPlayer.Empty > 9 && g.Round < 50 {
		trim = 3
	}*/

	return trim
}
