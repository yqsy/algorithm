package main

import "fmt"

type Board [][]bool

type Context struct {
	board Board

	solvedBoards []Board

	// 横向候选
	rows []bool

	// 竖向候选
	cols []bool

	// 正斜着候选(/)  左上角idx=0
	skewsForward []bool

	// 反斜着候选(\)  左下角idx=0
	skewsBack []bool
}

func solve(pos, n int, ctx *Context) {
	if pos == n*n {
		var copyBoard Board
		//copy(copyBoard, ctx.board)
		ctx.solvedBoards = append(ctx.solvedBoards, copyBoard)

		var ok bool = true
		for i := 0; i < n; i++ {
			if !ctx.rows[i] || !ctx.cols[i] {
				ok = false
				break
			}
		}

		for i := 0; i < 2*n-1; i++ {
			if !ctx.skewsForward[i] || !ctx.skewsBack[i] {
				ok = false
				break
			}
		}

		if ok {
			fmt.Println(ctx.board)
		}

		return
	}

	row, col := pos/n, pos%n
	if !ctx.rows[row] && !ctx.cols[col] && !ctx.skewsForward[row+col] && !ctx.skewsBack[(n-1-row)+col] {
		ctx.board[row][col] = true
		ctx.rows[row] = true
		ctx.cols[col] = true
		ctx.skewsForward[row+col] = true
		ctx.skewsBack[(n-1-row)+col] = true
		solve(pos+1, n, ctx)
		ctx.board[row][col] = false
		ctx.rows[row] = false
		ctx.cols[col] = false
		ctx.skewsForward[row+col] = false
		ctx.skewsBack[(n-1-row)+col] = false
		solve(pos+1, n, ctx)
	} else {
		solve(pos+1, n, ctx)
	}
}

func solveNQueens(n int) [][]string {
	ctx := Context{}
	ctx.board = make([][]bool, n)
	for i := 0; i < n; i++ {
		ctx.board[i] = make([]bool, n)
	}

	ctx.rows = make([]bool, n)
	ctx.cols = make([]bool, n)
	ctx.skewsForward = make([]bool, 2*n-1)
	ctx.skewsBack = make([]bool, 2*n-1)

	solve(0, n, &ctx)

	fmt.Print(ctx.solvedBoards)

	return [][]string{}
}

func main() {
	solveNQueens(8)
}
