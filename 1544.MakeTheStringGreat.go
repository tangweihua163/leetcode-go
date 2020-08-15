package leetcode_go

func makeGood(s string) string {

	var stack = make([]byte, 0, len(s)+1)
	stack = append(stack, '#')
	var top = 0

	for i := 0; i < len(s); i++ {

		if lower(stack[top]) == lower(s[i]) && stack[top] != s[i] {
			stack = stack[:top]
			top--
		} else {
			stack = append(stack, s[i])
			top++
		}
	}

	return string(stack[1:])
}

func lower(b byte) byte {

	if b >= 'A' && b <= 'Z' {
		return b + ('z' - 'Z')
	}

	return b
}
