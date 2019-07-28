package leetcode_go

func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	maxIndex := 0
	minIndex := 0

	max := 0
	for i, v := range prices {
		if v > prices[maxIndex] {
			maxIndex = i
			if v-prices[minIndex] > max {
				max = v - prices[minIndex]
			}
		} else if v < prices[minIndex] {
			minIndex = i
			maxIndex = i
		}
	}

	return max
}
