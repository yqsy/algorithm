package avl

import (
	"testing"
	"fmt"
)

func TestAvlSimplePrint(t *testing.T) {
	avl := &AVL{}

	//for i := 0; i < 10; i ++ {
	//	a := strconv.Itoa(i)
	//	avl.Put(a, a)
	//}

	avl.Put("0", "0")
	avl.Put("1", "1")
	avl.Put("2", "2")

	fmt.Println(avl.Prettify())
}
