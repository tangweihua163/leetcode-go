package leetcode_go

func searchInRotatedSortedArray(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	var left, right = 0, len(nums) - 1
	for left <= right {

		if right-left < 5 {
			for i := left; i <= right; i++ {
				if nums[i] == target {
					return i
				}
			}
			return -1
		}

		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}

		if target < nums[left] && target > nums[right] {
			return -1
		}

		if nums[left] < nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else if nums[mid] < nums[right] {
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}

	return -1
}
