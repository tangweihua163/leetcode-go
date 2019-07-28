package leetcode_go

func insertionSortList(head *ListNode) *ListNode {

	var resultList, p *ListNode

	for head != nil {
		head, p = deleteHead(head)
		resultList = insert(resultList, p)
	}

	return resultList
}

func deleteHead(list *ListNode) (resultList *ListNode, headNode *ListNode) {
	if list == nil {
		return nil, nil
	}

	headNode = list

	resultList = list.Next

	headNode.Next = nil

	return
}

func insert(list *ListNode, node *ListNode) *ListNode {
	if list == nil {
		return node
	}

	if node == nil {
		return list
	}

	if node.Val <= list.Val {
		node.Next = list
		return node
	}

	var front, back *ListNode

	front = list.Next
	back = list

	for front != nil && back != nil {
		if node.Val > front.Val {
			front = front.Next
			back = back.Next
			continue
		}

		node.Next = front
		back.Next = node
		break
	}

	if front == nil {
		back.Next = node
	}

	return list

}
