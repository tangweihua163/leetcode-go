package leetcode_go

func longestPalindrome(s string) string {
	var resultStr string
	var maxLength int

	// 奇数长度
	if len(s)&1 == 1 {

		mid := len(s) / 2

		// 奇数回文
		for i := 0; i <= mid; i++ {

			if (mid-i)*2+1 < maxLength {
				break
			}

			ss, l := odd(s, mid-i)
			if l > maxLength {
				resultStr = ss
				maxLength = l
			}

			ss, l = odd(s, mid+i)
			if l > maxLength {
				resultStr = ss
				maxLength = l
			}
		}

		// 偶数回文
		for i := 0; i < mid; i++ {

			if (mid-i)*2 < maxLength {
				break
			}

			ss, l := even(s, mid-i-1, mid-i)
			if l > maxLength {
				resultStr = ss
				maxLength = l
			}

			ss, l = even(s, mid+i, mid+i+1)
			if l > maxLength {
				resultStr = ss
				maxLength = l
			}
		}

	}

	// 偶数长度
	if len(s)&1 == 0 {

		mid := len(s) / 2

		// 偶数回文
		for i := 0; i <= mid-1; i++ {

			if (mid-i)*2 < maxLength {
				break
			}

			ss, l := even(s, mid-1-i, mid-i)
			if l > maxLength {
				resultStr = ss
				maxLength = l
			}

			ss, l = even(s, mid-1+i, mid+i)
			if l > maxLength {
				resultStr = ss
				maxLength = l
			}
		}

		// 奇数回文
		for i := 0; i < mid; i++ {

			if (mid-i)*2+1 < maxLength {
				break
			}

			ss, l := odd(s, mid-1-i)
			if l > maxLength {
				resultStr = ss
				maxLength = l
			}

			ss, l = odd(s, mid+i)
			if l > maxLength {
				resultStr = ss
				maxLength = l
			}
		}
	}

	return resultStr
}

// 以 i 为中心最长的奇数长度的回文
func odd(s string, i int) (result string, length int) {
	left := i - 1
	right := i + 1
	length = 1

	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			length += 2
		} else {
			break
		}
		left--
		right++
	}

	return s[left+1 : right], length
}

// 以 i、j 为中心偶数长度的回文
func even(s string, i, j int) (result string, length int) {
	left := i
	right := j

	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			length += 2
		} else {
			break
		}
		left--
		right++
	}

	return s[left+1 : right], length
}
