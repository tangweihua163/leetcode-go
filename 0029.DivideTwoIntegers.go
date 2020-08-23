package leetcode_go

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	var flag = 1

	if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
		flag = 1
	} else if dividend < 0 {
		dividend = -dividend
		flag = -1
	} else if divisor < 0 {
		divisor = -divisor
		flag = -1
	}

	if dividend < divisor {
		return 0
	}

	q := quo(dividend, divisor)

	if flag < 0 {
		q = -q
	}

	if q > math.MaxInt32 {
		q = math.MaxInt32
	}

	return q
}

func quo(a, b int) int {
	if a < b {
		return 0
	}

	if a == b {
		return 1
	}

	var i int

	for i = 0; i < 32; i++ {
		if a > (b << i) {
			continue
		} else {
			break
		}
	}

	return 1<<(i-1) + quo(a-(b<<(i-1)), b)
}
