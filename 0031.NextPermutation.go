package leetcode_go

func nextPermutation(nums []int) {

	if len(nums) <= 1 {
		return
	}

	if len(nums) == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}

	i := len(nums) - 2

	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i == -1 {
		reverseArray(nums)
		return
	}

	// 执行到此处，nums[i] < nums[i+1] >= nums[i+2]

	j := search(nums, i+1, len(nums)-1, nums[i])

	nums[i], nums[j] = nums[j], nums[i]

	reverseArray(nums[i+1:])

}

func reverseArray(nums []int) {
	left := 0
	right := len(nums) - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func search(nums []int, start, end int, target int) int {

	if end-start <= 5 {
		for start <= end {
			if nums[start] > target {
				start++
			} else {
				break
			}
		}
		return start - 1
	}

	mid := (start + end) / 2

	if nums[mid] > target {
		return search(nums, mid+1, end, target)
	} else {
		return search(nums, start, mid-1, target)
	}

}
