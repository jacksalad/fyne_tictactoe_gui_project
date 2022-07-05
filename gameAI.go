package main

// 走法结构体
type Move struct {
	row, col int
}

// 评价函数
func evaluate(board [3][3]byte) int {
	// Rows
	for row := 0; row < 3; row++ {
		if board[row][0] == board[row][1] && board[row][1] == board[row][2] {
			if board[row][0] == 'X' {
				return 1
			}
			if board[row][0] == 'O' {
				return -1
			}
		}
	}
	// Colums
	for col := 0; col < 3; col++ {
		if board[0][col] == board[1][col] && board[1][col] == board[2][col] {
			if board[0][col] == 'X' {
				return 1
			}
			if board[0][col] == 'O' {
				return -1
			}
		}
	}
	// Diagonals
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		if board[0][0] == 'X' {
			return 1
		}
		if board[0][0] == 'O' {
			return -1
		}
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		if board[0][2] == 'X' {
			return 1
		}
		if board[0][2] == 'O' {
			return -1
		}
	}
	return 0
}

// minimax函数
func minimax(board [3][3]byte, isMax bool) int {
	score := evaluate(board)
	if score != 0 {
		return score
	}
	if isFull(board) {
		return 0
	}

	if isMax {
		maxVal := -10
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				if board[row][col] == '-' {
					board[row][col] = 'X'
					maxVal = max(maxVal, minimax(board, !isMax))
					board[row][col] = '-'
				}
			}
		}
		return maxVal
	} else {
		minVal := 10
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				if board[row][col] == '-' {
					board[row][col] = 'O'
					minVal = min(minVal, minimax(board, !isMax))
					board[row][col] = '-'
				}
			}
		}
		return minVal
	}
}

// 最佳走法函数
func bestMove(board [3][3]byte) Move {
	var move Move
	move.row, move.col = -1, -1
	minVal := 10
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row][col] == '-' {
				board[row][col] = 'O'
				moveVal := minimax(board, true)
				board[row][col] = '-'
				if moveVal < minVal {
					move.row = row
					move.col = col
					minVal = moveVal
				}
			}
		}
	}
	return move
}

// 判断是否棋盘已满
func isFull(board [3][3]byte) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row][col] == '-' {
				return false
			}
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
