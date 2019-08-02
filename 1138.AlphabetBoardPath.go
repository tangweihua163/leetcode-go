package leetcode_go

func alphabetBoardPath(target string) string {
	var result []byte

	row, col := 0, 0

	for i := 0; i < len(target); i++ {
		index := int(target[i] - 'a')

		r := index / 5
		c := index % 5

		dr := r - row
		dc := c - col

		if row == 5 {

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

			result = append(result, '!')

			row = r
			col = c

			continue

		}

		if r == 5 {
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

			result = append(result, '!')

			row = r
			col = c

			continue

		}

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

		result = append(result, '!')

		row = r
		col = c

		continue

	}

	return string(result)
}
