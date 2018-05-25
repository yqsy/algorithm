package binarytree

import (
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

func (node *Node) PreOrder(out *[]int) {
	*out = append(*out, node.data)
	if node.left != nil {
		node.left.PreOrder(out)
	}

	if node.right != nil {
		node.right.PreOrder(out)
	}
}

func (node *Node) InfixOrder(out *[]int) {
	if node.left != nil {
		node.left.InfixOrder(out)
	}

	*out = append(*out, node.data)

	if node.right != nil {
		node.right.InfixOrder(out)
	}
}

func (node *Node) PostOrder(out *[]int) {
	if node.left != nil {
		node.left.PostOrder(out)
	}

	if node.right != nil {
		node.right.PostOrder(out)
	}

	*out = append(*out, node.data)
}

type BinaryTree struct {
	head *Node
}

// 前序(栈)
func (binaryTree *BinaryTree) PreOrderStack() []int {
	if binaryTree.head == nil {
		return []int{}
	}

	var out []int
	s := stack.New()
	s.Push(binaryTree.head)

	for {
		if s.Len() < 1 {
			break
		}

		node := s.Pop().(*Node)
		out = append(out, node.data)

		if node.right != nil {
			s.Push(node.right)
		}

		if node.left != nil {
			s.Push(node.left)
		}
	}

	return out
}

// 中序(栈)
func (binaryTree *BinaryTree) InfixOrderStack() []int {
	if binaryTree.head == nil {
		return []int{}
	}

	var out []int
	s := stack.New()

	node := binaryTree.head
	for {
		for {
			if node == nil {
				break
			}

			s.Push(node)
			node = node.left
		}

		if s.Len() < 1 {
			break
		}

		popNode := s.Pop().(*Node)
		out = append(out, popNode.data)
		if popNode.right != nil {
			node = popNode.right
		}
	}

	return out
}

// 后序(栈)
func (binaryTree *BinaryTree) PostOrderStack() []int {
	if binaryTree.head == nil {
		return []int{}
	}

	var out []int
	s := stack.New()
	s.Push(binaryTree.head)

	for {
		if s.Len() < 1 {
			break
		}

		node := s.Pop().(*Node)
		out = append(out, node.data)

		if node.left != nil {
			s.Push(node.left)
		}

		if node.right != nil {
			s.Push(node.right)
		}
	}

	// reverse
	half := len(out) / 2

	for i := 0; i < half; i++ {
		out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
	}

	return out
}

// 前序
func (binaryTree *BinaryTree) PreOrder() []int {
	if binaryTree.head == nil {
		return []int{}
	}

	var out []int
	binaryTree.head.PreOrder(&out)
	return out
}

// 中序
func (binaryTree *BinaryTree) InfixOrder() []int {
	if binaryTree.head == nil {
		return []int{}
	}
	var out []int
	binaryTree.head.InfixOrder(&out)
	return out
}

// 后序
func (binaryTree *BinaryTree) PostOrder() []int {
	if binaryTree.head == nil {
		return []int{}
	}
	var out []int
	binaryTree.head.PostOrder(&out)
	return out
}

// 层次遍历
func (binaryTree *BinaryTree) AllLayer() [][]int {
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

// 序列化二叉树
func (binaryTree *BinaryTree) Serialize() []int {
	if binaryTree.head == nil {
		return []int{}
	}

	var serialized []int

	s := queue.New()
	s.Enqueue(binaryTree.head)

	for {
		currentRowLen := s.Len()

		if currentRowLen < 1 {
			break
		}

		for i := 0; i < currentRowLen; i++ {
			node := s.Dequeue().(*Node)
			serialized = append(serialized, node.data)
			l, r := node.left, node.right
			if l != nil {
				s.Enqueue(l)
			}

			if r != nil {
				s.Enqueue(r)
			}
		}

	}
	return serialized
}

// 反序列化树
// 1234567
// 生成
//    1
//  2  3
// 4 5 6 7
func Deserialize(nodes []int) *BinaryTree {
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
