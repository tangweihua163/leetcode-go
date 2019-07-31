package leetcode_go

func reverseKGroup(head *ListNode, k int) *ListNode {

	var list, tail *ListNode

	for head != nil {

		var n int
		var h, t *ListNode

		for head != nil && n != k {
			p := head
			head = head.Next
			p.Next = nil

			if h == nil {
				h = p
				t = p
				n++
			} else {
				p.Next = h
				h = p
				n++
			}
		}

		if n != k {
			h = reverseList(h)
		}

		if list == nil {
			list = h
			tail = t
		} else {
			tail.Next = h
			tail = t
		}

	}

	return list
}
