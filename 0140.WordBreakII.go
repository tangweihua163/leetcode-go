package leetcode_go

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	m := make(map[int][][]string)
	mm := make([]int, len(s))
	r, _ := find1(s, 0, wordDict, m, mm)

	result := make([]string, 0, len(r))
	for i := 0; i < len(r); i++ {
		result = append(result, strings.Join(r[i], " "))
	}

	fmt.Println(result)
	return result
}

func find1(s string, i int, d []string, m map[int][][]string, mm []int) ([][]string, bool) {
	if i == len(s) {
		return [][]string{}, true
	}

	if v := mm[i]; v != 0 {
		if v == 1 {
			return m[i], true
		} else {
			return [][]string{}, false
		}
	}

	var result [][]string
	var bb bool
	for j := 0; j < len(d); j++ {

		if len(s)-i < len(d[j]) {
			continue
		}

		if string(s[i:i+len(d[j])]) != d[j] {
			continue
		}

		tmp, b := find1(s, i+len(d[j]), d, m, mm)

		if b {
			if len(tmp) == 0 {
				result = append(result, []string{d[j]})
			} else {
				for k := 0; k < len(tmp); k++ {
					result = append(result, append([]string{d[j]}, tmp[k]...))
				}
			}
			bb = true
		}
	}

	if bb {
		mm[i] = 1
		m[i] = result
	} else {
		mm[i] = -1
	}

	return result, bb
}
