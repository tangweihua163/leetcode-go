package leetcode_go

import "math"

func maxSubArray(nums []int) int {

	a := nums[0]
	b := math.MinInt64

	maxSum := max(a, b)

	for i := 1; i < len(nums); i++ {
		b = max(a, b)

		if b > maxSum {
			maxSum = b
		}

		if a > 0 {
			a += nums[i]
		} else {
			a = nums[i]
		}

		if a > maxSum {
			maxSum = a
		}

	}

	return maxSum
}
