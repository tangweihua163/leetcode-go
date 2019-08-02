package leetcode_go

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {
	var p *ListNode

	p = &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	sortList(p)
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}

	p = &ListNode{3, &ListNode{4, nil}}
	sortList(p)
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
