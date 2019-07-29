package leetcode_go

import "sort"

func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	var result [][]int

	for i := 0; i < len(nums)-2; i++ {

		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left+1 < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				for right-1 > left && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			}
		}
	}
	return result
}
