package leetcode_go

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	println(lengthOfLongestSubstring("sglsdalfldoeruooosoosoosddslllssdsjfl") == 7)

	println(lengthOfLongestSubstring("s") == 1)

	println(lengthOfLongestSubstring("ss") == 1)

	println(lengthOfLongestSubstring("12345") == 5)

	println(lengthOfLongestSubstring("12344") == 4)
}
