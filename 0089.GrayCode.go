package leetcode_go

func grayCode(n int) []int {
	result := make([]int, 1<<n)
	result = append(result, 0)
	for i := 0; i <= n-1; i++ {
		var m = 1 << i
		for j := len(result) - 1; j >= 0; j-- {
			result = append(result, result[j]+m)
		}
	}
	return result
}
