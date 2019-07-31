package leetcode_go

import (
	"fmt"
	"testing"
)

func TestMinInt(t *testing.T) {
	fmt.Println(MinInt([]int{1, 2, 3, 4, 5}...))
}

func TestMaxInt(t *testing.T) {
	fmt.Println(MaxInt([]int{197, 37192, 48, 349, 712}...))
}

func TestHeapSortASC(t *testing.T) {
	A := []int{25, 346, 65, 52, 23, 5, 452, 23, 434224}
	HeapSortASC(A)
	fmt.Println(A)
}

func TestHeapSortDESC(t *testing.T) {
	A := []int{25, 346, 65, 52, 23, 5, 452, 23, 434224}
	HeapSortDESC(A)
	fmt.Println(A)
}

func TestQuickSort(t *testing.T) {
	A := []int{25, 346, 65, 52, 23, 5, 452, 23, 434224}
	QuickSort(A)
	fmt.Println(A)
}
