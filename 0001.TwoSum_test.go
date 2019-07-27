package leetcode_go

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

	fmt.Println(twoSum([]int{7, 2, 11, 15}, 9))

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 10))
}
