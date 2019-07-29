package leetcode_go

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 1, 4, 5, 7, 8, 6, 3, 2}

	nextPermutation(nums)

	fmt.Println(nums)
}
