package nqueen

import (
	"testing"
	"time"
)

func TestNqueen(t *testing.T) {
	t1 := time.Now()
	result := solveNQueens(11)
	_ = result
	escaped := float64(time.Since(t1).Nanoseconds()) / 1000
	t.Logf("esacped: %.2f us\n", escaped)
}
