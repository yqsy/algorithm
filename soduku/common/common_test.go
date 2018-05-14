package common

import "testing"

func TestPos(t *testing.T) {

	var pos Pos

	if pos != 0 {
		t.Fatal("error")
	}

	for i := 0; i < 81; i++ {
		pos = pos.NextPos()
	}

	row, col := pos.GetRowCol()

	if row != 9 || col != 0 {
		t.Fatalf("NextPos error row:%v col:%v", row, col)
	}

}
