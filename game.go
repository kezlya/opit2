package main

import (
	//	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	X           int
	Y           int
	Timebank    int
	TimePerMove int
	Width       int
	Height      int
	Round       int

	CurrentPieceName string
	NextPieceName    string
	CurrentPiece     Piece
	NextPiece        Piece

	Players  []Player
	MyPlayer *Player

	Strategy Strategy
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
			g.CurrentPieceName = value
		case "next_piece_type":
			g.NextPieceName = value
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
					grid := InitGrid(cleanSource)
					g.Players[i].Field = grid.ToField()
					break
				}
			}
		}
	}
}

func (g *Game) initPieces() {
	realY := g.MyPlayer.Field.Height + g.Y
	g.CurrentPiece = InitPiece(g.CurrentPieceName, g.X, realY)
	g.NextPiece = InitPiece(g.NextPieceName, g.X, realY)
}

func (g *Game) calculateMoves() *Piece {
	mf := g.MyPlayer.Field
	positions := mf.FindPositions(g.CurrentPiece)
	for i := range positions {
		nPositions := positions[i].FieldAfter.FindPositions(g.NextPiece)
		for j := range nPositions {
			//if ((g.Round + 1) % 20) == 0 {
			//	np.FieldAfter.Grid = np.FieldAfter.Grid[:np.FieldAfter.Height-1]
			//}
			nPositions[j].SetScore(g.Strategy, positions[i].FieldAfter.CountBH, positions[i].FieldAfter.CountFH, 0)

		}
		newScore := 10000000000000
		if len(nPositions) > 0 {
			OrderedBy(SCORE).Sort(nPositions)
			newScore = nPositions[0].Score.Total
			//fmt.Println(newScore)
		}
		positions[i].SetScore(g.Strategy, mf.CountBH, mf.CountFH, newScore)
		//fmt.Println(p.Score, p.Score.Burn)
	}

	if len(positions) > 0 {
		OrderedBy(SCORE).Sort(positions)
		return &positions[0]
	}
	return nil
}
