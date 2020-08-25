package leetcode_go

func restoreString(s string, indices []int) string {
	bs := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		bs[indices[i]] = s[i]
	}
	return string(bs)
}
