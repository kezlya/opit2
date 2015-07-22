package main

import (

)

type Position struct {
	Rotation int
	X int
}

func FindFitPositions(piece string, col []int) Position{
	pos := Position{}
	pos.Rotation=1
	pos.X=0

	return pos
}