package main

import "fmt"

// 游戏状态结构体
type GameState struct {
	board [3][3]byte // 棋盘
	isX   bool       // 回合
}

// 初始化函数
func NewGameState() *GameState {
	var gs GameState
	gs.board = [3][3]byte{
		{'-', '-', '-'},
		{'-', '-', '-'},
		{'-', '-', '-'},
	}
	gs.isX = true
	return &gs
}

// 下棋
func (this *GameState) put(i, j int) {
	var role byte
	if this.isX {
		role = 'X'
	} else {
		role = 'O'
	}
	this.board[i][j] = role
	this.isX = !this.isX
}

func (this *GameState) back(i, j int) {
	this.board[i][j] = '-'
	this.isX = !this.isX
}

// 判断是否为空棋格
func (this GameState) isEmpty(row, col int) bool {
	if this.board[row][col] == '-' {
		return true
	}
	return false
}

// 判断是否获胜
func (this GameState) isWin() bool {
	// Rows
	for row := 0; row < 3; row++ {
		if this.board[row][0] == this.board[row][1] && this.board[row][1] == this.board[row][2] && this.board[row][0] != '-' {
			return true
		}
	}
	// Colums
	for col := 0; col < 3; col++ {
		if this.board[0][col] == this.board[1][col] && this.board[1][col] == this.board[2][col] && this.board[0][col] != '-' {
			return true
		}
	}
	// Diagonals
	if this.board[0][0] == this.board[1][1] && this.board[1][1] == this.board[2][2] && this.board[0][0] != '-' {
		return true
	}
	if this.board[0][2] == this.board[1][1] && this.board[1][1] == this.board[2][0] && this.board[0][2] != '-' {
		return true
	}
	return false
}

// 判断是否游戏结束
func (this GameState) isEnd() bool {
	if this.isWin() {
		return true
	}
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if this.isEmpty(row, col) {
				return false
			}
		}
	}
	return true
}

// 显示棋盘状态
func (this GameState) showBoard() {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			fmt.Printf("%c ", this.board[row][col])
		}
		fmt.Println()
	}
}

// func main() {
// 	var gs GameState
// 	gs.board = [3][3]byte{
// 		{'-', '-', '-'},
// 		{'-', '-', '-'},
// 		{'-', '-', '-'},
// 	}
// 	showBoard(gs.board)
// 	var row, col int
// 	for !isEnd(gs.board) {
// 		fmt.Scanf("%d %d", &row, &col)
// 		if isEmpty(gs.board, row, col) {
// 			gs.board[row][col] = 'x'
// 			showBoard(gs.board)
// 			fmt.Println()
// 			if isEnd(gs.board) {
// 				fmt.Println("Game Over")
// 				break
// 			}
// 			oMove := bestMove(gs.board)
// 			gs.board[oMove.row][oMove.col] = 'o'
// 			showBoard(gs.board)
// 			fmt.Println("------------------------------------------------")
// 		}
// 	}
// }
