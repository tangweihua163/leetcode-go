package leetcode_go

func twoSumII(numbers []int, target int) []int {

	lo := 0
	hi := len(numbers) - 1

	for lo < hi {

		if numbers[lo]+numbers[hi] == target {
			return []int{lo + 1, hi + 1}
		}

		if numbers[lo]+numbers[hi] > target {
			hi--
		}

		if numbers[lo]+numbers[hi] < target {
			lo++
		}
	}

	return []int{0, 0}
}
