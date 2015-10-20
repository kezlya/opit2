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

func (p Picks) IsTowers() bool {
	counter := 0
	for i, pick := range p {
		if i > 0 {
			diff := p[i-1] - pick
			if diff > 3 || diff < -3 {
				counter++
			}
		}
	}
	return counter >= 2
}

func (p Picks) SumStep() int {
	sum := 0
	for i, v := range p {
		if i > 0 {
			diff := v - p[i-1]
			if diff > 0 {
				sum = sum + diff
			} else {
				sum = sum - diff
			}
		}
	}
	return sum
}
