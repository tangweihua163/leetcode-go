package leetcode_go

import "math"

func reverse(x int) int {
	var res int
	for x != 0 {
		if abs(res) > math.MaxInt32/10 {
			return 0
		}

		res = res*10 + x%10
		x = x / 10
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
