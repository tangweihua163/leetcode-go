package leetcode_go

const X = 'X'
const O = 'O'
const A = 'A'

func solveSurroundedRegions(board [][]byte) {

	if len(board) < 3 || len(board[0]) < 3 {
		return
	}

	for i := 1; i <= len(board)-2; i++ {
		for j := 1; j <= len(board[0])-2; j++ {
			if board[i][j] == O {
				searchBoard(board, i, j)
			}
		}
	}

	for i := 0; i <= len(board)-1; i++ {
		for j := 0; j <= len(board[0])-1; j++ {
			if board[i][j] == A {
				board[i][j] = O
			}
		}
	}

}

func searchBoard(board [][]byte, i, j int) {
	b := make(map[int]bool)
	s := true

	height := len(board)
	width := len(board[0])

	queue := make([]int, 1)
	queue[0] = i*width + j
	board[i][j] = A
	b[i*width+j] = true

	for len(queue) != 0 {
		i := queue[0] / width
		j := queue[0] % width

		if i == 0 || j == 0 || i == height-1 || j == width-1 {
			s = false
		}

		if i > 0 && board[i-1][j] == O {
			board[i-1][j] = A
			b[(i-1)*width+j] = true
			queue = append(queue, (i-1)*width+j)
		}
		if i < height-1 && board[i+1][j] == O {
			board[i+1][j] = A
			b[(i+1)*width+j] = true
			queue = append(queue, (i+1)*width+j)
		}
		if j > 0 && board[i][j-1] == O {
			board[i][j-1] = A
			b[i*width+j-1] = true
			queue = append(queue, i*width+j-1)
		}
		if j < width-1 && board[i][j+1] == O {
			board[i][j+1] = A
			b[i*width+j+1] = true
			queue = append(queue, i*width+j+1)
		}

		queue = queue[1:]
	}

	if s {
		for k := range b {
			board[k/width][k%width] = X
		}
	}

}
