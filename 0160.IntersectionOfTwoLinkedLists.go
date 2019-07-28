package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	countA, tailA := countAndTail(headA)
	countB, tailB := countAndTail(headB)

	if tailA != tailB {
		return nil
	}

	if countA > countB {
		n := countA - countB
		for n > 0 {
			headA = headA.Next
			n--
		}
	} else {
		n := countB - countA
		for n > 0 {
			headB = headB.Next
			n--
		}
	}

	for headB != headA {
		headB = headB.Next
		headA = headA.Next
	}

	return headA
}

func countAndTail(head *ListNode) (count int, tail *ListNode) {
	if head == nil {
		return 0, nil
	}

	count = 1
	tail = head

	for tail.Next != nil {
		count++
		tail = tail.Next
	}

	return
}
