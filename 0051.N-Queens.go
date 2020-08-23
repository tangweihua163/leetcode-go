package leetcode_go

func solveNQueens(n int) [][]string {
	stack := make([]int, 0, n)
	result := make([][]int, 0)
	visited := make([]bool, n)

	result = solve1(n, stack, visited, result)

	return graph(result, n)
}

func solve1(n int, stack []int, visited []bool, result [][]int) [][]int {

	if len(stack) == n {
		dst := make([]int, n)
		copy(dst, stack)
		return append(result, dst)
	}

	var i1 = len(stack)
	for j1 := 0; j1 < n; j1++ {

		if visited[j1] {
			continue
		}

		b := true
		for i2 := 0; i2 < i1; i2++ {
			j2 := stack[i2]
			if conflict(i1, j1, i2, j2) {
				b = false
				break
			}
		}

		if b {
			visited[j1] = true
			result = solve1(n, append(stack, j1), visited, result)
			visited[j1] = false
		}
	}

	return result
}

func graph(arr [][]int, n int) [][]string {
	result := make([][]string, len(arr))

	var line = make([]byte, n)
	for i := 0; i < n; i++ {
		line[i] = '.'
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < n; j++ {
			dst := make([]byte, n)
			copy(dst, line)
			dst[arr[i][j]] = 'Q'
			result[i] = append(result[i], string(dst))
		}
	}

	return result
}

func conflict(a, b, c, d int) bool {
	return abs1(a-c) == abs1(b-d)
}

func abs1(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
