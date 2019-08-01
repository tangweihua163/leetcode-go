package leetcode_go

func topKFrequent(nums []int, k int) []int {

	var result = make([]int, 0, k)
	var count = make(map[int]int)

	for _, v := range nums {
		count[v]++
	}

	var heap [][]int
	for key, value := range count {
		heap = append(heap, []int{value, key})
	}

	sliceBuildMinHeap(heap[:k])

	for i := k; i < len(heap); i++ {
		if heap[i][0] > heap[0][0] {
			heap[0] = heap[i]
			sliceAdjustDownMinHeap(heap[:k], 0)
		}
	}

	heap = heap[:k]

	for i := 0; i < k; i++ {
		result = append(result, heap[0][1])
		heap[0] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		sliceAdjustDownMinHeap(heap, 0)
	}

	left := 0
	right := len(result) - 1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}

	return result
}
