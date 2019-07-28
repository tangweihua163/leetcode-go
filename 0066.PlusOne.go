package leetcode_go

func plusOne(digits []int) []int {

	c := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + c
		digits[i] = sum % 10
		c = sum / 10
	}

	if c > 0 {
		return append([]int{1}, digits...)
	}

	return digits
}
