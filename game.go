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
					var empty Field
					g.Players[i].Field = empty.init(cleanSource)
					g.Height = g.Players[i].Field.Height()
					break
				}
			}
		}
	}
}

func (g *Game) calculateMoves() *Position {
	positions := g.MyPlayer.Field.Positions(g.CurrentPiece, g.Strategy)
	if g.MyPlayer.Combo >= 2 {
		burned := g.keepBurning(positions)
		if burned != nil {
			return burned
		}
	}

	return g.clasic(positions)
}

func (g *Game) keepBurning(positions []Position) *Position {
	var burned []Position
	var burnStrategy = Strategy{
		BurnK:   g.Strategy.BurnK + 100,
		StepK:   g.Strategy.StepK,
		DamageK: g.Strategy.DamageK,
		PostyK:  g.Strategy.PostyK,
	}
	for i, position := range positions {
		if position.Burn > 0 {
			position.FieldAfter.Burn()
			nextPositions := position.FieldAfter.Positions(g.NextPiece, burnStrategy)
			if len(nextPositions) > 0 {
				OrderedBy(SCORE).Sort(nextPositions)
				minNextScore := nextPositions[0].Score
				positions[i].Score += minNextScore
			} else {
				positions[i].Score += 10000000000000
			}
			burned = append(burned, positions[i])
		}
	}

	if len(burned) > 0 {
		OrderedBy(SCORE).Sort(burned)
		return &burned[0]
	}

	return nil
}

func (g *Game) clasic(positions []Position) *Position {
	for i, position := range positions {
		if position.Burn > 0 {
			position.FieldAfter.Burn()
		}
		nextPositions := position.FieldAfter.Positions(g.NextPiece, g.Strategy)
		if len(nextPositions) > 0 {
			OrderedBy(SCORE).Sort(nextPositions)
			minNextScore := nextPositions[0].Score
			positions[i].Score += minNextScore
		} else {
			positions[i].Score += 10000000000000
		}
	}
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
