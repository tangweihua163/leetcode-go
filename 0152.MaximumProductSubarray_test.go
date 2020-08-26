package leetcode_go

import (
	"fmt"
	"testing"
)

func TestMaximumProductSubarray(t *testing.T) {
	fmt.Println(maxProduct([]int{-2, 1, 0, -3}))
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}
