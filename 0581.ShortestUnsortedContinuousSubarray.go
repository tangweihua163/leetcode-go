package leetcode_go

import "math"

func findUnsortedSubarray(nums []int) int {
	leftIndex := -1
	rightIndex := -1

	result := len(nums)

	top := math.MinInt64
	f := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= top {
			if f == 1 {
				top = nums[i]
				leftIndex++
			}
		} else {
			for leftIndex >= 0 && top > nums[i] {
				leftIndex--
				if leftIndex >= 0 {
					top = nums[leftIndex]
				}
			}
			if leftIndex == -1 {
				break
			}
			f = 0
		}
	}

	if leftIndex > -1 {
		result -= leftIndex + 1
	}

	top = math.MaxInt64
	f = 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= top {
			if f == 1 {
				top = nums[i]
				rightIndex++
			}
		} else {
			for rightIndex >= 0 && top < nums[i] {
				rightIndex--
				if rightIndex >= 0 {
					top = nums[len(nums)-1-rightIndex]
				}
			}
			if rightIndex == -1 {
				break
			}
			f = 0
		}
	}

	if rightIndex > -1 {
		result -= rightIndex + 1
	}

	if result < 0 {
		return 0
	}

	return result

}
