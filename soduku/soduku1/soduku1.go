package main

import (
	"github.com/yqsy/algorithm/soduku/common"
	"fmt"
)

type Contex struct {
	// 已经求解的
	solved int

	// 数独表
	table common.Table
}

func solve(ctx *Contex) {

}

func main() {
	problem := "904002100032100907100790000800070000765040813000030002000018006208007390007900501\r\n"

	table, err := common.ConvertLineToTable(problem)
	if err != nil {
		panic(err)
	}

	fmt.Println(table)

	cols := table.GetCol(4)
	fmt.Println(cols)
}
