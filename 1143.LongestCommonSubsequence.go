package leetcode_go

func longestCommonSubsequence(text1 string, text2 string) int {

	l1 := len(text1)
	l2 := len(text2)

	mark := make([]int, l1*l2)

	return lcs(text1, text2, 0, 0, mark)

}

func lcs(text1, text2 string, i, j int, mark []int) int {
	if i >= len(text1) || j >= len(text2) {
		return 0
	}

	if v := mark[i*len(text2)+j]; v > 0 {
		return v
	}

	if text1[i] == text2[j] {
		a := 1 + lcs(text1, text2, i+1, j+1, mark)
		mark[i*len(text2)+j] = a
		return a
	} else {
		a := lcs(text1, text2, i+1, j, mark)
		b := lcs(text1, text2, i, j+1, mark)
		max := Max(a, b)
		mark[i*len(text2)+j] = max
		return max
	}
}
