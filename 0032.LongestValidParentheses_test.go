package leetcode_go

import (
	"fmt"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	//fmt.Println(longestValidParentheses("())")==2)
	//fmt.Println(longestValidParentheses("(())")==4)
	//fmt.Println(longestValidParentheses("()()")==4)
	//fmt.Println(longestValidParentheses(")()(")==2)

	fmt.Println(longestValidParentheses("(()()((()())"), 6)
}
