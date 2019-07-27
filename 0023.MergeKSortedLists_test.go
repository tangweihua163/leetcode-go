package leetcode_go

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {

	var list *ListNode

	list = mergeKLists([]*ListNode{
		{-1, &ListNode{1, nil}},
		{-3, &ListNode{1, &ListNode{4, nil}}},
		{-2, &ListNode{-1, &ListNode{0, &ListNode{2, nil}}}},
	})

	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
