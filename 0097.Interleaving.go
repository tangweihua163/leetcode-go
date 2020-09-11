package leetcode_go

func isInterleave(s1 string, s2 string, s3 string) bool {
	dp := make(map[int]bool)
	return isInterleaveR(s1, s2, s3, dp, len(s2))
}

func isInterleaveR(s1, s2, s3 string, dp map[int]bool, l int) (r bool) {

	index := len(s1)*l + len(s2)
	if v, ok := dp[index]; ok {
		return v
	}

	if len(s1)+len(s2) != len(s3) {
		dp[index] = false
		return false
	}

	if len(s1) == 0 {
		dp[index] = s2 == s3
		return s2 == s3
	}
	if len(s2) == 0 {
		dp[index] = s1 == s3
		return s1 == s3
	}

	i1 := len(s1) - 1
	i2 := len(s2) - 1
	i3 := len(s3) - 1

	// 都相同
	if s1[i1] == s3[i3] && s2[i2] == s3[i3] {
		b := isInterleaveR(s1[:i1], s2, s3[:i3], dp, l)
		if b {
			dp[index] = true
			return true
		}
		a := isInterleaveR(s1, s2[:i2], s3[:i3], dp, l)
		if a {
			dp[index] = true
			return true
		}
		dp[index] = false
		return false
	}

	// 都不相同
	if s1[i1] != s3[i3] && s2[i2] != s3[i3] {
		dp[index] = false
		return false
	}

	// 有一个相同
	if s3[i3] == s1[i1] {
		b := isInterleaveR(s1[:i1], s2, s3[:i3], dp, l)
		dp[index] = b
		return b
	}
	if s3[i3] == s2[i2] {
		b := isInterleaveR(s1, s2[:i2], s3[:i3], dp, l)
		dp[index] = b
		return b
	}

	dp[index] = false
	return false
}
