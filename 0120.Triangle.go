package leetcode_go

func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	i := n - 2
	for i >= 0 {
		k := i
		for k >= 0 {
			v, _ := MinInt(triangle[i+1][k], triangle[i+1][k+1])
			triangle[i][k] += v
			k--
		}
		i--
	}

	return triangle[0][0]
}
