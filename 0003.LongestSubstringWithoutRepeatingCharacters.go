package leetcode_go

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	var start, end int
	var maxlength = end - start + 1
	set := make(map[byte]int)
	set[s[start]] = 1

	for i := 1; i < len(s); i++ {
		if set[s[i]] == 0 {
			end = i
			set[s[end]] = 1
			if maxlength < end-start+1 {
				maxlength = end - start + 1
			}
			continue
		}

		for s[start] != s[i] {
			set[s[start]] = 0
			start++
		}

		start++

	}

	return maxlength

}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}

	var start, end int
	var maxlength = end - start + 1

	var bytes [256]byte

	bytes[s[start]] = 1

	for i := 1; i < len(s); i++ {
		if bytes[s[i]] == 0 {
			end = i
			bytes[s[end]] = 1
			if maxlength < end-start+1 {
				maxlength = end - start + 1
			}
			continue
		}

		for s[start] != s[i] {
			bytes[s[start]] = 0
			start++
		}

		start++

	}

	return maxlength
}
