package leetcode_go

func mergeKLists(lists []*ListNode) *ListNode {

	k := len(lists)

	heap := make([]*ListNode, k)

	var j int
	for _, l := range lists {
		if l == nil {
			k--
			continue
		}
		heap[j] = l
		j++
	}

	heap = heap[:k]

	buildHeap(heap)

	var list, tail *ListNode

	for len(heap) > 0 {

		p := heap[0]

		if tail == nil {
			list = p
			tail = p
			p = p.Next
			tail.Next = nil
		} else {
			tail.Next = p
			tail = tail.Next
			p = p.Next
			tail.Next = nil
		}

		if p == nil {
			heap[0] = heap[len(heap)-1]
			heap = heap[:len(heap)-1]
		} else {
			heap[0] = p
		}

		adjustDown(heap, 0)

	}

	return list

}

func buildHeap(array []*ListNode) {
	start := (len(array) - 1) / 2

	for i := start; i >= 0; i-- {
		adjustDown(array, i)
	}
}

func adjustDown(array []*ListNode, i int) {

	if i > len(array)/2-1 {
		return
	}

	l := len(array)
	left := i*2 + 1
	right := i*2 + 2

	if right < l {
		if array[i].Val > array[right].Val && array[left].Val >= array[right].Val {
			array[i], array[right] = array[right], array[i]
			adjustDown(array, right)
		} else if array[i].Val > array[left].Val && array[right].Val > array[left].Val {
			array[i], array[left] = array[left], array[i]
			adjustDown(array, left)
		}
	} else {
		if array[i].Val > array[left].Val {
			array[i], array[left] = array[left], array[i]
		}
	}

}
