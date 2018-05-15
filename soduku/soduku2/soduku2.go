package main

import (
	"github.com/yqsy/algorithm/soduku/common"
	"fmt"
)

func solve(table *common.Table, candidate *common.Candidate, pos int) bool {
	// pos 0 ~ 80
	if pos > 80 {
		return true
	}

	row, col := pos/9, pos%9
	if (*table)[row][col] == 0 {
		// 本格子待猜测
		for k, _ := range (*candidate)[row][col] {
			(*table)[row][col] = k
			if table.IsValid(row, col) {
				if solve(table, candidate, pos+1) {
					return true
				}
			}
			(*table)[row][col] = 0
		}

		return false
	} else {
		// 本格子已完成
		return solve(table, candidate, pos+1)
	}
}

func main() {
	problem := "000000010400000000020000000000050407008000300001090000300400200050100000000806000\r\n"

	//problem := "080001030500804706000270000920400003103958402400002089000029000305106008040300010"

	table, err := common.ConvertLineToTable(problem)
	if err != nil {
		panic(err)
	}

	fmt.Println(table)

	candidate := table.GetCandidate()

	var pos int
	if solve(table, candidate, pos) {
		fmt.Println("ok")
		fmt.Println(table)
	} else {
		fmt.Println("error")
		fmt.Println(table)
	}
}
