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
	//TODO: choose plasements clother to the wall
	//zone := _getZone()
	positions := g.MyPlayer.Field.Positions(g.CurrentPiece)
	burndPositions := _getBurned(positions)

	if len(burndPositions) > 0 {
		OrderedBy(BURN, DAMAGE).Sort(burndPositions)
		return &burndPositions[0]
	}

	/*if len(burndPositions) > 0 && (MyPlayer.Combo > 0 || zone == "dangerous") {
		return _keepUpBurn(burndPositions)
	}*/

	return g.choosePosition(positions)
}

func (g *Game) choosePosition(positions []Position) *Position {
	if len(positions) > 1 {
		OrderedBy(SCORE, DAMAGE, HIGHY).Sort(positions)
		bIndex := 0
		sumScore := 1000
		for current_i, pos := range positions {
			nextPiecePositions := pos.FieldAfter.Positions(g.NextPiece)
			for _, nextPos := range nextPiecePositions {
				if pos.Score+nextPos.Score < sumScore {
					sumScore = pos.Score + nextPos.Score
					bIndex = current_i
				}
				//TODO:check burn game in unittest
			}
		}
		return &positions[bIndex]
	}

	if len(positions) == 0 {
		return nil
	}

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

func (g *Game) keepUpBurn(burnedPos []Position) *Position {
	if len(burnedPos) > 1 {
		OrderedBy(BURN, DAMAGE).Sort(burnedPos)
		bIndex := 0
		for current_i, pos := range burnedPos {
			pos.FieldAfter.Burn()
			nextPiecePositions := pos.FieldAfter.Positions(g.NextPiece)
			for _, nextPos := range nextPiecePositions {
				if nextPos.IsBurn > 0 {
					bIndex = current_i
					break
				}
			}
		}
		return &burnedPos[bIndex]
	}
	return &burnedPos[0]
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
