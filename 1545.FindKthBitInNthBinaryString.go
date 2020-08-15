package leetcode_go

func findKthBit(n int, k int) byte {

	var a byte = '0'
	var b byte = '1'

	s := [][]byte{{}, []byte("0"), []byte("011")}

	var t int
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
		return s[n][k-1]
	} else {
		if s[n][k-1] == a {
			return b
		} else {
			return a
		}
	}
}
