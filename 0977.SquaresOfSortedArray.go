package leetcode_go

import (
	"math"
)

func sortedSquares(A []int) []int {

	result := make([]int, 0, len(A))

	min := math.MaxInt64
	minIndex := -1

	low := 0
	high := len(A) - 1

	for low <= high {
		mid := (low + high) / 2

		if A[mid] == 0 {
			min = 0
			minIndex = mid
			break
		}

		multi := A[mid] * A[mid]
		if multi < min {
			min = multi
			minIndex = mid
		}

		if A[mid] > 0 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	i := minIndex - 1
	j := minIndex + 1

	result = append(result, min)

	for i > -1 || j < len(A) {
		if i == -1 {
			result = append(result, A[j]*A[j])
			j++
			continue
		}

		if j == len(A) {
			result = append(result, A[i]*A[i])
			i--
			continue
		}

		if A[i]+A[j] >= 0 {
			result = append(result, A[i]*A[i])
			i--
			continue
		} else {
			result = append(result, A[j]*A[j])
			j++
			continue
		}

	}

	return result

}
