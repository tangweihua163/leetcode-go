package leetcode_go

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	//
	fmt.Println(trap([]int{2, 1, 0, 2}))
}
