package leetcode_go

func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	i := n - 2
	j := n - 1
	for i >= 0 {
		k := i
		for k >= 0 {
			triangle[i][k] += min(triangle[j][k], triangle[j][k+1])
			k--
		}
		i--
		j--
	}

	return triangle[0][0]
}

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
