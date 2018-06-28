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



}
