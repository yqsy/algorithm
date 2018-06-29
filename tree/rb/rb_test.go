package rb

import "testing"

func TestSimpleInsert(t *testing.T) {
	rb := &RB{}
	rb.Put(19, "S")
	rb.Put(5, "E")
	rb.Put(1, "A")
	rb.Put(18, "R")
	rb.Put(3, "C")
	rb.Put(8, "H")
	rb.Put(24, "X")
	rb.Put(13, "M")
	rb.Put(16, "P")
	rb.Put(12, "L")

	// key check
	if rb.head.key != 13 ||
		rb.head.left.key != 5 ||
		rb.head.right.key != 18 ||
		rb.head.left.left.key != 3 ||
		rb.head.left.right.key != 12 ||
		rb.head.right.left.key != 16 ||
		rb.head.right.right.key != 24 ||
		rb.head.left.left.left.key != 1 ||
		rb.head.left.right.left.key != 8 ||
		rb.head.right.right.left.key != 19 {
		t.Fatal("err")
	}

	// color check
	if rb.head.left.left.left.color != RED ||
		rb.head.left.right.left.color != RED ||
		rb.head.right.right.left.color != RED {
		t.Fatal("err")
	}
}

func TestSimpleInsert2(t *testing.T) {
	rb := &RB{}
	rb.Put(1, "A")
	rb.Put(3, "C")
	rb.Put(5, "E")
	rb.Put(8, "H")
	rb.Put(12, "L")
	rb.Put(13, "M")
	rb.Put(16, "P")
	rb.Put(18, "R")
	rb.Put(19, "S")
	rb.Put(24, "X")

	// key check
	if rb.head.key != 8 ||
		rb.head.left.key != 3 ||
		rb.head.right.key != 18 ||
		rb.head.left.left.key != 1 ||
		rb.head.left.right.key != 5 ||
		rb.head.right.left.key != 13 ||
		rb.head.right.right.key != 24 ||
		rb.head.right.left.left.key != 12 ||
		rb.head.right.left.right.key != 16 ||
		rb.head.right.right.left.key != 19 {
		t.Fatal("err")
	}

	// color check
	if rb.head.right.left.color != RED ||
		rb.head.right.right.left.color != RED {
		t.Fatal("err")
	}
}
