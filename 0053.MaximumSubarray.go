package leetcode_go

func maxSubArray(nums []int) int {
	a := nums[0]
	b := a
	for i := 1; i < len(nums); i++ {
		if a > 0 {
			a += nums[i]
		} else {
			a = nums[i]
		}
		if a > b {
			b = a
		}
	}
	return b
}
