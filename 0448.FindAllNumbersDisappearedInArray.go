package leetcode_go

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {

		if nums[i] == i+1 {
			continue
		}

		t := 0
		index := i
		for nums[index] != index+1 {
			t, nums[index] = nums[index], t
			index = t - 1
			if index < 0 {
				break
			}
		}

	}

	index := 0

	for i, v := range nums {
		if v != i+1 {
			nums[index] = i + 1
			index++
		}
	}

	return nums[:index]
}
