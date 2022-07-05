package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var gs *GameState = NewGameState() // 游戏引擎
var myApp fyne.App                 // 主程序
var myWindow fyne.Window           // 主窗口

type myButton struct {
	bt   *widget.Button
	r, c int
}

// 添加网格按钮
func genButtonGrid(r, c int) [][]*myButton {
	ans := make([][]*myButton, r)
	for i := 0; i < r; i++ {
		ans[i] = make([]*myButton, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			x, y := i, j
			ans[i][j] = &myButton{widget.NewButton("-", nil), x, y}
			ans[i][j].bt.OnTapped = func() {
				println(x, y)
				if !gs.isEmpty(x, y) {
					return
				}
				gs.put(x, y)
				role := gs.board[x][y]
				gs.showBoard()
				ans[x][y].bt.SetText(string(role))
				var boardcp [3][3]byte
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						boardcp[i][j] = gs.board[i][j]
					}
				}
				oMove := bestMove(boardcp)
				gs.put(oMove.row, oMove.col)
				role = gs.board[oMove.row][oMove.col]
				ans[oMove.row][oMove.col].bt.SetText(string(role))
				if gs.isEnd() {
					println("END!")
					myWindow.Close()
				}
			}
		}
	}
	return ans
}

func main() {
	myApp = app.New()
	myWindow = myApp.NewWindow("TIC-TAC-TOE")
	buttonGrid := genButtonGrid(3, 3)
	myWindow.SetContent(container.NewGridWithColumns(1,
		container.NewGridWithColumns(3, buttonGrid[0][0].bt, buttonGrid[0][1].bt, buttonGrid[0][2].bt),
		container.NewGridWithColumns(3, buttonGrid[1][0].bt, buttonGrid[1][1].bt, buttonGrid[1][2].bt),
		container.NewGridWithColumns(3, buttonGrid[2][0].bt, buttonGrid[2][1].bt, buttonGrid[2][2].bt),
	))
	myWindow.Resize(fyne.NewSize(240, 240))
	myWindow.ShowAndRun()
}
