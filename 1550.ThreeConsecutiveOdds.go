package leetcode_go

func threeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	for i := 0; i <= len(arr)-3; i++ {
		if arr[i]%2 == 0 {
			continue
		}

		if arr[i+1]%2 == 0 {
			i = i + 1
			continue
		}

		if arr[i+2]%2 == 0 {
			i = i + 2
			continue
		}

		return true
	}

	return false
}
