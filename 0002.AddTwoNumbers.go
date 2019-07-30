package leetcode_go

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	c := 0

	var head, tail *ListNode

	var sum int

	for {

		sum = 0

		if l1 != nil && l2 != nil {
			sum = l1.Val + l2.Val + c
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 == nil && l2 == nil {
			if c != 0 {
				if tail != nil {
					tail.Next = &ListNode{c, nil}
				}
			}
			break
		} else if l1 == nil && l2 != nil {
			sum = l2.Val + c
			l2 = l2.Next
		} else if l2 == nil && l1 != nil {
			sum = l1.Val + c
			l1 = l1.Next
		}

		c = sum / 10

		p := &ListNode{sum % 10, nil}

		if head == nil {
			head = p
			tail = head
		} else if tail != nil {
			tail.Next = p
			tail = p
		}

	}

	var p *ListNode
	var iter = head
	for iter != nil {
		if iter.Val != 0 {
			p = iter
		}
		iter = iter.Next
	}

	if p == nil {
		return &ListNode{0, nil}
	} else {
		p.Next = nil
	}

	return head
}
