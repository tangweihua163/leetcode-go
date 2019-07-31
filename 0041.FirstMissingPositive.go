package leetcode_go

func firstMissingPositive(nums []int) int {

	nums = nums[partition(nums)+1:]

	if len(nums) == 0 {
		return 1
	}

	BuildMinHeap(nums)

	last := 0
	for len(nums) > 0 {
		if nums[0]-last >= 2 {
			return last + 1
		}

		last = nums[0]

		nums[0] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		AdjustDownMinHeap(nums, 0)
	}

	return last + 1
}

func partition(nums []int) int {
	i := -1
	x := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return i
}
