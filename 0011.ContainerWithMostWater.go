package leetcode_go

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max := 0
	left := 0
	right := len(height) - 1
	width := len(height) - 1

	for left < right {
		min, _ := MinInt(height[left], height[right])
		max, _ = MaxInt(max, min*width)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		width--
	}

	return max
}
