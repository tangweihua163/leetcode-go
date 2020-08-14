package leetcode_go

func rotate1(nums []int, k int) {

	k = k % len(nums)

	reverse1(nums[:len(nums)-k])
	reverse1(nums[len(nums)-k:])

	reverse1(nums)
}

func reverse1(nums []int) {
	if len(nums) < 2 {
		return
	}

	var left = 0
	var right = len(nums) - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
