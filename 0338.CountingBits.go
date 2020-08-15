package leetcode_go

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	var result = make([]int, num+1)
	result[0] = 0
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			result[i] = result[i/2]
		} else {
			result[i] = result[i-1] + 1
		}
	}

	return result
}
