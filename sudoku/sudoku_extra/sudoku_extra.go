package sudoku_extra

import (
	"errors"
	"strconv"
)

type Table struct {
	Tb [][]byte

	// 横向的9个数字候选区
	rows [][]bool

	// 竖向的9个数字候选区
	cols [][]bool

	// 小9宫格的9个数字候选区
	cells [][]bool
}

func Solve(table *Table, pos int) bool {
	// pos 0 ~ 80
	if pos > 80 {
		return true
	}

	row, col := pos/9, pos%9
	if table.Tb[row][col] == 0 {
		// 本格子待猜测
		for i := 0; i < 9; i++ {
			if table.rows[row][i] || table.cols[col][i] || table.cells[row-row%3+col/3][i] {
				continue
			}

			table.Tb[row][col] = byte(i + 1)
			table.rows[row][i] = true
			table.cols[col][i] = true
			table.cells[row-row%3+col/3][i] = true

			if Solve(table, pos+1) {
				return true
			}

			table.Tb[row][col] = 0
			table.rows[row][i] = false
			table.cols[col][i] = false
			table.cells[row-row%3+col/3][i] = false
		}

		return false
	} else {
		// 本格子已完成
		return Solve(table, pos+1)
	}
}

func ConvertLineToTable(line string) (*Table, error) {
	if len(line) > 2 && line[len(line)-2:] == "\r\n" {
		line = line[:len(line)-2]
	}

	if len(line) != 81 {
		return nil, errors.New("not 9 x 9 soduku problem")
	}

	// 初始表
	table := &Table{}
	table.Tb = make([][]byte, 9)
	for i := 0; i < 9; i++ {
		table.Tb[i] = make([]byte, 9)
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			n, err := strconv.Atoi(string(line[9*row+col]))
			if err != nil {
				return nil, errors.New("line convert error")
			}

			table.Tb[row][col] = byte(n)
		}
	}

	table.rows = make([][]bool, 9)
	table.cols = make([][]bool, 9)
	table.cells = make([][]bool, 9)

	// 初始候选区
	for i := 0; i < 9; i++ {
		table.rows[i] = []bool{false, false, false, false, false, false, false, false, false}
		table.cols[i] = []bool{false, false, false, false, false, false, false, false, false}
		table.cells[i] = []bool{false, false, false, false, false, false, false, false, false}
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if table.Tb[row][col] == 0 {
				continue
			}
			num := table.Tb[row][col] - 1
			table.rows[row][num] = true
			table.cols[col][num] = true
			table.cells[row-row%3+col/3][num] = true
		}
	}

	return table, nil
}
