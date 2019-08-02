package leetcode_go

var result []byte

var row, col, r, c, dr, dc int

func alphabetBoardPath(target string) string {
	result = []byte{}

	row, col = 0, 0

	for i := 0; i < len(target); i++ {
		index := int(target[i] - 'a')

		r = index / 5
		c = index % 5

		dr = r - row
		dc = c - col

		if r == 5 {
			adjustCol()
			adjustRow()
			result = append(result, '!')
			row = r
			col = c
		} else {
			adjustRow()
			adjustCol()
			result = append(result, '!')
			row = r
			col = c
		}
	}

	return string(result)
}

func adjustRow() {
	if dr > 0 {
		for dr != 0 {
			result = append(result, 'D')
			dr--
		}
	}
	if dr < 0 {
		for dr != 0 {
			result = append(result, 'U')
			dr++
		}
	}
}

func adjustCol() {
	if dc > 0 {
		for dc != 0 {
			result = append(result, 'R')
			dc--
		}
	}
	if dc < 0 {
		for dc != 0 {
			result = append(result, 'L')
			dc++
		}
	}
}
