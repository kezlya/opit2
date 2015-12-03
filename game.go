package main

import (
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
					grid := GridFromString(cleanSource)
					g.Players[i].Field = grid.ToField()
					break
				}
			}
		}
	}
}

func (g *Game) calculateMoves() *Piece {
	mf := g.MyPlayer.Field
	positions := mf.ValidPosition(g.CurrentPiece)
	hBlocked, hFixable := mf.FindHoles()
	countBh := len(hBlocked)
	countFh := len(hFixable)
	if len(hFixable) > 0 {
		fixes := mf.FixHoles(g.CurrentPiece, hFixable)
		positions = append(positions, fixes...)
	}

	for i, p := range positions {
		if p.Score.Burn > 0 {
			p.FieldAfter.Burn()
		}
		pp := p.FieldAfter.Picks

		nPositions := p.FieldAfter.ValidPosition(g.NextPiece)
		nhBlocked, nhFixable := p.FieldAfter.FindHoles()
		ncountBh := len(nhBlocked)
		ncountFh := len(nhFixable)
		if len(nhFixable) > 0 {
			nfixes := p.FieldAfter.FixHoles(g.NextPiece, nhFixable)
			nPositions = append(nPositions, nfixes...)
		}

		positions[i].Score.BHoles = ncountBh - countBh
		positions[i].Score.FHoles = ncountFh - countFh
		positions[i].setHighY()
		positions[i].setStep(g.MyPlayer.Field.Picks)
		positions[i].setCHoles(nhBlocked)

		for j, np := range nPositions {

			if np.Score.Burn > 0 {
				np.FieldAfter.Burn()
			}
			if ((g.Round + 1) % 20) == 0 {
				np.FieldAfter.Grid = np.FieldAfter.Grid[:np.FieldAfter.Height-1]
			}

			npp := np.FieldAfter.Picks
			nEmpty := np.FieldAfter.Height - npp.Max()
			nnhBlocked, nnhFixable := np.FieldAfter.FindHoles()
			nPositions[j].Score.BHoles = len(nnhBlocked) - ncountBh
			nPositions[j].Score.FHoles = len(nnhFixable) - ncountFh
			nPositions[j].setHighY()
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
		positions[i].setTotalScore(g.Strategy, g.MyPlayer.Field.Empty, countBh)
	}

	if len(positions) > 0 {
		OrderedBy(SCORE).Sort(positions)
		return &positions[0]
	}
	return nil
}
