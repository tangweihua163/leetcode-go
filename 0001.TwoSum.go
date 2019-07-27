package leetcode_go

func twoSum(nums []int, target int) []int {
	index := make(map[int]int)

	for iv, v := range nums {

		c := target - v
		ic, ok := index[c]

		if ok {
			return []int{ic, iv}
		}

		index[v] = iv
	}

	return []int{0, 0}
}
