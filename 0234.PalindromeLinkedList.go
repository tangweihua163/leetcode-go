package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

	length := lengthOfLinkedList(head)

	if length <= 1 {
		return true
	}

	mid := (length - 1) / 2

	p := head
	for mid != 0 {
		p = p.Next
		mid--
	}
	q := p.Next
	p.Next = nil

	q = reverseLinkedList(q)

	return equalPrefix(head, q)
}

func lengthOfLinkedList(head *ListNode) (length int) {
	for head != nil {
		length++
		head = head.Next
	}
	return
}

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var list *ListNode

	for head != nil {
		p := head
		head = head.Next
		p.Next = list
		list = p
	}

	return list
}

func equalPrefix(head1, head2 *ListNode) bool {
	for head1 != nil && head2 != nil {
		if head1.Val != head2.Val {
			return false
		}
		head1 = head1.Next
		head2 = head2.Next
	}

	return true
}
