package leetcode_go

func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	p := head
	q := head

	for p.Next != nil && q.Next != nil && q.Next.Next != nil {
		p = p.Next
		q = q.Next.Next

		if p == q {
			return true
		}
	}

	return false

}
