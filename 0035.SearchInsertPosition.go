package leetcode_go

func searchInsert(nums []int, target int) int {
	var low = 0
	var high = len(nums)

	for low < high {
		mid := (low + high) / 2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] >= target {
			high = mid
		}
	}

	return low
}
