package leetcode_go

import (
	"math"
	"sort"
)

func minCost(n int, cuts []int) int {

	cuts = append(cuts, 0, n)

	sort.Ints(cuts)

	var m = make(map[int]int)

	for k := 1; k <= len(cuts)-1; k++ {
		for i := 0; i <= len(cuts)-1-k; i++ {
			j := i + k

			var a, b int
			a = cuts[i]
			b = cuts[j]

			if k == 1 {
				m[(n+1)*a+b] = 0
				continue
			}

			m[(n+1)*a+b] = math.MaxInt64
			for p := i + 1; p < j; p++ {
				m[(n+1)*a+b] = min(m[(n+1)*a+b], m[(n+1)*a+cuts[p]]+m[(n+1)*cuts[p]+b]+b-a)
			}
		}
	}

	return m[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
