package leetcode_go

func isPalindromeString(s string) bool {

	if len(s) <= 1 {
		return true
	}

	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		if ss[i] >= 'a' && ss[i] <= 'z' {
			ss[i] = ss[i] - ('a' - 'A')
		}
	}

	left := 0
	right := len(ss) - 1

	for left < len(ss) && !isAlpha(ss[left]) {
		left++
	}
	for right >= 0 && !isAlpha(ss[right]) {
		right--
	}

	for left < right {
		if ss[left] != ss[right] {
			return false
		}
		left++
		right--

		for left < len(ss) && !isAlpha(ss[left]) {
			left++
		}
		for right >= 0 && !isAlpha(ss[right]) {
			right--
		}
	}

	return true
}

func isAlpha(b byte) bool {
	return b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' || b >= '0' && b <= '9'
}
