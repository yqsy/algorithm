package avl

import (
	"math"
	"github.com/yqsy/algorithm/tree/common"
)

type Node struct {
	key   string
	value string

	left, right *Node

	// 该结点到叶节点的最远距离 (低 -> 高 开始点起)
	height int
}

func NewNode(key string, value string, height int) *Node {
	return &Node{key: key, value: value, height: height}
}

func isAllNodesNil(nodes []*Node) bool {
	for _, val := range nodes {
		if val != nil {
			return false
		}
	}
	return true
}

type AVL struct {
	head *Node
}

func (avl *AVL) Prettify() string {
	if avl.head == nil {
		return ""
	}

	nodes := []*Node{avl.head}
	height := avl.head.height
	s := ""
	avl.prettifyNode(nodes, 1, height, &s)
	return s
}

func (avl *AVL) Get(key string) (string, bool) {
	node := avl.getNode(avl.head, key)
	if node != nil {
		return node.value, true
	} else {
		return "", false
	}
}

func (avl *AVL) Put(key, value string) {
	avl.head = avl.putNode(avl.head, key, value)
}

func (avl *AVL) Delete(key string) {
	avl.head = avl.deleteNode(avl.head, key)
}

func (avl *AVL) minNode(node *Node) (*Node) {
	if node == nil {
		return nil
	}

	if node.left == nil {
		return node
	} else {
		return avl.minNode(node.left)
	}
}

func (avl *AVL) nodeHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (avl *AVL) calcNodeHeight(node *Node) int {
	if node == nil {
		return 0
	}
	ld := avl.calcNodeHeight(node.left)
	rd := avl.calcNodeHeight(node.right)
	return common.MaxInt(ld, rd) + 1
}

func (avl *AVL) getNode(node *Node, key string) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		return avl.getNode(node.left, key)
	} else if key > node.key {
		return avl.getNode(node.right, key)
	} else {
		return node
	}
}

func (avl *AVL) putNode(node *Node, key, value string) *Node {
	if node == nil {
		return NewNode(key, value, 1)
	}

	if key < node.key {
		node.left = avl.putNode(node.left, key, value)
	} else if key > node.key {
		node.right = avl.putNode(node.right, key, value)
	} else {
		node.value = value
	}
	node.height = common.MaxInt(avl.nodeHeight(node.left), avl.nodeHeight(node.right)) + 1
	avl.rebalancedNode(node)
	return node
}

func (avl *AVL) deleteMinNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.left == nil {
		return node.right
	} else {
		node.left = avl.deleteMinNode(node.left)
		node.height = common.MaxInt(avl.nodeHeight(node.left), avl.nodeHeight(node.right)) + 1
		avl.rebalancedNode(node)
		return node
	}
}

func (avl *AVL) deleteNode(node *Node, key string) *Node {
	if node == nil {
		return nil
	}

	if key > node.key {
		node.right = avl.deleteNode(node.right, key)
	} else if key < node.key {
		node.left = avl.deleteNode(node.left, key)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		min := avl.minNode(node.right)
		min.left = node.left
		min.right = avl.deleteMinNode(node.right)
		node = min
	}
	node.height = common.MaxInt(avl.nodeHeight(node.left), avl.nodeHeight(node.right)) + 1
	avl.rebalancedNode(node)
	return node
}

func (avl *AVL) getBalanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return avl.nodeHeight(node.left) - avl.nodeHeight(node.right)
}

func (avl *AVL) rightRotationNode(Y *Node) *Node {
	X := Y.left
	T2 := X.right
	Y.left = T2
	X.right = Y
	Y.height = common.MaxInt(avl.nodeHeight(Y.left), avl.nodeHeight(Y.right))
	X.height = common.MaxInt(avl.nodeHeight(X.left), avl.nodeHeight(X.right))
	return X
}

func (avl *AVL) leftRotationNode(X *Node) *Node {
	Y := X.right
	T2 := Y.left
	X.right = T2
	Y.left = X
	X.height = common.MaxInt(avl.nodeHeight(X.left), avl.nodeHeight(X.right))
	Y.height = common.MaxInt(avl.nodeHeight(Y.left), avl.nodeHeight(Y.right))
	return Y
}

func (avl *AVL) rebalancedNode(node *Node) {
	balanceFactor := avl.getBalanceFactor(node)

	// 左边树高
	if balanceFactor > 1 {
		leftBalanceFactor := avl.getBalanceFactor(node.left)
		if leftBalanceFactor > 0 {
			// case1
			avl.rightRotationNode(node)
		} else if leftBalanceFactor < 0 {
			// case2
			avl.leftRotationNode(node.left)
			avl.rightRotationNode(node)
		}
	} else if balanceFactor < -1 {
		rightBalanceFactor := avl.getBalanceFactor(node.right)
		if rightBalanceFactor < 0 {
			// case4
			avl.leftRotationNode(node)
		} else if rightBalanceFactor > 0 {
			// case3
			avl.rightRotationNode(node.right)
			avl.leftRotationNode(node)
		}
	}
}

func (avl *AVL) prettifyNode(nodes []*Node, level, maxLevel int, s *string) {
	if len(nodes) < 1 || isAllNodesNil(nodes) {
		return
	}

	floor := maxLevel - level
	endgeLines := int(math.Pow(2, math.Max(float64(floor-1), 0)))
	firstSpaces := int(math.Pow(2, float64(floor))) - 1
	betweenSpaces := int(math.Pow(2, float64(floor+1))) - 1

	common.AddSpaces(firstSpaces, s)

	var newNodes []*Node
	for _, node := range nodes {
		if node != nil {
			newNodes = append(newNodes, node.left)
			newNodes = append(newNodes, node.right)
			*s += node.key
		} else {
			newNodes = append(newNodes, (*Node)(nil))
			newNodes = append(newNodes, (*Node)(nil))
			*s += " "
		}
		common.AddSpaces(betweenSpaces, s)
	}

	*s += "\n"
	// 画 / \ 树枝

	for i := 1; i <= endgeLines; i++ {
		for j := 0; j < len(nodes); j++ {
			common.AddSpaces(firstSpaces-i, s)
			if nodes[j] == nil {
				common.AddSpaces(endgeLines+endgeLines+i+1, s)
				continue
			}

			if nodes[j].left != nil {
				*s += "/"
			} else {
				*s += " "
			}

			common.AddSpaces(i+i-1, s)

			if nodes[j].right != nil {
				*s += "\\"
			} else {
				*s += " "
			}

			common.AddSpaces(endgeLines+endgeLines-i, s)
		}

		*s += "\n"
	}

	avl.prettifyNode(newNodes, level+1, maxLevel, s)
}
