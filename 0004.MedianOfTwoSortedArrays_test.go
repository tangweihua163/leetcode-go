package leetcode_go

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	var nums1 = []int{0, 3, 5, 6, 8, 9}
	var nums2 = []int{3, 6, 8, 9, 12, 54, 67, 89}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArrays(nums2, nums1))

	nums1 = []int{0, 3, 5, 6, 8, 9}
	nums2 = []int{13, 16, 18, 19, 22, 54, 67, 89}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{}
	nums2 = []int{13, 16, 18, 19, 22, 54, 67, 89}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
