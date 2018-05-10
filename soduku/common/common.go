package common

import (
	"errors"
	"strconv"
	"github.com/yqsy/algorithm/repeat_number/repeat_hash"
)

type Row [9]int

type Table [9]Row

type Cols [9]int

// 0 ~ 80
type Pos int

func (pos *Pos) GetRowCol() (int, int) {
	row := *pos / 9
	col := *pos % 9
	return int(row), int(col)
}

func (pos *Pos) NextPos() Pos {
	nextPos := *pos + 1
	return nextPos
}

// 刚填完一个cell对cell的row和col做合法校验
func (table *Table) IsValid(pos Pos) bool {
	row, col := pos.GetRowCol()

	// row check
	checkRow := table[row]
	if repeat_hash.IsArrayRepeatIgnoreZero(checkRow[:]) {
		return false
	}

	// col check
	checkCol := table.GetCol(col)
	if repeat_hash.IsArrayRepeatIgnoreZero(checkCol[:]) {
		return false
	}

	return true
}

// 给定列号,获取该列所有的元素
func (table *Table) GetCol(col int) *Cols {
	cols := &Cols{}

	for row := 0; row < 9; row++ {
		cols[row] = table[row][col]
	}

	return cols
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
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			n, err := strconv.Atoi(string(line[9*row+col]))
			if err != nil {
				return nil, errors.New("line convert error")
			}

			table[row][col] = n
		}
	}

	return table, nil
}
