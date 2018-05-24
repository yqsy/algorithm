package binarytree

import (
	"testing"
	"fmt"
)

func TestCreateTree(t *testing.T) {
	binaryTree := CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})

	fmt.Println(binaryTree.GetAllLayer())

	binaryTree = CreateTree([]int{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(binaryTree.GetAllLayer())

	binaryTree = CreateTree([]int{1, 2, 3, 4, 5, 6})

	fmt.Println(binaryTree.GetAllLayer())

}
