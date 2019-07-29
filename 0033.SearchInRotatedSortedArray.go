package leetcode_go

func searchInRotatedSortedArray(nums []int, target int) int {
	return searchInRange(nums, 0, len(nums)-1, target)
}

func searchInRange(nums []int, left, right int, target int) int {

	if right-left <= 5 {
		return linearSearch(nums, left, right, target)
	}

	if nums[left] < nums[right] {
		return binarySearch(nums, left, right, target)
	}

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
			if target >= nums[left] && target <= nums[mid] {
				return binarySearch(nums, left, mid-1, target)
			} else {
				left = mid + 1
				continue
			}
		}

		if nums[mid] <= nums[right] {
			if target >= nums[mid] && target <= nums[right] {
				return binarySearch(nums, mid+1, right, target)
			} else {
				right = mid - 1
				continue
			}
		}
	}

	return -1

}

func binarySearch(nums []int, left, right int, target int) int {

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

func linearSearch(nums []int, left, right int, target int) int {

	for i := left; i <= right; i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}
