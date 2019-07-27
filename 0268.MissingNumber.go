package leetcode_go

func missingNumber(nums []int) int {

	l := len(nums)

	sum := (l * (l + 1)) >> 1

	for _, v := range nums {
		sum -= v
	}

	return sum
}
