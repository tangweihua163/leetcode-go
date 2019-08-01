package leetcode_go

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	fmt.Println(kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8))
}
