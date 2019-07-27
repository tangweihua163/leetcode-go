package leetcode_go

import "strconv"

func countAndSay(n int) string {
	var old = "1"

	for i := 2; i <= n; i++ {
		old = generateNew(old)
	}

	return old
}

func generateNew(str string) string {

	var start int
	var result string

	bytes := []byte(result)
	i := 1

	for ; i < len(str); i++ {
		if str[i] == str[start] {
			continue
		}

		bytes = strconv.AppendInt(bytes, int64(i-start), 10)
		bytes = append(bytes, str[start])

		start = i
	}

	bytes = strconv.AppendInt(bytes, int64(i-start), 10)
	bytes = append(bytes, str[start])

	return string(bytes)

}
