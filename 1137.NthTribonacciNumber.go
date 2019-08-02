package leetcode_go

func tribonacci(n int) int {

	if n <= 2 {
		return (n + 1) / 2
	}

	a, b, c := 0, 1, 1
	for i := 0; i <= n-3; i++ {
		a, b, c = b, c, a+b+c
	}

	return c
}
