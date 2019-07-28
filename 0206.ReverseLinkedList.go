package leetcode_go

func reverseList(head *ListNode) *ListNode {
	var list *ListNode
	for head != nil {
		p := head
		head = head.Next

		p.Next = list
		list = p
	}
	return list
}
