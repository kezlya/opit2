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

func (p Picks) Damage(a Picks) (int, int, int) {
	highY := 0
	lowY := 1000
	damage := 0
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
		}
	}
	return damage, lowY, highY
}

func (p Picks) IsHole() bool {
	for i, c := range p {
		if p.IsRight(i, 1) && (c-p[i+1] < -2 || c-p[i+1] > 2) {
			return true
		}
	}
	return false
}
