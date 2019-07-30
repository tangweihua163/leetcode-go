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
