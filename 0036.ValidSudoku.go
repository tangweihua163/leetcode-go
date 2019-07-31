package leetcode_go

func isValidSudoku(board [][]byte) bool {

	var row = make([]map[byte]int, len(board))
	var col = make([]map[byte]int, len(board))
	var square = make([]map[byte]int, len(board))

	for i := 0; i < len(board); i++ {
		row[i] = make(map[byte]int)
		col[i] = make(map[byte]int)
		square[i] = make(map[byte]int)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {

			if board[i][j] == '.' {
				continue
			}

			if row[i][board[i][j]] == 1 {
				return false
			} else {
				row[i][board[i][j]] = 1
			}

			if col[j][board[i][j]] == 1 {
				return false
			} else {
				col[j][board[i][j]] = 1
			}

			if square[i/3*3+j/3][board[i][j]] == 1 {
				return false
			} else {
				square[i/3*3+j/3][board[i][j]] = 1
			}

		}
	}

	return true
}
