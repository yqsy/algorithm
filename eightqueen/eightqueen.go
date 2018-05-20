package main

import "fmt"

type Context struct {
	board [][]byte

	solvedBoards [][]string

	// 竖向候选
	cols []bool

	// 正斜着候选(/)  左上角idx=0
	skewsForward []bool

	// 反斜着候选(\)  左下角idx=0
	skewsBack []bool
}

func (ctx *Context) setCell(row, col, n int, b byte, flag bool) {
	ctx.board[row][col] = b
	ctx.cols[col] = flag
	ctx.skewsForward[row+col] = flag
	ctx.skewsBack[(n-1-row)+col] = flag
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
		table := make([]string, n)
		for i, chessRow := range ctx.board {
			table[i] = string(chessRow[:])
		}
		ctx.solvedBoards = append(ctx.solvedBoards, table)
		return
	}

	// 横向每一格进行尝试
	for i := 0; i < n; i++ {
		if !ctx.check(row, i, n) {
			continue
		}
		ctx.setCell(row, i, n, 'Q', true)
		solve(row+1, n, ctx)
		ctx.setCell(row, i, n, '.', false)
	}
}

func solveNQueens(n int) [][]string {
	ctx := Context{}
	ctx.board = make([][]byte, n)
	for i := 0; i < n; i++ {
		ctx.board[i] = make([]byte, n)

		for j := 0; j < n; j++ {
			ctx.board[i][j] = '.'
		}
	}
	ctx.cols = make([]bool, n)
	ctx.skewsForward = make([]bool, 2*n-1)
	ctx.skewsBack = make([]bool, 2*n-1)

	solve(0, n, &ctx)

	return ctx.solvedBoards
}

func main() {
	result := solveNQueens(8)
	fmt.Print(result)
}
