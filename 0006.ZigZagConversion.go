package leetcode_go

func convert1(s string, numRows int) string {

	var result = make([]byte, len(s))

	var a, b = numRows, 0

	for i := 0; i < numRows; i++ {
		if a == 0 || b == 0 {
			for j := i; j < len(s); j += numRows {
				result = append(result, s[j])
			}
		} else {
			for j := i; j < len(s); j += numRows {
				result = append(result, s[j])
				if j+a < len(s) {
					result = append(result, s[j+a])
				}
			}
		}
		a--
		b++
	}

	return string(result)
}
