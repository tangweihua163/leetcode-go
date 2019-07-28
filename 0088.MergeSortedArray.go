package leetcode_go

func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1
	i := m - 1
	j := n - 1

	for j >= 0 {
		if i < 0 || nums2[j] >= nums1[i] {
			nums1[index] = nums2[j]
			index--
			j--
		} else {
			nums1[index] = nums1[i]
			index--
			i--
		}
	}

}
