package leetcode_go

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	var s = []int{1, 2, 5}
	fmt.Println(coinChange(s, 11))

	var ss = []int{2}
	fmt.Println(coinChange(ss, 3))
}
