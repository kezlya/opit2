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
