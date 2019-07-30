package leetcode_go

import (
	"fmt"
	"testing"
)

func TestRotateImage(t *testing.T) {
	var a = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	rotate(a)
	fmt.Println(a)
}
