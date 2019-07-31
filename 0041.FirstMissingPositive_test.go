package leetcode_go

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	nums := []int{2, 3, 0, 1}

	fmt.Println(partition(nums))
}

func TestFirstMissingPositive(t *testing.T) {
	nums := []int{-1, 3, 52}
	fmt.Println(firstMissingPositive(nums))
}
