package leetcode_go

import "math"

func maxProduct(nums []int) int {

	var i, j = 0, 0
	for j < len(nums) && j+1 < len(nums) && j+2 < len(nums) {
		if nums[j] == 1 {
			if i > 0 && nums[i-1] > 0 {
				j++
			} else {
				nums[i] = nums[j]
				i++
				j++
			}
		} else if nums[j] == -1 && nums[j+1] == -1 && nums[j+2] == -1 {
			j += 2
		} else {
			nums[i] = nums[j]
			j++
			i++
		}
	}

	for j < len(nums) {
		nums[i] = nums[j]
		i++
		j++
	}

	nums = nums[:i]

	N := len(nums)
	dp := make([]int, N*N)

	var max = math.MinInt64

	for k := 0; k <= N-1; k++ {
		for i := 0; i < N-k; i++ {
			j := i + k
			if k == 0 {
				dp[i*N+j] = nums[i]
			} else {
				dp[i*N+j] = nums[i] * dp[(i+1)*N+j]
			}

			if dp[i*N+j] > max {
				max = dp[i*N+j]
			}

			if nums[j] == 0 {
				i = j
			}
		}
	}

	return max
}
