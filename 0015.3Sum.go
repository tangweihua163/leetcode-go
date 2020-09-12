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

func threeSum2(nums []int) [][]int {

	sort.Ints(nums)

	result := make([][]int, 0, 10)
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for i := 0; i < len(nums)-1; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		m[nums[i]]--

		for j := i + 1; j < len(nums); j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			t := 0 - nums[i] - nums[j]

			if t < nums[j] {
				break
			}

			m[nums[j]]--

			if m[t] > 0 {
				result = append(result, []int{nums[i], nums[j], t})
			}

			m[nums[j]]++

			if t == nums[j] {
				break
			}
		}
	}

	return result
}
