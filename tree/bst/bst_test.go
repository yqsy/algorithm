package bst

import (
	"testing"
	"fmt"
)

func TestBSTPrettify(t *testing.T) {
	bst := BST{}
	bst.Put("3", "C")
	bst.Put("1", "A")
	bst.Put("2", "B")
	bst.Put("4", "D")
	bst.Put("5", "E")
	bst.Put("6", "F")
	bst.Put("7", "G")
	fmt.Println(bst.Prettify())
}

func TestBST_Simple(t *testing.T) {
	bst := BST{}
	bst.Put("3", "C")
	bst.Put("1", "A")
	bst.Put("2", "B")
	bst.Put("4", "D")
	bst.Put("5", "E")
	bst.Put("6", "F")
	bst.Put("7", "G")

	v, _ := bst.Get("3")
	if v != "C" {
		t.Fatal("err")
	}

	v, _ = bst.Get("5")
	if v != "E" {
		t.Fatal("err")
	}

	v, _ = bst.Get("7")
	if v != "G" {
		t.Fatal("err")
	}

	bst.Put("8", "H")

	v, _ = bst.Get("8")
	if v != "H" {
		t.Fatal("err")
	}

	min, _ := bst.Min()
	if min != "1" {
		t.Fatal("err")
	}

	max, _ := bst.Max()
	if max != "8" {
		t.Fatal("err")
	}

	floor, _ := bst.Floor("9")
	if floor != "8" {
		t.Fatal("err")
	}

	ceiling, _ := bst.Ceiling("4")
	if ceiling != "4" {
		t.Fatal("err")
	}

	sel, _ := bst.Select(3)
	if sel != "4" {
		t.Fatal("err")
	}

	rank, _ := bst.Rank("6")
	if rank != 5 {
		t.Fatal("err")
	}

	bst.Delete("4")
	v, ok := bst.Get("4")
	if ok {
		t.Fatal("err")
	}

	bst.DeleteMin()
	v, ok = bst.Get("1")
	if ok {
		t.Fatal("err")
	}

	bst.DeleteMax()
	v, ok = bst.Get("8")
	if ok {
		t.Fatal("err")
	}
}
