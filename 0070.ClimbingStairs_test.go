package leetcode_go

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	fmt.Println(climbStairs(1) == 1)
	fmt.Println(climbStairs(2) == 2)
	fmt.Println(climbStairs(3) == 3)
	fmt.Println(climbStairs(4) == 5)
	fmt.Println(climbStairs(5) == 8)
	fmt.Println(climbStairs(6) == 13)
	fmt.Println(climbStairs(7) == 21)
}
