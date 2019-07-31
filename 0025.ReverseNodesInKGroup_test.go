package leetcode_go

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	l = reverseKGroup(l, 2)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
