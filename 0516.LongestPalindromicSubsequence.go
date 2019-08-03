package leetcode_go

func longestPalindromeSubseq(s string) int {
	mark := make([]int, len(s)*len(s))

	return lps(s, 0, len(s)-1, mark)
}

func lps(s string, i, j int, mark []int) int {
	if j < i {
		return 0
	}

	if j == i {
		return 1
	}

	if v := mark[i*len(s)+j]; v > 0 {
		return v
	}

	if s[i] == s[j] {
		max := 2 + lps(s, i+1, j-1, mark)
		mark[i*len(s)+j] = max
		return max
	} else {
		a := lps(s, i+1, j, mark)
		b := lps(s, i, j-1, mark)
		max := Max(a, b)
		mark[i*len(s)+j] = max
		return max
	}
}
