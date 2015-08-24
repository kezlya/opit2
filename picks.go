package main

type Picks []int

func (p Picks) Max() int {
	result := 0
	for _, pick := range p {
		if result < pick {
			result = pick
		}
	}
	return result
}

func (p Picks) MaxR(x, n int) int {
	pick := p[x]
	for i := 1; i <= n; i++ {
		if pick < p[x+i] {
			pick = p[x+i]
		}
	}
	return pick
}

func (p Picks) IsRight(x, n int) bool {
	if x+n < len(p) {
		return true
	}
	return false
}

func (p Picks) Equal(b Picks) bool {
	if len(p) != len(b) {
		return false
	}
	for i := range p {
		if p[i] != b[i] {
			return false
		}
	}
	return true
}

func (p Picks) Damage(a Picks) (int, int, int, int) {
	highY := 0
	lowY := 1000
	step := 0
	damage := 0
	_left := -1
	_right := 0
	for i, col := range p {
		diff := a[i] - col
		if diff > 0 {
			damage = damage + diff

			if col < lowY {
				lowY = col
			}

			if a[i] > highY {
				highY = a[i]
			}

			if step == -1 {
				step = i
			}

			_right = i
		}
	}

	if _left > 0 {
		lDif := a[_left-1] - a[_left]
		if lDif < 0 {
			lDif = -lDif
		}
		step += lDif
	}
	if _right < len(a)-1 {
		rDif := a[_right+1] - a[_right]
		if rDif < 0 {
			rDif = -rDif
		}
		step += rDif
	}

	return damage, lowY, highY, step
}
