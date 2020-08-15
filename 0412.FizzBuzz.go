package leetcode_go

import "strconv"

func fizzBuzz(n int) []string {

	var ss = make([]string, 0, n)

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			ss = append(ss, "FizzBuzz")
		} else if i%3 == 0 {
			ss = append(ss, "Fizz")
		} else if i%5 == 0 {
			ss = append(ss, "Buzz")
		} else {
			ss = append(ss, strconv.Itoa(i))
		}
	}
	return ss
}
