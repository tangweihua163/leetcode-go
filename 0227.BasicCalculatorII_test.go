package leetcode_go

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	fmt.Println(calculate("3+2*2+3"))
	fmt.Println(calculate("3/2"))
	fmt.Println(calculate("3+5 / 2"))
	fmt.Println(calculate("3+22*2"))
	fmt.Println(calculate("3+2*2-2"))

	fmt.Println(_calculate("3+2*2+3"))
	fmt.Println(_calculate("3/2"))
	fmt.Println(_calculate("3+5 / 2"))
	fmt.Println(_calculate("3+22*2"))
	fmt.Println(_calculate("3+2*2-2"))
}
