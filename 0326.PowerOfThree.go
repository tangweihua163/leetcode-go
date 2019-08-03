package leetcode_go

func isPowerOfThree(n int) bool {

	multi := make([]int, 40)

	if n == 1 {
		return true
	}
	for n%3 != 0 {
		return false
	}

	multi[0] = 1
	multi[1] = 3
	multi[2] = 9
	multi[3] = 27
	i := 1
	for {
		pi := pow3(multi, i)
		if pi == n {
			return true
		} else if pi > n {
			start, end := i>>1, i
			return searchPow3(start, end, multi, n)
		} else {
			i <<= 1
		}
	}

}

func pow3(multi []int, i int) int {
	if multi[i] != 0 {
		return multi[i]
	} else {
		p := pow3(multi, i>>1) * pow3(multi, i-i>>1)
		multi[i] = p
		return p
	}
}

func searchPow3(start, end int, multi []int, n int) bool {
	for start < end {
		if end-start <= 1 {
			return false
		}
		mid := (start + end) >> 1
		pm := pow3(multi, mid)
		if pm == n {
			return true
		} else if pm < n {
			start = mid
		} else {
			end = mid
		}
	}
	return false
}
