package leetcode_go

func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return head
	}

	if head.Next == nil {
		if head.Val == val {
			return nil
		}
		return head
	}

	front := head.Next
	back := head

	for front != nil {
		if front.Val == val {
			front = front.Next
			back.Next = front
			continue
		}
		front = front.Next
		back = back.Next
	}

	if head.Val == val {
		return head.Next
	}

	return head
}
