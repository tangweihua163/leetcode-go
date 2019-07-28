package leetcode_go

func majorityElement(nums []int) int {

	if len(nums) <= 2 {
		return nums[0]
	}

	low := 0
	high := len(nums) - 1
	mid := 0

	for low <= high {

		mid = low - 1
		x := nums[high]

		for j := low; j <= high; j++ {
			if nums[j] <= x {
				mid++
				nums[mid], nums[j] = nums[j], nums[mid]
			}
		}

		if mid == len(nums)/2 {
			return nums[mid]
		}

		if mid > len(nums)/2 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return nums[mid]
}

func majorityElement1(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v] = m[v] + 1
	}

	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}

	return nums[0]
}

func majorityElement2(nums []int) int {

	result := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			result = nums[i]
			count = 1
		} else {
			if nums[i] != result {
				count--
			} else {
				count++
			}
		}
	}

	return result
}
