package leetcode_go

func trap(height []int) int {

	if len(height) <= 2 {
		return 0
	}

	last := make([]int, len(height))
	sum := 0
	last[0] = -1
	for i := 1; i < len(height); i++ {

		if height[i-1] == height[i] {
			last[i] = last[i-1]
			continue
		}

		if height[i-1] < height[i] {
			l := last[i-1]
			ll := i - 1
			for l != -1 {
				if height[l] < height[i] {
					sum += (height[l] - height[ll]) * (i - l - 1)
					ll = l
					l = last[l]
					continue
				}
				if height[l] == height[i] {
					sum += (height[l] - height[ll]) * (i - l - 1)
					last[i] = last[l]
					break
				}
				if height[l] > height[i] {
					sum += (height[i] - height[ll]) * (i - l - 1)
					last[i] = l
					break
				}
			}

			if l == -1 {
				last[i] = -1
			}
		}

		if height[i-1] > height[i] {
			last[i] = i - 1
		}
	}

	return sum
}

func trap2(height []int) int {

	// 1 至少三个才能装到水
	if len(height) < 3 {
		return 0
	}

	// 2 统计黑色总面积、最高高度
	var block = 0
	var maxHeight = 0
	for _, h := range height {
		if h > maxHeight {
			maxHeight = h
		}
		block += h
	}

	// 3 统计装满水后空白部分的高度
	var l = 0
	var lh = 0
	var r = len(height) - 1
	var rh = 0
	var space = 0

	// 3.1 从左边扫描空白
	for l < r && height[l] < maxHeight {

		if height[l] > lh {
			lh = height[l]
		}

		space = space + (maxHeight - lh)

		l++
	}

	// 3.2 从右边扫描空白
	for l < r && height[r] < maxHeight {

		if height[r] > rh {
			rh = height[r]
		}

		space = space + (maxHeight - rh)

		r--
	}

	// 4 总面积 - 黑色面积 - 空白面积 = 水的面积
	return maxHeight*len(height) - space - block

}
