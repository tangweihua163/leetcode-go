package leetcode_go

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{1, &ListNode{5, &ListNode{6, &ListNode{9, nil}}}}
	l2 := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{10, &ListNode{12, nil}}}}}

	l := mergeTwoLists(l1, l2)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
