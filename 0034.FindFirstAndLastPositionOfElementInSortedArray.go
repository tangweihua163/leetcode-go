package leetcode_go

func searchRange(nums []int, target int) []int {

	return searchRangeInRange(nums, 0, len(nums)-1, target)

}

func searchRangeInRange(nums []int, left, right int, target int) []int {

	result := []int{-1, -1}

	for left <= right {

		if right-left <= 5 {
			return linearSearchRange(nums, left, right, target)
		}

		mid := (left + right) / 2

		if nums[mid] == target {
			if nums[mid-1] != target {
				result[0] = mid
			} else {
				result[0] = searchFirst(nums, left, mid-1, target)
			}
			if nums[mid+1] != target {
				result[1] = mid
			} else {
				result[1] = searchLast(nums, mid+1, right, target)
			}
			return result
		} else if nums[mid] < target {
			return searchRangeInRange(nums, mid+1, right, target)
		} else {
			return searchRangeInRange(nums, left, mid-1, target)
		}

	}
	return result
}

func searchFirst(nums []int, left, right int, target int) int {

	for left <= right {
		if right-left < 5 {
			result := linearSearchRange(nums, left, right, target)
			return result[0]
		}

		mid := (left + right) / 2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func searchLast(nums []int, left, right int, target int) int {

	for left <= right {
		if right-left < 5 {
			result := linearSearchRange(nums, left, right, target)
			return result[1]
		}

		mid := (left + right) / 2

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return right
}

func linearSearchRange(nums []int, left, right int, target int) []int {

	var result = []int{-1, -1}

	for left <= right {
		if nums[left] == target {
			result[0] = left
		} else {
			left++
		}

		if nums[right] == target {
			result[1] = right
		} else {
			right--
		}

		if result[0] != -1 && result[1] != -1 {
			break
		}
	}
	return result
}
