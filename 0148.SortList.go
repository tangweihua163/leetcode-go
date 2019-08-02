package leetcode_go

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	sortSingleLinkedList(head, nil)
	return head
}

func sortSingleLinkedList(left, right *ListNode) {
	if left == right || left.Next == right {
		return
	}
	mid := partitionSingleLinkedList(left, right)

	if mid == left || mid == right {
		return
	}

	sortSingleLinkedList(left, mid)
	sortSingleLinkedList(mid, right)
	return
}

func partitionSingleLinkedList(left, right *ListNode) *ListNode {
	var back, front = left, left.Next
	mv := back.Val
	for front != right {
		if front.Val < mv {
			front.Val, back.Val = back.Val, front.Val
			back = back.Next
		}
		front = front.Next
	}
	return back
}
