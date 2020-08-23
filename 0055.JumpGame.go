package leetcode_go

func canJump(nums []int) bool {

	var max = 0
	var i int
	for i = 0; i < len(nums) && i <= max; i++ {
		if nums[i]+i > max {
			max = nums[i] + i
		}
	}

	return i == len(nums)
}
