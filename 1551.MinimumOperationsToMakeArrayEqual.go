package leetcode_go

func minOperations(n int) int {
	if n%2 == 1 {
		return (n - 1) / 2 * ((n-1)/2 + 1)
	} else {
		return n/2 + (n-2)/2*((n-2)/2+1)
	}
}
