package leetcode_go

import (
	"fmt"
	"testing"
)

func TestRotateArray1(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	rotate1(s, 0)
	fmt.Println(s)

	s = []int{1, 2, 3, 4, 5, 6, 7}
	rotate1(s, 1)
	fmt.Println(s)

	s = []int{1, 2, 3, 4, 5, 6, 7}
	rotate1(s, 7)
	fmt.Println(s)

	s = []int{1, 2, 3, 4, 5, 6, 8}
	rotate1(s, 0)
	fmt.Println(s)
}
