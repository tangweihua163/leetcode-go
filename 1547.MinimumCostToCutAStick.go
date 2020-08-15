package leetcode_go

import (
	"math"
	"sort"
)

func minCost(n int, cuts []int) int {

	sort.Ints(cuts)
	var m = make([]int, (n+1)*(n+1))

	findmincost(n, 0, n, cuts, m)

	return m[n]
}

func findmincost(n int, l int, r int, cuts []int, m []int) int {
	if len(cuts) == 0 {
		return 0
	}

	var index = (n+1)*l + r
	if m[index] != 0 {
		return m[index]
	}

	var min = math.MaxInt32
	for i := 0; i < len(cuts); i++ {
		var t = findmincost(n, l, cuts[i], cuts[:i], m) + findmincost(n, cuts[i], r, cuts[i+1:], m) + r - l
		if t < min {
			min = t
		}
	}

	m[index] = min

	return min
}
