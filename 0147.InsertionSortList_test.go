package leetcode_go

import (
	"fmt"
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	var list *ListNode

	list = &ListNode{1, &ListNode{4, &ListNode{9, &ListNode{2, &ListNode{1, nil}}}}}

	list = insertionSortList(list)

	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
