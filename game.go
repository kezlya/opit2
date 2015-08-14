package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Timebank       int
	TimePerMove    int
	Width          int
	Height         int
	OriginalHeight int
	Round          int
	CurrentPieceX  int
	CurrentPieceY  int
	Players        []Player
	MyPlayer       *Player
	CurrentPiece   string
	NextPiece      string

	BurnK   int
	HoleK   int
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
			g.CurrentPiece = value
		case "next_piece_type":
			g.NextPiece = value
		case "this_piece_position":
			cor := strings.Split(value, ",")
			g.CurrentPieceX, _ = strconv.Atoi(cor[0])
			g.CurrentPieceY, _ = strconv.Atoi(cor[1])
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
			cleanSource := strings.Replace(value, ";3,3,3,3,3,3,3,3,3,3", "", g.OriginalHeight)
			for i, p := range g.Players {
				if p.Name == who {
					var empty Field
					g.Players[i].Field = empty.init(cleanSource)
					break
				}
			}
		}
	}
}

func (g *Game) calculateMoves() *Position {

	positions := g.MyPlayer.Field.Positions(g.CurrentPiece, g.DamageK, g.PostyK, g.HoleK, g.BurnK)

	for i, position := range positions {
		position.FieldAfter.Burn()
		nextPositions := position.FieldAfter.Positions(g.NextPiece, g.DamageK, g.PostyK, g.HoleK, g.BurnK)
		OrderedBy(SCORE).Sort(nextPositions)
		minNextScore := nextPositions[0].Score
		positions[i].Score += minNextScore
	}

	OrderedBy(SCORE).Sort(positions)

	/*if len(burndPositions) > 0 && (MyPlayer.Combo > 0 || zone == "dangerous") {
		return _keepUpBurn(burndPositions)
	}*/

	return &positions[0]
}

func (g *Game) getZone() string {
	picks := g.MyPlayer.Field.Picks()
	y := picks.Max()
	rowsLeft := g.MyPlayer.Field.Height() - y
	if rowsLeft > 13 {
		return "safe"
	}
	return "dangerous"
}

func (g *Game) printMoves() {
	pos := g.calculateMoves()

	var buffer bytes.Buffer
	for i := 0; i < pos.Rotation; i++ {
		buffer.WriteString("turnright,")
	}
	if pos.Rotation == 1 {
		g.CurrentPieceX = g.CurrentPieceX + 1
		if g.CurrentPiece == "I" {
			g.CurrentPieceX = g.CurrentPieceX + 1
		}
	}
	if g.CurrentPieceX > pos.X {
		for i := 0; i < g.CurrentPieceX-pos.X; i++ {
			buffer.WriteString("left,")
		}
	}
	if g.CurrentPieceX < pos.X {
		for i := 0; i < pos.X-g.CurrentPieceX; i++ {
			buffer.WriteString("right,")
		}
	}
	buffer.WriteString("drop")
	fmt.Println(buffer.String())
}
