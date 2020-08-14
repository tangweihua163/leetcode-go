package leetcode_go

func getSum(a int, b int) int {
	var n = 1
	var i = 0

	var c = 0
	var s = 0

	for n != 0 {

		if a&(1<<i) == 0 && b&(1<<i) == 0 {
			if c == 1 {
				s = s | 1<<i
				c = 0
			}
		} else if a&(1<<i) == 1<<i && b&(1<<i) == 1<<i {
			if c == 1 {
				s = s | 1<<i
			} else {
				c = 1
			}
		} else {
			if c == 0 {
				s = s | 1<<i
			}
		}

		i++
		n = n << 1
	}

	return s
}
