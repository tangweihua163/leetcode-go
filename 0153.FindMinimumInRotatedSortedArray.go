package leetcode_go

func findMin(nums []int) int {
	return findMinInRange(nums, 0, len(nums)-1)
}

func findMinInRange(nums []int, left, right int) int {

	if nums[left] < nums[right] {
		return nums[left]
	}

	var mid int
	for left <= right {

		if nums[left] < nums[right] {
			return nums[left]
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
			left = mid + 1
		} else {
			right = mid
		}
	}

	return mid

}
