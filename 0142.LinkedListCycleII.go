package leetcode_go

func detectCycle(head *ListNode) *ListNode {

	mark := make(map[*ListNode]int)

	for head != nil {
		if _, ok := mark[head]; ok {
			return head
		} else {
			mark[head] = 1
		}

		head = head.Next
	}

	return nil

}
