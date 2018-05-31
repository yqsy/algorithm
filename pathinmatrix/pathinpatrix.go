package pathinmatrix

import (
	"errors"
)

func recursion(matrix [][]byte, pos, previousPos, xLen, yLen int, tmp *[]byte, dst string) bool {
	row, col := pos/xLen, pos%xLen
	*tmp = append(*tmp, matrix[row][col])

	if len(*tmp) > len(dst) {
		*tmp = (*tmp)[:len(*tmp)-1]
		return false
	}

	if string(*tmp) == dst {
		return true
	}

	if pos%xLen > 0 {
		leftPos := pos - 1
		if leftPos != previousPos && (leftPos > 0 && leftPos < xLen*yLen) {
			if recursion(matrix, leftPos, pos, xLen, yLen, tmp, dst) {
				return true
			}
		}
	}

	if pos%xLen < xLen-1 {
		rightPos := pos + 1
		if rightPos != previousPos && (rightPos > 0 && rightPos < xLen*yLen) {
			if recursion(matrix, rightPos, pos, xLen, yLen, tmp, dst) {
				return true
			}
		}
	}

	if pos >= xLen {
		upPos := pos - xLen
		if upPos != previousPos && (upPos > 0 && upPos < xLen*yLen) {
			if recursion(matrix, upPos, pos, xLen, yLen, tmp, dst) {
				return true
			}
		}
	}

	if pos <= xLen*yLen-1-xLen {
		downPos := pos + xLen
		if downPos != previousPos && (downPos > 0 && downPos < xLen*yLen) {
			if recursion(matrix, downPos, pos, xLen, yLen, tmp, dst) {
				return true
			}
		}
	}

	*tmp = (*tmp)[:len(*tmp)-1]
	return false
}

func hasPath(matrix [][]byte, dst string) (bool, error) {
	// 纵轴长度
	yLen := len(matrix)
	if yLen < 1 {
		return false, errors.New("y err")
	}

	// 横轴长度
	xLen := len(matrix[0])
	if xLen < 1 {
		return false, errors.New("x err")
	}

	// 检查每一row是否长度一致
	for i := 0; i < yLen; i++ {
		if len(matrix[i]) != xLen {
			return false, errors.New("matrix err")
		}
	}

	tmp := make([]byte, 0)
	for i := 0; i < xLen*yLen; i++ {
		if recursion(matrix, i, 0, xLen, yLen, &tmp, dst) {
			return true, nil
		}
	}
	return false, nil
}
