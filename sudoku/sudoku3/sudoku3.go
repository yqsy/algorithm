package sudoku3

import (
	"github.com/yqsy/algorithm/sudoku/common"
)

func Solve(table *common.Table, candidate *common.Candidate, pos int) bool {
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
				candidate.ClearCandidate(k, row, col)
				if Solve(table, candidate, pos+1) {
					return true
				} else {
					candidate.RestoreCandidate(k, row, col)
				}
			}
			(*table)[row][col] = 0
		}

		return false
	} else {
		// 本格子已完成
		return Solve(table, candidate, pos+1)
	}
}
