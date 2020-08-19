package leetcode_go

import "math"

func minDays(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	var min = math.MaxInt32
	var a = n%2 + minDays(n/2)
	if min > a {
		min = a
	}

	var b = n%3 + minDays(n/3)
	if min > b {
		min = b
	}

	return min + 1
}
