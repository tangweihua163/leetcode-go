package leetcode_go

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	var m, n = len(nums1), len(nums2)
	var left, right = (m + n + 1) / 2, (m + n + 2) / 2

	return float64(findKth(nums1, 0, nums2, 0, left)+findKth(nums1, 0, nums2, 0, right)) / 2.0
}

func findKth(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}

	if j >= len(nums2) {
		return nums1[i+k-1]
	}

	if k == 1 {
		return min(nums1[i], nums2[j])
	}

	var mid1, mid2 = math.MaxInt64, math.MinInt64
	if i+k/2-1 < len(nums1) {
		mid1 = nums1[i+k/2-1]
	}
	if j+k/2-1 < len(nums2) {
		mid2 = nums2[j+k/2-1]
	}

	if mid1 < mid2 {
		return findKth(nums1, i+k/2, nums2, j, k-k/2)
	} else {
		return findKth(nums1, i, nums2, j+k/2, k-k/2)
	}
}

/*

class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int m = nums1.size(), n = nums2.size(), left = (m + n + 1) / 2, right = (m + n + 2) / 2;
        return (findKth(nums1, 0, nums2, 0, left) + findKth(nums1, 0, nums2, 0, right)) / 2.0;
    }
    int findKth(vector<int>& nums1, int i, vector<int>& nums2, int j, int k) {
        if (i >= nums1.size()) return nums2[j + k - 1];
        if (j >= nums2.size()) return nums1[i + k - 1];
        if (k == 1) return min(nums1[i], nums2[j]);
        int midVal1 = (i + k / 2 - 1 < nums1.size()) ? nums1[i + k / 2 - 1] : INT_MAX;
        int midVal2 = (j + k / 2 - 1 < nums2.size()) ? nums2[j + k / 2 - 1] : INT_MAX;
        if (midVal1 < midVal2) {
            return findKth(nums1, i + k / 2, nums2, j, k - k / 2);
        } else {
            return findKth(nums1, i, nums2, j + k / 2, k - k / 2);
        }
    }
};

*/
