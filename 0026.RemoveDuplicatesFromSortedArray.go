package leetcode_go

func removeDuplicates(nums []int) int {

	index := 0

	var val = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != val {
			index++
			nums[index] = nums[i]
			val = nums[i]
		}
	}

	return index + 1
}
