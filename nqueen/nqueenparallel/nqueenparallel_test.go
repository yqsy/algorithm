package nqueenparallel

import (
	"testing"
	"time"
)

func TestNqueen(t *testing.T) {
	cm := NewComputingModule()
	cm.Prepare(11)

	t1 := time.Now()
	allSolvedBoards := cm.Solve(11)
	_ = allSolvedBoards
	escaped := float64(time.Since(t1).Nanoseconds()) / 1000
	t.Logf("esacped: %.2f us\n", escaped)
}
