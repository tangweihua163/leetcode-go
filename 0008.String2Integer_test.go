package leetcode_go

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {

	fmt.Println(myAtoi(""))
	fmt.Println(myAtoi("1"))
	fmt.Println(myAtoi("11112"))
	fmt.Println(myAtoi("333"))
	fmt.Println(myAtoi("33333333333333"))
	fmt.Println(myAtoi("2147483647"))
	fmt.Println(myAtoi("2147483648"))
	fmt.Println(myAtoi("-2147483648"))
	fmt.Println(myAtoi("-2147483649"))

}
