package leetcode_go

func findKthLargest(nums []int, k int) int {

	buildMinHeap(nums[:k])

	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			adjustDownMinHeap(nums[:k], 0)
		}
	}

	return nums[0]
}

func buildMinHeap(array []int) {
	start := (len(array) - 1) / 2

	for i := start; i >= 0; i-- {
		adjustDownMinHeap(array, i)
	}
}

func adjustDownMinHeap(array []int, i int) {

	if i > len(array)/2-1 {
		return
	}

	l := len(array)
	left := i*2 + 1
	right := i*2 + 2

	if right < l {
		if array[i] > array[right] && array[left] >= array[right] {
			array[i], array[right] = array[right], array[i]
			adjustDownMinHeap(array, right)
		} else if array[i] > array[left] && array[right] >= array[left] {
			array[i], array[left] = array[left], array[i]
			adjustDownMinHeap(array, left)
		}
	} else {
		if array[i] > array[left] {
			array[i], array[left] = array[left], array[i]
		}
	}

}
