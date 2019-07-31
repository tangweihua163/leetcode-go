package leetcode_go

func solveSudoku(board [][]byte) {

	var candidate = make([]int, 0, 162)

	var row = make([]byte, 81)
	var col = make([]byte, 81)
	var square = make([]byte, 81)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b := board[i][j]
			if b == '.' {
				candidate = append(candidate, i, j)
				continue
			}
			row[i*9+int(b-'1')] = 1
			col[j*9+int(b-'1')] = 1
			square[(i/3*3+j/3)*9+int(b-'1')] = 1
		}
	}

	cur := len(candidate) - 1

	solve(board, cur, candidate, row, col, square)

}

func solve(board [][]byte, cur int, candidate []int, row, col, square []byte) bool {

	if cur < 0 {
		return true
	}

	var b byte
	for b = '1'; b <= '9'; b++ {
		var i = candidate[cur-1]
		var j = candidate[cur]

		r := row[i*9 : i*9+9]
		c := col[j*9 : j*9+9]
		s := square[(i/3*3+j/3)*9 : (i/3*3+j/3)*9+9]

		if r[b-'1'] == 0 && c[b-'1'] == 0 && s[b-'1'] == 0 {
			cur -= 2
			r[b-'1'] = 1
			c[b-'1'] = 1
			s[b-'1'] = 1
			board[i][j] = b

			if solve(board, cur, candidate, row, col, square) {
				return true
			}

			board[i][j] = '.'
			r[b-'1'] = 0
			c[b-'1'] = 0
			s[b-'1'] = 0
			cur += 2
		}
	}

	return false
}
