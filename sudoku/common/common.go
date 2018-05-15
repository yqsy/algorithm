package common

import (
	"errors"
	"strconv"
)

type Table [][]byte

type Candidate [][]map[byte]struct{}

// 合法性校验
func (table *Table) IsValid(row, col int) bool {

	// 横向一列 check
	for i := 0; i < 9; i++ {
		if col != i && (*table)[row][col] == (*table)[row][i] {
			return false
		}
	}

	// 竖向一列 check
	for i := 0; i < 9; i++ {
		if row != i && (*table)[row][col] == (*table)[i][col] {
			return false
		}
	}

	// 小九宫格 check
	cellRowBegin := row / 3 * 3
	cellColBegin := col / 3 * 3
	for i := cellRowBegin; i < cellRowBegin+3; i++ {
		for j := cellColBegin; j < cellColBegin+3; j++ {
			if row == i && col == j {
				continue
			}

			if (*table)[row][col] == (*table)[i][j] {
				return false
			}
		}
	}

	return true
}

// 获取候选数表格
func (table *Table) GetCandidate() *Candidate {
	candidate := &Candidate{}
	*candidate = make([][]map[byte]struct{}, 9)
	for row := 0; row < 9; row++ {
		(*candidate)[row] = make([]map[byte]struct{}, 9)
		for col := 0; col < 9; col++ {
			(*candidate)[row][col] = make(map[byte]struct{})

			// 安插候选数
			for i := 1; i <= 9; i++ {
				(*candidate)[row][col][byte(i)] = struct{}{}
			}
		}
	}

	// 首次候选数排除
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			exist := (*table)[row][col]
			if exist != 0 {
				// 横向一列 清除
				for i := 0; i < 9; i ++ {
					delete((*candidate)[row][i], exist)
				}

				// 竖向一列 清除
				for i := 0; i < 9; i++ {
					delete((*candidate)[i][col], exist)
				}

				// 小九宫格 清除
				cellRowBegin := row / 3 * 3
				cellColBegin := col / 3 * 3
				for i := cellRowBegin; i < cellRowBegin+3; i++ {
					for j := cellColBegin; j < cellColBegin+3; j++ {
						delete((*candidate)[i][j], exist)
					}
				}
			}
		}
	}

	return candidate
}

// 动态清除候选数
func (candidate *Candidate) ClearCandidate(exist byte, row, col int) {
	// 横向一列 清除
	for i := 0; i < 9; i ++ {
		if col != i {
			delete((*candidate)[row][i], exist)
		}
	}

	// 竖向一列 清除
	for i := 0; i < 9; i++ {
		if row != i {
			delete((*candidate)[i][col], exist)
		}
	}

	// 小九宫格 清除
	cellRowBegin := row / 3 * 3
	cellColBegin := col / 3 * 3
	for i := cellRowBegin; i < cellRowBegin+3; i++ {
		for j := cellColBegin; j < cellColBegin+3; j++ {

			if row != i && col != j {
				delete((*candidate)[i][j], exist)
			}
		}
	}
}

// 恢复动态清除的候选数
func (candidate *Candidate) RestoreCandidate(exist byte, row, col int) {
	// 横向一列 恢复
	for i := 0; i < 9; i ++ {
		if col != i {
			(*candidate)[row][i][exist] = struct{}{}
		}
	}

	// 竖向一列 恢复
	for i := 0; i < 9; i++ {
		if row != i {
			(*candidate)[i][col][exist] = struct{}{}
		}
	}

	// 小九宫格 恢复
	cellRowBegin := row / 3 * 3
	cellColBegin := col / 3 * 3
	for i := cellRowBegin; i < cellRowBegin+3; i++ {
		for j := cellColBegin; j < cellColBegin+3; j++ {

			if row != i && col != j {
				(*candidate)[i][j][exist] = struct{}{}
			}
		}
	}
}

// only support 9 x 9
// like "904002100032100907100790000800070000765040813000030002000018006208007390007900501\r\n"
// remove \r\n and convert to table
func ConvertLineToTable(line string) (*Table, error) {
	if len(line) > 2 && line[len(line)-2:] == "\r\n" {
		line = line[:len(line)-2]
	}

	if len(line) != 81 {
		return nil, errors.New("not 9 x 9 soduku problem")
	}

	table := &Table{}
	*table = make([][]byte, 9)
	for i := 0; i < 9; i++ {
		(*table)[i] = make([]byte, 9)
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			n, err := strconv.Atoi(string(line[9*row+col]))
			if err != nil {
				return nil, errors.New("line convert error")
			}

			(*table)[row][col] = byte(n)
		}
	}

	return table, nil
}
