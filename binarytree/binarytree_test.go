package binarytree

import (
	"testing"
	"fmt"
)

func TestCreateTree(t *testing.T) {
	binaryTree := Deserialize([]int{1, 2, 3, 4, 5, 6, 7, 8})

	fmt.Println(binaryTree.AllLayer())

	binaryTree = Deserialize([]int{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(binaryTree.AllLayer())

	binaryTree = Deserialize([]int{1, 2, 3, 4, 5, 6})

	fmt.Println(binaryTree.AllLayer())

	serialized := binaryTree.Serialize()
	fmt.Println(serialized)
}

func TestIterTree(t *testing.T) {
	binaryTree := Deserialize([]int{1, 2, 3, 4, 5, 6, 7, 8})

	out := binaryTree.PreOrder()
	fmt.Println(out)

	out = binaryTree.InfixOrder()
	fmt.Println(out)

	out = binaryTree.PostOrder()
	fmt.Println(out)
}

func TestIterTreeStack(t *testing.T) {
	binaryTree := Deserialize([]int{1, 2, 3, 4, 5, 6, 7, 8})

	out := binaryTree.PreOrderStack()
	fmt.Println(out)

	out = binaryTree.InfixOrderStack()
	fmt.Println(out)

	out = binaryTree.PostOrderStack()
	fmt.Println(out)
}
