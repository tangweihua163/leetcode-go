package leetcode_go

import "fmt"

func maxNonOverlapping(nums []int, target int) int {

	n := len(nums)

	m := make([]int, n+1)
	m[n] = -1

	for i := 0; i < n; i++ {
		m[i] = -1
		var j = i
		var sum = 0
		for j < n {
			sum += nums[j]
			if sum == target {
				m[i] = j
				break
			}
			j++
		}
	}

	mm := make([]int, n+1)
	mm[n] = 0

	for i := n - 1; i >= 0; i-- {
		if m[i] == -1 {
			mm[i] = mm[i+1]
		} else {
			var a = mm[i+1]
			var b = mm[m[i]+1] + 1
			if a < b {
				mm[i] = b
			} else {
				mm[i] = a
			}
		}
		fmt.Println(mm)
	}

	return mm[0]
}
