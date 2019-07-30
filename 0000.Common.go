package leetcode_go

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MinInt(arr ...int) (val int, index int) {
	min := arr[0]
	minIndex := 0
	for i, v := range arr {
		if v < min {
			min = v
			minIndex = i
		}
	}
	return min, minIndex
}

func MaxInt(arr ...int) (val int, index int) {
	max := arr[0]
	maxIndex := 0
	for i, v := range arr {
		if v > max {
			max = v
			maxIndex = i
		}
	}
	return max, maxIndex
}

// 堆排序（升序）
func HeapSortASC(A []int) {
	BuildMaxHeap(A)
	for i := len(A) - 1; i > 0; i-- {
		A[i], A[0] = A[0], A[i]
		A = A[:i]
		AdjustDownMaxHeap(A, 0)
	}
}

// 构建最大堆
func BuildMaxHeap(A []int) {
	start := (len(A) - 1) / 2
	for i := start; i >= 0; i-- {
		AdjustDownMaxHeap(A, i)
	}
}

// 调整最大堆节点
func AdjustDownMaxHeap(A []int, i int) {
	if i > len(A)/2-1 {
		return
	}

	l := len(A)
	left := i*2 + 1
	right := i*2 + 2

	if right < l {
		if A[i] < A[right] && A[left] <= A[right] {
			A[i], A[right] = A[right], A[i]
			AdjustDownMaxHeap(A, right)
		} else if A[i] < A[left] && A[right] <= A[left] {
			A[i], A[left] = A[left], A[i]
			AdjustDownMaxHeap(A, left)
		}
	} else {
		if A[i] < A[left] {
			A[i], A[left] = A[left], A[i]
		}
	}
}

// 堆排序（降序）
func HeapSortDESC(A []int) {
	BuildMinHeap(A)
	for i := len(A) - 1; i > 0; i-- {
		A[i], A[0] = A[0], A[i]
		A = A[:i]
		AdjustDownMinHeap(A, 0)
	}
}

// 构建最小堆
func BuildMinHeap(A []int) {
	start := (len(A) - 1) / 2
	for i := start; i >= 0; i-- {
		AdjustDownMinHeap(A, i)
	}
}

// 调整最小堆节点
func AdjustDownMinHeap(A []int, i int) {
	if i > len(A)/2-1 {
		return
	}

	l := len(A)
	left := i*2 + 1
	right := i*2 + 2

	if right < l {
		if A[i] > A[right] && A[left] >= A[right] {
			A[i], A[right] = A[right], A[i]
			AdjustDownMinHeap(A, right)
		} else if A[i] > A[left] && A[right] >= A[left] {
			A[i], A[left] = A[left], A[i]
			AdjustDownMinHeap(A, left)
		}
	} else {
		if A[i] > A[left] {
			A[i], A[left] = A[left], A[i]
		}
	}
}
