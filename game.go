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

	picks := g.MyPlayer.Field.Picks()

	positions := g.MyPlayer.Field.ValidPosition(g.CurrentPiece, picks)

	_, hFixable := g.MyPlayer.Field.FindHoles(picks)
	if len(hFixable) > 0 {
		fixes := g.MyPlayer.Field.FixHoles(g.CurrentPiece, hFixable)
		positions = append(positions, fixes...)
	}

	for i, p := range positions {

		if p.Score.Burn > 0 { ///I don't have Burn yet
			p.FieldAfter.Burn()
		}
		nPicks := p.FieldAfter.Picks()
		np := p.FieldAfter.ValidPosition(g.NextPiece, nPicks)
		_, nhFixable := p.FieldAfter.FindHoles(nPicks)
		p.Score = p.score()
		// damage >3 discard

		if len(nhFixable) > 0 {
			nfixes := p.FieldAfter.FixHoles(g.NextPiece, nhFixable)
			np = append(np, nfixes...)
		}

		if len(np) > 0 {
			OrderedBy(SCORE).Sort(np)

			minNextScore := np[0].Score.Total
			positions[i].Score.NScore = minNextScore
		} else {
			positions[i].Score.NScore += 10000000000000 //maybe romove current piece
		}
	}
	//fmt.Println("classic", len(positions))
	if len(positions) > 0 {
		OrderedBy(SCORE).Sort(positions)
		return &positions[0]
	}
	return nil
}

func (g *Game) isSafe() bool {
	picks := g.MyPlayer.Field.Picks()
	y := picks.Max()
	rowsLeft := g.MyPlayer.Field.Height() - y
	if rowsLeft > 10 {
		return true
	}
	return false
}
