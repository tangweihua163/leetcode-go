package leetcode_go

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode

	for {

		if l1 == nil && l2 == nil {
			break
		}

		if l1 == nil {
			if tail == nil {
				return l2
			}
			tail.Next = l2
			break
		}

		if l2 == nil {
			if tail == nil {
				return l1
			}
			tail.Next = l1
			break
		}

		var p *ListNode
		if l1.Val < l2.Val {
			p = l1
			l1 = l1.Next
			p.Next = nil
		} else {
			p = l2
			l2 = l2.Next
			p.Next = nil
		}

		if head == nil {
			head = p
			tail = head
		} else if tail != nil {
			tail.Next = p
			tail = tail.Next
		}
	}

	return head
}
