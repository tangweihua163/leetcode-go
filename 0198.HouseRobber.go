package leetcode_go

func rob(nums []int) (max int) {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		max, _ = MaxInt(nums[0], nums[1])
		return
	}

	nums[1], _ = MaxInt(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i], _ = MaxInt(nums[i-1], nums[i-2]+nums[i])
	}

	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	return
}
