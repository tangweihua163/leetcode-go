package leetcode_go

func generateParenthesis(n int) []string {

	var result []string
	stack := make([]byte, 0, n*2)
	return generate(n, n, stack, 0, result)
}

func generate(a, b int, stack []byte, c int, result []string) []string {

	if len(stack) == cap(stack) {
		result = append(result, string(stack))
		return result
	}

	if c == 0 {
		stack = append(stack, '(')
		c = 1
		a--
		result = generate(a, b, stack, c, result)
		stack = stack[:len(stack)-1]
		c--
	} else {
		if a > 0 {
			stack = append(stack, '(')
			c++
			a--
			result = generate(a, b, stack, c, result)
			stack = stack[:len(stack)-1]
			c--
			a++
		}

		if b > 0 {
			stack = append(stack, ')')
			c--
			b--
			result = generate(a, b, stack, c, result)
			stack = stack[:len(stack)-1]
			c++
			b++
		}

	}

	return result
}
