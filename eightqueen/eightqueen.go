package main

import "fmt"

type Board [][]bool

type Context struct {
	board Board

	solvedBoards []Board

	// 竖向候选
	cols []bool

	// 正斜着候选(/)  左上角idx=0
	skewsForward []bool

	// 反斜着候选(\)  左下角idx=0
	skewsBack []bool
}

func (ctx *Context) copy(n int) Board {
	copyBoard := make([][]bool, n)
	for i := 0; i < n; i++ {
		copyBoard[i] = make([]bool, n)
	}
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			copyBoard[row][col] = ctx.board[row][col]
		}
	}
	return copyBoard
}

func (ctx *Context) setCell(row, col, n int, b bool) {
	ctx.board[row][col] = b
	ctx.cols[col] = b
	ctx.skewsForward[row+col] = b
	ctx.skewsBack[(n-1-row)+col] = b
}

func (ctx *Context) check(row, col, n int) bool {
	if !ctx.cols[col] &&
		!ctx.skewsForward[row+col] &&
		!ctx.skewsBack[(n-1-row)+col] {
		return true
	} else {
		return false
	}
}

func solve(row, n int, ctx *Context) {
	if row == n {
		ctx.solvedBoards = append(ctx.solvedBoards, ctx.copy(n))
		return
	}

	// 横向每一格进行尝试
	for i := 0; i < n; i++ {
		if !ctx.check(row, i, n) {
			continue
		}
		ctx.setCell(row, i, n, true)
		solve(row+1, n, ctx)
		ctx.setCell(row, i, n, false)
	}
}

func solveNQueens(n int) [][]string {
	ctx := Context{}
	ctx.board = make([][]bool, n)
	for i := 0; i < n; i++ {
		ctx.board[i] = make([]bool, n)
	}
	ctx.cols = make([]bool, n)
	ctx.skewsForward = make([]bool, 2*n-1)
	ctx.skewsBack = make([]bool, 2*n-1)

	solve(0, n, &ctx)

	result := make([][]string, len(ctx.solvedBoards))
	for i := 0; i < len(ctx.solvedBoards); i++ {
		var oneBoard []string
		for row := 0; row < n; row++ {
			var line string
			for col := 0; col < n; col++ {
				if ctx.solvedBoards[i][row][col] {
					line = line + "Q"
				} else {
					line = line + "."
				}
			}
			oneBoard = append(oneBoard, line)
		}
		result[i] = oneBoard
	}

	return result
}

func main() {
	result := solveNQueens(8)
	fmt.Print(result)
}
