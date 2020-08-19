package leetcode_go

func canCompleteCircuit(gas []int, cost []int) int {

	for i := 0; i < len(gas); i++ {
		if judge(gas, cost, i) {
			return i
		}
	}

	return -1
}

func judge(gas, cost []int, index int) bool {

	var n = 0
	var i = index
	for j := 0; j < len(gas); j++ {
		if n+gas[i] < cost[i] {
			return false
		}

		n = n + gas[i] - cost[i]

		i = (i + 1) % len(gas)
	}

	return true
}
