package leetcode_go

func sortColors(nums []int) {
	var i, j = 0, len(nums) - 1

	for k := 0; k <= j; k++ {
		if nums[k] == 0 {
			nums[k], nums[i] = nums[i], nums[k]
			i++
		} else if nums[k] == 2 {
			nums[k], nums[j] = nums[j], nums[k]
			j--
			k--
		}
	}

}
