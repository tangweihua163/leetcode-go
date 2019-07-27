package leetcode_go

func singleNumberII(nums []int) int {
	times := make(map[int]int)

	for _, v := range nums {
		times[v] = times[v] + 1
	}

	for k, v := range times {
		if v == 1 {
			return k
		}
	}

	return -1
}
