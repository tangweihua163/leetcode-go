package leetcode_go

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {

	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	list = removeNthFromEnd(list, 2)

	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
