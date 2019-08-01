package leetcode_go

func minCut(s string) int {

	var min = make([]int, len(s)*len(s))
	var palindromeRange = make([]bool, len(s)*len(s))

	if len(s) <= 1 {
		return 0
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return 0
		} else {
			return 1
		}
	}

	for k := 0; k < len(s); k++ {

		for i := 0; i+k < len(s); i++ {

			index := i*len(s) + i + k
			min[index] = len(s)

			if k == 0 {
				min[index] = 0
				palindromeRange[index] = true
			} else if k == 1 {
				if s[i] == s[i+k] {
					min[index] = 0
					palindromeRange[index] = true
				} else {
					min[index] = 1
					palindromeRange[index] = false
				}
			} else {
				if s[i] == s[i+k] && palindromeRange[(i+1)*len(s)+i+k-1] {
					palindromeRange[index] = true
					min[index] = 0
				} else {
					palindromeRange[index] = false
					for l := i; l < i+k; l++ {
						t := min[i*len(s)+l] + min[(l+1)*len(s)+i+k] + 1
						if min[index] > t {
							min[index] = t
						}
					}
				}
			}
		}
	}

	return min[len(s)-1]
}
