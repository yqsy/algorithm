package main

import (
	"github.com/yqsy/algorithm/soduku/common"
	"fmt"
)

type Contex struct {
	// 已经求解的
	solved *int

	// 数独表
	table *common.Table
}

func (ctx *Contex) isSolved() bool {
	if *ctx.solved == 9*9 {
		return true
	} else {
		return false
	}
}

func solve(ctx *Contex, pos common.Pos) bool {
	if ctx.isSolved() {
		return true
	}

	row, col := pos.GetRowCol()
	if ctx.table[row][col] != 0 {
		*ctx.solved += 1
		newPos := pos.NextPos()
		if solve(ctx, newPos) {
			return true
		} else {
			*ctx.solved -= 1
			return false
		}
	}

	for i := 1; i <= 9; i ++ {
		ctx.table[row][col] = i
		if ctx.table.IsValid(pos) {
			*ctx.solved += 1
			newPos := pos.NextPos()
			if solve(ctx, newPos) {
				return true
			} else {
				if i == 9 {
					ctx.table[row][col] = 0
					*ctx.solved -= 1
				}
			}
		}
	}

	return false
}

func main() {
	problem := "904002100032100907100790000800070000765040813000030002000018006208007390007900501\r\n"

	table, err := common.ConvertLineToTable(problem)
	if err != nil {
		panic(err)
	}

	fmt.Println(table)

	ctx := &Contex{solved: new(int), table: table}
	var pos common.Pos
	if solve(ctx, pos) {
		fmt.Println("ok")
		fmt.Println(table)
	} else {
		fmt.Println("error")
		fmt.Println(table)
	}
}
