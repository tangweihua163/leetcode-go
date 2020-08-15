package leetcode_go

func largestRectangleArea(heights []int) int {

	var l = new(int)
	largest(heights, l)
	return *l

}

func largest(height []int, l *int) {

	if len(height) == 0 {
		return
	}

	var index, min int
	min = height[0]
	for i, h := range height {
		if h <= min {
			min = h
			index = i
		}
	}

	if len(height)*min > *l {
		*l = len(height) * min
	}

	largest(height[:index], l)
	largest(height[index:], l)
}
