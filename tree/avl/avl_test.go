package avl

import (
	"testing"
	"fmt"
	"strconv"
)

func TestAvlSimplePrint(t *testing.T) {
	avl := &AVL{}

	for i := 0; i < 11; i ++ {
		avl.Put(i, strconv.Itoa(i))
	}

	fmt.Println(avl.Prettify())
}

// https://stackoverflow.com/questions/3955680/how-to-check-if-my-avl-tree-implementation-is-correct

func Test1a(t *testing.T) {
	avl := &AVL{}

	avl.Put(20, "20")
	avl.Put(4, "4")

	head := avl.head
	if head.left == nil || head.left.key != 4 {
		t.Fatal("err")
	}

	avl.Put(15, "15")

	head = avl.head
	if head.key != 15 {
		t.Fatal("err")
	}

	if head.left == nil || head.left.key != 4 {
		t.Fatal("err")
	}

	if head.right == nil || head.right.key != 20 {
		t.Fatal("err")
	}
}


