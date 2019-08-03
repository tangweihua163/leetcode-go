package leetcode_go

import (
	"fmt"
	"testing"
)

func TestIsPowerOfThree(t *testing.T) {
	//fmt.Println(isPowerOfThree(0)==false)
	//fmt.Println(isPowerOfThree(1)==true)
	//fmt.Println(isPowerOfThree(2)==false)
	//fmt.Println(isPowerOfThree(3)==true)
	//fmt.Println(isPowerOfThree(8)==false)
	//fmt.Println(isPowerOfThree(81)==true)
	//fmt.Println(isPowerOfThree(87)==false)
	//fmt.Println(isPowerOfThree(243))
	//fmt.Println(isPowerOfThree(19683))

	sum := 1
	for i := 1; i <= 64; i++ {
		s := sum * 3
		fmt.Println(i, s)
		if s < sum {
			fmt.Println("----")
		}
		sum = s
	}
}
