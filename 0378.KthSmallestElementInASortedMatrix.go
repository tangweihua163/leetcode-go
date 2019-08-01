package leetcode_go

func kthSmallest(matrix [][]int, k int) int {
	l := len(matrix)
	max := matrix[l-1][l-1]
	for i := 1; i <= k-1; i++ {
		matrix[0][0] = max
		adjust(matrix, 0, 0, k)
	}
	return matrix[0][0]
}

func adjust(matrix [][]int, i, j, k int) {

	l := len(matrix)

	for !(i <= j && i*l+j >= k || i >= j && j*l+i >= k) {

		if j < l-1 && i == l-1 {
			if matrix[i][j] > matrix[i][j+1] {
				matrix[i][j], matrix[i][j+1] = matrix[i][j+1], matrix[i][j]
				j++
				continue
			}
			break
		}

		if i < l-1 && j == l-1 {
			if matrix[i][j] > matrix[i+1][j] {
				matrix[i][j], matrix[i+1][j] = matrix[i+1][j], matrix[i][j]
				i++
				continue
			}
			break
		}

		if i < l-1 && j < l-1 {
			if matrix[i][j] > matrix[i][j+1] && matrix[i+1][j] > matrix[i][j+1] {
				matrix[i][j], matrix[i][j+1] = matrix[i][j+1], matrix[i][j]
				j++
				continue
			} else if matrix[i][j] > matrix[i+1][j] {
				matrix[i][j], matrix[i+1][j] = matrix[i+1][j], matrix[i][j]
				i++
				continue
			}
			break
		}
	}
}
