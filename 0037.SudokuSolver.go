package leetcode_go

var row = make([]map[byte]int, 9)
var col = make([]map[byte]int, 9)
var square = make([]map[byte]int, 9)

var candidate []int

var suc = false

func solveSudoku(board [][]byte) {

	suc = false
	candidate = make([]int, 0, 162)

	for i := 0; i < len(board); i++ {
		row[i] = make(map[byte]int)
		col[i] = make(map[byte]int)
		square[i] = make(map[byte]int)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == '.' {
				candidate = append(candidate, i, j)
				row[i][board[i][j]] = 0
				col[j][board[i][j]] = 0
				square[i/3*3+j/3][board[i][j]] = 0
				continue
			}
			row[i][board[i][j]] = 1
			col[j][board[i][j]] = 1
			square[i/3*3+j/3][board[i][j]] = 1
		}
	}

	cur := len(candidate) - 1

	solve(board, cur)

}

func solve(board [][]byte, cur int) {

	if cur < 0 {
		suc = true
		return
	}

	var num byte
	for num = '1'; num <= '9'; num++ {
		var i = candidate[cur-1]
		var j = candidate[cur]
		r := row[i]
		c := col[j]
		s := square[i/3*3+j/3]
		if r[num] == 0 && c[num] == 0 && s[num] == 0 {
			cur -= 2
			r[num] = 1
			c[num] = 1
			s[num] = 1
			board[i][j] = num

			solve(board, cur)

			if suc {
				return
			}

			board[i][j] = '.'
			r[num] = 0
			c[num] = 0
			s[num] = 0
			cur += 2
		}
	}
}
