package leetcode_go

import "math"

func myAtoi(str string) int {

	if len(str) == 0 {
		return 0
	}

	var sign = 1
	var base = 0
	var i int
	var n int = len(str)

	// 忽略前导空格
	for i < n && str[i] == ' ' {
		i++
	}

	// 确定符号
	if i < n && (str[i] == '-' || str[i] == '+') {
		if str[i] == '-' {
			sign = -1
		}
		i++
	}

	//扫描有效数字
	for i < n && str[i] >= '0' && str[i] <= '9' {
		if base > math.MaxInt32/10 || base == math.MaxInt32/10 && (str[i]-'0') > 7 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		base = 10*base + int(str[i]-'0')

		i++
	}

	return base * sign
}
