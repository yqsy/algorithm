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

func TestCreateTreePreInfixPost(t *testing.T) {
	// 1 2 3 4 5 6 7 8 9
	pre := []int{1, 2, 4, 8, 9, 5, 3, 6, 7}
	infix := []int{8, 4, 9, 2, 5, 1, 6, 3, 7}
	post := []int{8, 9, 4, 5, 2, 6, 7, 3, 1}

	tree := buildTreeFromPreInfix(pre, infix)

	fmt.Println(tree.AllLayer())

	tree = buildTreeFromInfixPost(infix, post)

	fmt.Println(tree.AllLayer())
}

func TestNextNode(t *testing.T) {

	binaryTree := Deserialize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	node := binaryTree.FindNode(5)

	fmt.Println(node.data)

	nextNode := binaryTree.NextInfixNode(node)

	if nextNode != nil {
		fmt.Println(nextNode.data)
	}

}
