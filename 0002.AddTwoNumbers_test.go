package leetcode_go

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	l1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}

	l2 := &ListNode{5, &ListNode{7, &ListNode{0, &ListNode{5, &ListNode{0, nil}}}}}

	l := addTwoNumbers(l1, l2)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}
