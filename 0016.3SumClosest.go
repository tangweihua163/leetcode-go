package leetcode_go

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	gap := math.MaxInt32
	closestSum := 0

	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {

			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return target
			}

			if sum > target && sum-target < gap {
				gap = sum - target
				closestSum = sum
				right--
			} else if sum < target && target-sum < gap {
				gap = target - sum
				closestSum = sum
				left++
			} else if sum < target {
				left++
			} else if sum > target {
				right--
			}
		}
	}

	return closestSum
}
