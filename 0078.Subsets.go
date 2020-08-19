package leetcode_go

func subsets(nums []int) [][]int {
	var result [][]int
	var set = make([]int, len(nums))
	return fff(set, 0, nums, result)
}

func fff(set []int, i int, nums []int, result [][]int) [][]int {

	if len(nums) == 0 {
		dst := make([]int, i)
		copy(dst, set[:i])
		result = append(result, dst)
		return result
	}

	set[i] = nums[0]
	result = fff(set, i+1, nums[1:], result)
	result = fff(set, i, nums[1:], result)

	return result
}
