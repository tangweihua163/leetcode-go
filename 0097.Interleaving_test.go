package leetcode_go

import (
	"fmt"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	s := []string{
		"bab",
		"abbbcba",
		"babbbabcba",
	}
	fmt.Println(isInterleave(s[0], s[1], s[2]))

	fmt.Println()

	s = []string{
		"aabcc", "dbbca", "aadbbcbcac",
	}
	fmt.Println(isInterleave(s[0], s[1], s[2]))

}
