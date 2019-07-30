package leetcode_go

func permute(nums []int) [][]int {

	if len(nums) == 0 {
		return [][]int{{}}
	}

	result := make([][]int, 0)

	result = perm(nums, 0, result)

	return result
}

func perm(nums []int, start int, result [][]int) [][]int {
	if start >= len(nums) {
		return result
	}

	if start == len(nums)-1 {
		sli := make([]int, 0, len(nums))
		sli = append(sli, nums...)
		result = append(result, sli)
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		result = perm(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}

	return result
}
