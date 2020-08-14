package leetcode_go

func generatePascalsTriangle(numRows int) [][]int {

	if numRows == 0 {
		return [][]int{}
	}

	var result = make([][]int, 0, numRows)

	var last = []int{1}
	result = append(result, last)

	for i := 1; i < numRows; i++ {
		var newRow = make([]int, len(last)+1)

		//fmt.Println(len(newRow))

		newRow[0] = 1

		var j int
		for j = 1; j < len(newRow)-1; j++ {
			newRow[j] = last[j-1] + last[j]
		}
		newRow[j] = 1

		result = append(result, newRow)

		last = newRow
	}

	return result
}
