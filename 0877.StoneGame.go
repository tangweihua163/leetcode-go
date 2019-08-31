package leetcode_go

func stoneGame(piles []int) bool {

	var m = make(map[int]int)

	var sum int
	for _, p := range piles {
		sum += p
	}

	return playStoneGame(piles, 0, len(piles)-1, m, sum) > sum/2
}

func playStoneGame(piles []int, start, end int, m map[int]int, sum int) int {
	if v, ok := m[start*len(piles)+end]; ok {
		return v
	}

	if end-start == 0 {
		m[start*len(piles)+end] = piles[start]
		return piles[start]
	}

	var min int
	op1 := playStoneGame(piles, start+1, end, m, sum-piles[start])
	op2 := playStoneGame(piles, start, end-1, m, sum-piles[end])

	min = op1
	if op1 > op2 {
		min = op2
	}

	m[start*len(piles)+end] = sum - min
	return sum - min
}
