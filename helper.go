package main

import (
//"strings"
)

func _fitsI(c []int) []Position {
	var pos []Position
	for i, v := range c {
		if (_isRight(i, 1) && v+1 < c[i+1]) || (_isLeft(i, 1) && v+1 < c[i-1]) {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 3) && v == c[i+1] && v == c[i+2] && v == c[i+3] {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if (_isRight(i, 1) && v < c[i+1]) || (_isLeft(i, 1) && v < c[i-1]) {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}
	return pos
}

func _fitsJ(c []int) []Position {
	var pos []Position
	for i, v := range c {
		if _isRight(i, 1) && v+2 == c[i+1] {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 2) && v == c[i+1] && v == c[i+2]+1 {
			p := Position{Rotation: 2, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 2) && v == c[i+1] && v == c[i+2] {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 1) && v == c[i+1] {
			p := Position{Rotation: 3, X: i}
			pos = append(pos, p)
		}
	}

	return pos
}

func _fitsL(c []int) []Position {
	var pos []Position
	for i, v := range c {
		if _isRight(i, 1) && v == c[i+1]+2 {
			p := Position{Rotation: 3, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 2) && v+1 == c[i+1] && v+1 == c[i+2] {
			p := Position{Rotation: 2, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 2) && v == c[i+1] && v == c[i+2] {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 1) && v == c[i+1] {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}

	return pos
}

func _fitsO(c []int) []Position {
	var pos []Position
	for i, v := range c {
		if (_isRight(i, 2) && v == c[i+1] && v < c[i+2]) || (_isLeft(i, 1) && v < c[i-1] && _isRight(i, 1) && v == c[i+1]) {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 1) && v == c[i+1] {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	return pos
}

func _fitsS(c []int) []Position {
	var pos []Position
	for i, v := range c {
		if _isRight(i, 2) && v == c[i+1] && v+1 == c[i+2] {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 1) && v == c[i+1]+1 {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}

	return pos
}

func _fitsT(c []int) []Position {
	var pos []Position
	for i, v := range c {
		if _isRight(i, 2) && v == c[i+1]+1 && v == c[i+2] {
			p := Position{Rotation: 2, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 2) && v == c[i+1] && v == c[i+2] {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 1) && v+1 == c[i+1] {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 1) && v == c[i+1]+1 {
			p := Position{Rotation: 3, X: i}
			pos = append(pos, p)
		}
	}

	return pos
}

func _fitsZ(c []int) []Position {
	var pos []Position
	for i, v := range c {
		if _isRight(i, 2) && v == c[i+1]+1 && v == c[i+2]+1 {
			p := Position{Rotation: 0, X: i}
			pos = append(pos, p)
		}
	}

	for i, v := range c {
		if _isRight(i, 1) && v+1 == c[i+1] {
			p := Position{Rotation: 1, X: i}
			pos = append(pos, p)
		}
	}

	return pos
}

func _isRight(i, right int) bool {
	if i+right < Width {
		return true
	}
	return false
}

func _isLeft(i, left int) bool {
	if i-left > 0 {
		return true
	}
	return false
}

func _getColumnsAfter(c []int, i, r int, piece string) ([]int, int) {
	a := make([]int, len(c))
	copy(a, c[:])
	y := 0
	switch piece {
	case "I":
		switch r {
		case 0:
			if _isRight(i, 3) {
				pick := _getPick(i, 3)
				a[i] = pick + 1
				a[i+1] = pick + 1
				a[i+2] = pick + 1
				a[i+3] = pick + 1
				y = pick + 1
			}
		case 1:
			a[i] = c[i] + 4
			y = c[i] + 4
		}
	case "J":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				a[i] = pick + 2
				a[i+1] = pick + 1
				a[i+2] = pick + 1
				y = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				if c[i+1] >= c[i]+2 {
					a[i] = c[i+1] + 1
					a[i+1] = c[i+1] + 1
					y = c[i+1] + 1
				} else {
					a[i] = c[i] + 3
					a[i+1] = c[i] + 3
					y = c[i] + 3
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i+2] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				a[i] = pick + 1
				a[i+1] = pick + 3
				y = pick + 3
			}
		}
	case "L":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				a[i] = pick + 1
				a[i+1] = pick + 1
				a[i+2] = pick + 2
				y = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				a[i] = pick + 3
				a[i+1] = pick + 1
				y = pick + 3
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				if c[i] >= c[i+1]+2 {
					a[i] = c[i] + 1
					a[i+1] = c[i] + 1
					y = c[i] + 1
				} else {
					a[i] = c[i+1] + 3
					a[i+1] = c[i+1] + 3
					y = c[i] + 3
				}
			}
		}
	case "O":
		if _isRight(i, 1) {
			pick := _getPick(i, 1)
			a[i] = pick + 2
			a[i+1] = pick + 2
			y = pick + 2
		}
	case "S":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i] || pick == c[i+1] {
					a[i] = pick + 1
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i+1] {
					a[i] = pick + 3
					a[i+1] = pick + 2
					y = pick + 3
				} else {
					a[i] = pick + 2
					a[i+1] = pick + 1
					y = pick + 2
				}
			}
		}
	case "T":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				a[i] = pick + 1
				a[i+1] = pick + 2
				a[i+2] = pick + 1
				y = pick + 2
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i] {
					a[i] = pick + 3
					a[i+1] = pick + 2
					y = pick + 3
				} else {
					a[i] = pick + 2
					a[i+1] = pick + 1
					y = pick + 2
				}
			}
		case 2:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i+1] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 2
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick + 1
					y = pick + 1
				}
			}
		case 3:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i+1] {
					a[i] = pick + 2
					a[i+1] = pick + 3
					y = pick + 3

				} else {
					a[i] = pick + 1
					a[i+1] = pick + 2
					y = pick + 2
				}
			}
		}
	case "Z":
		switch r {
		case 0:
			if _isRight(i, 2) {
				pick := _getPick(i, 2)
				if pick == c[i+1] || pick == c[i+2] {
					a[i] = pick + 2
					a[i+1] = pick + 2
					a[i+2] = pick + 1
					y = pick + 2
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 1
					a[i+2] = pick
					y = pick + 1
				}
			}
		case 1:
			if _isRight(i, 1) {
				pick := _getPick(i, 1)
				if pick == c[i] {
					a[i] = pick + 2
					a[i+1] = pick + 3
					y = pick + 3
				} else {
					a[i] = pick + 1
					a[i+1] = pick + 2
					y = pick + 2
				}
			}
		}
	}
	return a, y
}

func _getPick(i, v int) int {
	pick := MyPlayer.Columns[i]
	for n := 1; n <= v; n++ {
		if pick < MyPlayer.Columns[i+n] {
			pick = MyPlayer.Columns[i+n]
		}
	}
	return pick
}
