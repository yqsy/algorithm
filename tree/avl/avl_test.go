package avl

import (
	"testing"
	"strconv"
	"fmt"
)

func TestAvlSimplePrint(t *testing.T) {
	avl := &AVL{}

	for i := 0; i < 10; i ++ {
		a := strconv.Itoa(i)
		avl.Put(a, a)
	}

	fmt.Println(avl.Prettify())
}
