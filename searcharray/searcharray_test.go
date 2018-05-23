package searcharray

import "testing"

func TestOk(t *testing.T) {

	t1 := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	ok, err := Find(t1, 15)

	if err != nil || !ok {
		t.Fatal("can't find")
	}

	ok, err = Find(t1, 1)

	if err != nil || !ok {
		t.Fatal("can't find")
	}

	ok, err = Find(t1, 7)

	if err != nil || !ok {
		t.Fatal("can't find")
	}
}

func TestNotOk(t *testing.T) {
	t1 := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	ok, err := Find(t1, 0)

	if err != nil || ok {
		t.Fatal("find")
	}

	ok, err = Find(t1, 16)

	if err != nil || ok {
		t.Fatal("find")
	}

	ok, err = Find(t1, 14)

	if err != nil || ok {
		t.Fatal("find")
	}
}

func TestErr(t *testing.T) {

	var t1 [][]int

	ok, err := Find(t1, 0)

	if err == nil || ok {
		t.Fatal("err test")
	}

	t1 = [][]int{
		{1, 2, 3},
		{4},
	}

	ok, err = Find(t1, 0)

	if err == nil || ok {
		t.Fatal("err test")
	}

	t1 = [][]int{
		{1, 2, 3, 0},
		{4, 5, 6},
	}

	ok, err = Find(t1, 0)

	if err == nil || ok {
		t.Fatal("err test")
	}
}
