package leetcode_go

func robII(nums []int) (max int) {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		max = Max(nums[0], nums[1])
		return
	}

	a := nums[0] + robI(nums[2:len(nums)-1])
	b := robI(nums[1:])

	return Max(a, b)
}

func robI(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return Max(nums[0], nums[1])
	}

	a := nums[0]
	b := Max(a, nums[1])
	var c int
	for i := 2; i < len(nums); i++ {
		c = Max(b, a+nums[i])
		a, b = b, c
	}

	return c
}
