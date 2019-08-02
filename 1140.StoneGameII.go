package leetcode_go

func stoneGameII(piles []int) int {
	var sum int
	for i := 0; i < len(piles); i++ {
		sum += piles[i]
	}
	var m = make(map[[2]int]int)
	return play(piles, 0, 1, sum, m)
}

func play(piles []int, start int, M int, sum int, m map[[2]int]int) int {

	if start+2*M >= len(piles) {
		return sum
	}

	var max, get int
	for x := start; x <= 2*M-1+start && x < len(piles); x++ {
		get += piles[x]

		var op int
		if v, ok := m[[2]int{x + 1, Max(M, x-start+1)}]; ok {
			op = v
		} else {
			op = play(piles, x+1, Max(M, x-start+1), sum-get, m)
		}
		max = Max(sum-op, max)
	}
	m[[2]int{start, M}] = max
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
