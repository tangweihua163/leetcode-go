package leetcode_go

var m = map[byte]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {

	var result []string
	var stack = make([]byte, 0, len(digits))

	return combinations(digits, len(digits), stack, result)
}

func combinations(digits string, l int, stack []byte, result []string) []string {
	if len(digits) <= 0 {
		return result
	}

	str := digitOf(digits, 0)

	for i := 0; i < len(str); i++ {
		stack = append(stack, str[i])
		if len(stack) == l {
			result = append(result, string(stack))
		}

		if len(digits) > 1 {
			result = combinations(digits[1:], l, stack, result)
		}

		stack = stack[:len(stack)-1]
	}

	return result
}

func digitOf(digits string, index int) string {
	return m[digits[index]-byte('0')]
}
