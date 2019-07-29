package leetcode_go

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	maxArea := 0
	left := 0
	right := len(height) - 1
	width := len(height) - 1

	for left < right {
		maxArea = getMax(maxArea, getMin(height[left], height[right])*width)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		width--
	}

	return maxArea
}

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
