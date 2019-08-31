package leetcode_go

import "math"

func coinChange(coins []int, amount int) int {
	var m = make(map[int]int)
	return change(coins, amount, m)
}

func change(coins []int, amount int, m map[int]int) int {

	if v, ok := m[amount]; ok {
		return v
	}

	var min = math.MaxInt32
	var ex = false
	for _, coin := range coins {
		if amount < coin {
			continue
		}
		r := change(coins, amount-coin, m)
		m[amount-coin] = r
		if r != -1 && r < min {
			min = r
			ex = true
		}
	}

	if ex {
		m[amount] = min + 1
		return min + 1
	} else {
		m[amount] = -1
		return -1
	}

}
