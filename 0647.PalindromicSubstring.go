package leetcode_go

func countSubstrings(s string) int {
	count := 0
	n := len(s)
	mark := make([]int, n*n)
	for k := 0; k < n; k++ {
		for i := 0; i+k < n; i++ {
			if k <= 2 && s[i] == s[i+k] {
				mark[i*n+i+k] = 1
				count++
				continue
			}
			if s[i] == s[i+k] && mark[(i+1)*n+(i+k-1)] == 1 {
				mark[i*n+i+k] = 1
				count++
				continue
			}
		}
	}
	return count
}
