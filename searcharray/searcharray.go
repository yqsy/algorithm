package searcharray

import "errors"

func Find(table [][]int, findValue int) (bool, error) {
	if len(table) < 1 {
		return false, errors.New("table is empty")
	}

	if len(table[0]) < 1 {
		return false, errors.New("cols is empty")
	}

	cols := len(table[0])

	// check is matrix
	for i := 0; i < len(table); i++ {
		if len(table[i]) != cols {
			return false, errors.New("table is wrongful")
		}
	}

	// check all rows
	for row := 0; row < len(table); row++ {
		for col := 0; col < cols-1; col++ {
			if table[row][col] >= table[row][col+1] {
				return false, errors.New("table row wrongful")
			}
		}
	}

	// check all cols
	for col := 0; col < cols; col++ {
		for row := 0; row < len(table)-1; row++ {
			if table[row][col] >= table[row+1][col] {
				return false, errors.New("table col wrongful")
			}
		}
	}

	for col := cols - 1; col >= 0; col-- {
		for row := 0; row < len(table); row++ {
			if findValue == table[row][col] {
				return true, nil
			} else if findValue < table[row][col] {
				break
			}
		}
	}
	return false, nil
}
