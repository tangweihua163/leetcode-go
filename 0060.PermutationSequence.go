package leetcode_go

func getPermutation(n int, k int) string {
	var cal = 1
	for i := 1; i <= n; i++ {
		cal *= i
	}

	var a = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for i := 0; i < n-1; i++ {

		cal = cal / (n - i)

		j := k / cal
		if k%cal == 0 {
			j -= 1
		}

		k = k - j*cal

		a[i], a[j+i] = a[j+i], a[i]

		for l := j + i; l > i+1; l-- {
			if a[l] < a[l-1] {
				a[l], a[l-1] = a[l-1], a[l]
			}
		}
	}

	return string(a[:n])
}
