package leetcode_go

func moveZeroes(nums []int) {
	index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}

	for index < len(nums) {
		nums[index] = 0
		index++
	}
}
