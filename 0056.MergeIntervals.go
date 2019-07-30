package leetcode_go

func mergeIntervals(intervals [][]int) [][]int {

	if len(intervals) <= 1 {
		return intervals
	}

	sliceBuildMinHeap(intervals)

	intvl := intervals
	i := len(intervals) - 1

	for len(intervals) > 0 {
		f := false
		for len(intervals) > 1 {
			sliceAdjustDownMinHeap(intervals, 1)
			sli, ok := mergeRange(intervals[0], intervals[1])
			if !ok {
				break
			}
			intervals[0] = sli
			f = true
			intervals[1] = intervals[len(intervals)-1]
			intervals = intervals[:len(intervals)-1]

		}

		for len(intervals) > 2 {
			sliceAdjustDownMinHeap(intervals, 2)
			sli, ok := mergeRange(intervals[0], intervals[2])
			if !ok {
				break
			}
			intervals[0] = sli
			f = true
			intervals[2] = intervals[len(intervals)-1]
			intervals = intervals[:len(intervals)-1]
		}

		if f {
			continue
		}

		if len(intervals)-1 == i {
			intervals[0], intvl[i] = intvl[i], intervals[0]
			intervals = intervals[:len(intervals)-1]
			i--
			sliceAdjustDownMinHeap(intervals, 0)
		} else {
			intvl[i] = intervals[0]
			intervals[0] = intervals[len(intervals)-1]
			i--
			intervals = intervals[:len(intervals)-1]
			sliceAdjustDownMinHeap(intervals, 0)
		}

	}

	left := 0
	right := len(intvl) - 1
	for left < right {
		intvl[left], intvl[right] = intvl[right], intvl[left]
		left++
		right--
	}

	return intvl[:len(intvl)-1-i]
}

func sliceBuildMinHeap(A [][]int) {
	start := (len(A) - 1) / 2
	for i := start; i >= 0; i-- {
		sliceAdjustDownMinHeap(A, i)
	}
}

func sliceAdjustDownMinHeap(A [][]int, i int) {
	if i > len(A)/2-1 {
		return
	}

	l := len(A)
	left := i*2 + 1
	right := i*2 + 2

	if right < l {
		if A[i][0] > A[right][0] && A[left][0] >= A[right][0] {
			A[i], A[right] = A[right], A[i]
			sliceAdjustDownMinHeap(A, right)
		} else if A[i][0] > A[left][0] && A[right][0] >= A[left][0] {
			A[i], A[left] = A[left], A[i]
			sliceAdjustDownMinHeap(A, left)
		}
	} else {
		if A[i][0] > A[left][0] {
			A[i], A[left] = A[left], A[i]
		}
	}
}

func mergeRange(a, b []int) (interval []int, suc bool) {
	if a[1]-b[0] < 0 || b[1] < a[0] {
		suc = false
		return
	}
	return []int{min(a[0], b[0]), max(a[1], b[1])}, true
}
