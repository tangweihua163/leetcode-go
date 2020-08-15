package leetcode_go

import "strings"

func reformatDate(date string) string {
	s := strings.Split(date, " ")

	year := s[2]
	month := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}[s[1]]

	var day string
	if len(s[0]) == 3 {
		day = "0" + s[0]
		day = day[:2]
	} else {
		day = s[0][:2]
	}

	return year + "-" + month + "-" + day

}
