package leetcode_go

func findMinII(nums []int) int {
	return findMinInRangeII(nums, 0, len(nums)-1)
}

func findMinInRangeII(nums []int, left, right int) int {

	var mid int

	for left < right && nums[left] == nums[right] {
		right--
	}

	if right-left <= 5 {
		min := nums[left]
		for i := left; i <= right; i++ {
			if nums[i] < min {
				min = nums[i]
			}
		}
		return min
	}

	for left <= right {

		if nums[left] < nums[right] {
			return nums[left]
		}

		if nums[left] == nums[right] {
			right--
			continue
		}

		if right-left <= 5 {
			min := nums[left]
			for i := left; i <= right; i++ {
				if nums[i] < min {
					min = nums[i]
				}
			}
			return min
		}

		mid = (left + right) / 2

		if nums[mid] >= nums[left] {
			left = mid
		} else {
			right = mid
		}
	}
	return mid
}
