package binarytree

import (
	"github.com/golang-collections/collections/queue"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

type BinaryTree struct {
	head *Node
}

// 层次遍历
func (binaryTree *BinaryTree) GetAllLayer() [][]int {
	if binaryTree.head == nil {
		return [][]int{}
	}

	var nodes [][]int

	s := queue.New()
	s.Enqueue(binaryTree.head)

	for {
		currentRowLen := s.Len()

		if currentRowLen < 1 {
			break
		}

		var oneRow []int
		for i := 0; i < currentRowLen; i++ {
			node := s.Dequeue().(*Node)
			oneRow = append(oneRow, node.data)
			l, r := node.left, node.right
			if l != nil {
				s.Enqueue(l)
			}

			if r != nil {
				s.Enqueue(r)
			}
		}
		nodes = append(nodes, oneRow)
	}

	return nodes
}

// 反序列化树
// 1234567
// 生成
//    1
//  2  3
// 4 5 6 7
func CreateTree(nodes []int) *BinaryTree {
	if len(nodes) < 1 {
		return nil
	}

	binaryTree := &BinaryTree{}
	binaryTree.head = NewNode(nodes[0])

	if len(nodes) == 1 {
		return binaryTree
	}

	s := queue.New()
	s.Enqueue(binaryTree.head)

	nodes = nodes[1:]
	for i := 0; i < len(nodes)/2; i++ {
		l, r := nodes[2*i], nodes[2*i+1]
		node := s.Dequeue().(*Node)

		node.left = NewNode(l)
		node.right = NewNode(r)
		s.Enqueue(node.left)
		s.Enqueue(node.right)
	}

	if len(nodes)%2 != 0 {
		node := s.Dequeue().(*Node)
		node.left = NewNode(nodes[len(nodes)-1])
	}

	return binaryTree
}
