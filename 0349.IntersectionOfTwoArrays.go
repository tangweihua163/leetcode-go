package leetcode_go

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, n := range nums1 {
		m1[n]++
	}

	for _, n := range nums2 {
		m2[n]++
	}

	var result = make([]int, 0, len(m1))
	for k := range m1 {
		if _, ok := m2[k]; ok {
			result = append(result, k)
		}
	}

	return result
}
