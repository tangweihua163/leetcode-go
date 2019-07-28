package leetcode_go

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 2, 4, 1}))
	fmt.Println(maxProfit([]int{1}))
	fmt.Println(maxProfit([]int{1, 2}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{6, 5, 4, 3, 4, 5, 6, 5, 4, 3}))

	fmt.Println(maxProfit([]int{0, 1, -1, 0}))
}
