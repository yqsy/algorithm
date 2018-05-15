package sudoku1

import (
	"github.com/yqsy/algorithm/sudoku/common"
)

func Solve(table *common.Table, pos int) bool {
	// pos 0 ~ 80
	if pos > 80 {
		return true
	}

	row, col := pos/9, pos%9
	if (*table)[row][col] == 0 {
		// 本格子待猜测
		for i := 1; i <= 9; i++ {
			(*table)[row][col] = byte(i)
			if table.IsValid(row, col) {
				if Solve(table, pos+1) {
					return true
				}
			}
			(*table)[row][col] = 0
		}
		return false
	} else {
		// 本格子已完成
		return Solve(table, pos+1)
	}
}
