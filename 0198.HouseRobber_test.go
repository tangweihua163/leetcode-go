package leetcode_go

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	nums[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-1], nums[i-2]+nums[i])
	}

	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	return max
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
