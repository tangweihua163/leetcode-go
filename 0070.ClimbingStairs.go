package leetcode_go

func climbStairs(n int) int {

	if n <= 3 {
		return n
	}

	var a, b, r = 1, 2, 3

	for i := 4; i <= n; i++ {
		a, b = b, r
		r = a + b
	}

	return r
}
