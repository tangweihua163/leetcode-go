package leetcode_go

func findKthBit(n int, k int) byte {

	var t = 0
	var m = 0
	for i := 1; i <= n; i++ {
		m = 2*m + 1
	}

	for true {

		if n == 1 {
			break
		}

		if k == m/2+1 {
			if t%2 == 0 {
				return '1'
			} else {
				return '0'
			}
		}

		if k <= m/2 {
			m = m / 2
			n--
			continue
		}

		t = t + 1
		k = m - k + 1
		m = m / 2
		n--
	}

	if t%2 == 0 {
		return '0'
	} else {
		return '1'
	}
}
