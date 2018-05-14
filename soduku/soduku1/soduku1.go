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
	if ctx.table[row][col] == 0 {
		// 本格子待猜测
		for i := 1; i <= 9; i++ {
			ctx.table[row][col] = byte(i)
			if ctx.table.IsValid(pos) {
				*ctx.solved += 1
				nextPost := pos.NextPos()
				if solve(ctx, nextPost) {
					return true
				} else {
					*ctx.solved -= 1
				}
			}
		}

		// 本格子遍历完所有的数字都不行,重置下
		ctx.table[row][col] = 0
		return false
	} else {
		// 本格子已完成
		*ctx.solved += 1
		nextPost := pos.NextPos()
		if solve(ctx, nextPost) {
			return true
		} else {
			*ctx.solved -= 1
		}
		return false
	}
}

func main() {
	//problem := "000000010400000000020000000000050407008000300001090000300400200050100000000806000\r\n"

	problem := "080001030500804706000270000920400003103958402400002089000029000305106008040300010"

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
