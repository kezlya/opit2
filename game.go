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
	st := g.Strategy
	if g.MyPlayer.Empty < 5 {
		st.Step = g.Strategy.Step + 3
	}
	/*if g.MyPlayer.Empty < 3 {
		st.HighY = g.Strategy.HighY + (5 - g.MyPlayer.Empty)
	}*/

	trim := 0
	//trim 1 doesn't work
	/*if g.MyPlayer.Empty > 15 {
		trim = 1
		if g.CurrentPiece.Name == "I" && (g.MyPlayer.Picks[len(g.MyPlayer.Picks)-2]-g.MyPlayer.Picks[len(g.MyPlayer.Picks)-1]) > 3 {
			trim = 0
		}
		if g.CurrentPiece.Name == "L" &&
			(g.MyPlayer.Picks[len(g.MyPlayer.Picks)-3]-g.MyPlayer.Picks[len(g.MyPlayer.Picks)-2]) > 0 &&
			(g.MyPlayer.Picks[len(g.MyPlayer.Picks)-2]-g.MyPlayer.Picks[len(g.MyPlayer.Picks)-1]) > 1 {
			trim = 0
		}
	}*/

	/*	if g.MyPlayer.Empty > 10 && g.MyPlayer.Combo < 3 {
		trim = 3

	}*/

	positions := g.MyPlayer.Field.ValidPosition(g.CurrentPiece, g.MyPlayer.Picks, trim)
	hBlocked, hFixable := g.MyPlayer.Field.FindHoles(g.MyPlayer.Picks)
	countBh := len(hBlocked)
	countFh := len(hFixable)
	if len(hFixable) > 0 {
		fixes, _ := g.MyPlayer.Field.FixHoles(g.CurrentPiece, hFixable, g.MyPlayer.Picks.Max())
		positions = append(positions, fixes...)
		//countBh = countBh + countNotFixed
		//countFh = countFh - countNotFixed
	}

	for i, p := range positions {
		if p.Score.Burn > 0 {
			p.FieldAfter.Burn()
		}
		pp := p.FieldAfter.Picks()
		nPositions := p.FieldAfter.ValidPosition(g.NextPiece, pp, trim)

		nhBlocked, nhFixable := p.FieldAfter.FindHoles(pp)
		ncountBh := len(nhBlocked)
		ncountFh := len(nhFixable)
		if len(nhFixable) > 0 {
			nfixes, _ := p.FieldAfter.FixHoles(g.NextPiece, nhFixable, pp.Max())
			nPositions = append(nPositions, nfixes...)
			//ncountBh = ncountBh + ncountNotFixed
			//ncountFh = ncountFh - ncountNotFixed
		}

		positions[i].Score.BHoles = ncountBh - countBh
		positions[i].Score.FHoles = ncountFh - countFh
		positions[i].setHighY()
		positions[i].setStep(g.MyPlayer.Picks)
		positions[i].setCHoles(nhBlocked)

		for j, np := range nPositions {
			ncombo := 0
			if p.Score.Burn > 0 {
				ncombo = g.MyPlayer.Combo + 1
			}
			if np.Score.Burn > 0 {
				np.FieldAfter.Burn()
			}
			npp := np.FieldAfter.Picks()
			nnhBlocked, nnhFixable := np.FieldAfter.FindHoles(npp)
			nPositions[j].Score.BHoles = len(nnhBlocked) - ncountBh
			nPositions[j].Score.FHoles = len(nnhFixable) - ncountFh
			nPositions[j].setHighY()
			nPositions[j].setStep(pp)
			nPositions[j].setCHoles(nnhBlocked)
			nPositions[j].setTotalScore(st, ncombo, g.MyPlayer.Empty)
		}

		if len(nPositions) > 0 {
			OrderedBy(SCORE).Sort(nPositions)
			positions[i].Score.NScore = nPositions[0].Score.Total
		} else {
			positions[i].Score.NScore = 10000000000000 //maybe remove current piece
		}
		positions[i].setTotalScore(st, g.MyPlayer.Combo, g.MyPlayer.Empty)
		//fmt.Printf("%+v\n", p.sco)
	}

	if len(positions) > 0 {

		OrderedBy(SCORE).Sort(positions)
		/*for _, tempP := range positions {
			fmt.Printf("%+v\n", tempP.Score)
			fmt.Println(tempP.CurrentX, tempP.Name, tempP.Rotation)
		}*/

		return &positions[0]
	}
	return nil
}
