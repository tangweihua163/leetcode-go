package leetcode_go

func firstUniqChar(s string) int {
	var m = make([]int, 256)

	for i := 1; i < len(s); i++ {
		if m[s[i]] <= 2 {
			m[s[i]]++
		}
	}

	for i := 1; i < len(s); i++ {
		if m[s[i]] == 1 {
			return i
		}
	}

	return -1
}
