package leetcode_go

func partitions(s string) [][]string {

	result := make([][]string, 0)
	stack := make([]string, 0, len(s))

	result = parti(s, result, stack)

	return result
}

func parti(s string, result [][]string, stack []string) [][]string {
	if len(s) == 0 {
		dst := make([]string, len(stack))
		copy(dst, stack)
		result = append(result, dst)
		return result
	}

	for i := 0; i < len(s); i++ {
		if isPalin(s[:i+1]) {
			stack = append(stack, s[:i+1])
			result = parti(s[i+1:], result, stack)
			stack = stack[:len(stack)-1]
		}
	}

	return result
}

func isPalin(s string) bool {
	if len(s) < 2 {
		return true
	}

	var i, j = 0, len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
