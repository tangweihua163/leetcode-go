package leetcode_go

func findTheDifference(s string, t string) byte {

	var result = 0

	for _, v := range s {
		result ^= int(v)
	}

	for _, v := range t {
		result ^= int(v)
	}

	return byte(result)
}
