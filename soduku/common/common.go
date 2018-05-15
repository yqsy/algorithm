package common

import (
	"errors"
	"strconv"
)

type Table [][]byte

// 合法性校验
func (table *Table) IsValid(row int, col int) bool {

	// 横向一列 chec
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
