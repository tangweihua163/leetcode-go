package leetcode_go

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement([]int{3, 2, 3}))

	fmt.Println(majorityElement1([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement1([]int{3, 2, 3}))
}
