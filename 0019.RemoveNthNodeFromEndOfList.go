package leetcode_go

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}

	if n <= 0 {
		return head
	}

	var pi = head
	var pj = head
	var length int
	for pi != nil {
		length++
		pi = pi.Next
	}

	if n > length {
		return head
	}

	if n == length {
		result := head.Next
		head.Next = nil
		return result
	}

	pi = head
	for n > 0 {
		pi = pi.Next
		n--
	}

	for pi.Next != nil {
		pi = pi.Next
		pj = pj.Next
	}

	tmp := pj.Next
	pj.Next = pj.Next.Next
	tmp.Next = nil

	return head

}
