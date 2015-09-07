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
	BurnK   int
	StepK   int
	DamageK int
	PostyK  int
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
			for _, p := range g.Players {
				if p.Name == who {
					var pf Field
					pf.init(cleanSource)
					pfp := pf.Picks()

					p.Field = pf
					p.Picks = pfp
					p.Empty = pf.Height() - pfp.Max()
					break
				}
			}
		}
	}
}

func (g *Game) calculateMoves() *Piece {
	positions := g.MyPlayer.Field.ValidPosition(g.CurrentPiece, g.MyPlayer.Picks)
	hBlocked, hFixable := g.MyPlayer.Field.FindHoles(g.MyPlayer.Picks)
	if len(hFixable) > 0 {
		fixes := g.MyPlayer.Field.FixHoles(g.CurrentPiece, hFixable)
		positions = append(positions, fixes...)
	}

	for _, p := range positions {
		if p.Score.Burn > 0 {
			p.FieldAfter.Burn()
		}
		pp := p.FieldAfter.Picks()
		nPositions := p.FieldAfter.ValidPosition(g.NextPiece, pp)
		nhBlocked, nhFixable := p.FieldAfter.FindHoles(pp)
		if len(nhFixable) > 0 {
			nfixes := p.FieldAfter.FixHoles(g.NextPiece, nhFixable)
			nPositions = append(nPositions, nfixes...)
		}

		p.Score.BHoles = len(nhBlocked) - len(hBlocked)
		p.Score.FHoles = len(nhFixable) - len(hFixable)
		p.Score.LowY = p.CurrentY
		p.setHighY()
		p.setStep()
		p.setCHoles(nhBlocked)

		// damage >3 discard

		for j, np := range nPositions {
			if np.Score.Burn > 0 {
				np.FieldAfter.Burn()
			}
			npp := np.FieldAfter.Picks()
			nnhBlocked, nnhFixable := np.FieldAfter.FindHoles(npp)

			np.Score.BHoles = len(nnhBlocked) - len(nhBlocked)
			np.Score.FHoles = len(nnhFixable) - len(nhFixable)
			np.Score.LowY = np.CurrentY
			np.setHighY()
			np.setStep()
			np.setCHoles(nnhBlocked)
			np.setTotalScore()
		}

		if len(nPositions) > 0 {
			OrderedBy(SCORE).Sort(nPositions)
			p.Score.NScore = nPositions[0].Score.Total
		} else {
			p.Score.NScore = 10000000000000 //maybe romove current piece
		}
		p.setTotalScore()
	}

	if len(positions) > 0 {
		OrderedBy(SCORE).Sort(positions)
		return &positions[0]
	}
	return nil
}
