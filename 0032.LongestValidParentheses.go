package leetcode_go

func longestValidParentheses(s string) int {

	A := make([]int, 0, len(s))

	for _, v := range s {
		if v == int32('(') {
			A = append(A, -1)
		} else {
			A = append(A, -2)
		}

	}

	for i := 0; i < len(A); i++ {

		if A[i] == -2 {

			if i <= 0 {
				continue
			}

			if A[i-1] == -1 {

				A[i] = 2

				last := i - A[i]
				for last >= 0 && A[last] != -1 && A[last] != -2 {
					A[i] += A[last]
					last = i - A[i]
				}

				continue
			}

			if A[i-1] != -2 && A[i-1] != -1 {
				last := i - 1 - A[i-1]
				if last >= 0 && A[last] == -1 {
					A[i] = A[i-1] + 2
					last = i - A[i]
				}

				for last >= 0 && A[last] != -1 && A[last] != -2 {
					A[i] += A[last]
					last = i - A[i]
				}
			}

		}
	}

	max := 0
	for _, v := range A {
		if max < v {
			max = v
		}
	}

	return max

}
