package leetcode_go

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	fmt.Println(canJump([]int{2, 3, 1, 3, 4, 5, 4, 5, 3, 6, 2, 10, 2, 7, 6, 5, 4, 3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{2, 3, 1, 3, 4, 5, 4, 5, 3, 6, 2, 9, 2, 7, 6, 5, 4, 3, 2, 1, 0, 4}))
}
