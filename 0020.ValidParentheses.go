package leetcode_go

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	top := -1

	for i := 0; i < len(s); i++ {
		b := s[i]
		if top >= 0 && match(stack[top], b) {
			stack = stack[:len(stack)-1]
			top--
		} else {
			stack = append(stack, b)
			top++
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func match(b1, b2 byte) bool {
	if b1 == '[' && b2 == ']' {
		return true
	}

	if b1 == '{' && b2 == '}' {
		return true
	}

	if b1 == '(' && b2 == ')' {
		return true
	}

	return false
}
