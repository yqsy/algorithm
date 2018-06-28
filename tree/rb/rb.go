package rb

const (
	RED   = true
	BLACK = false
)

type Node struct {
	key   int
	value string

	left, right *Node

	// 以该结点为根的子树中的结点总数(包括自身结点)
	n int

	// 其父结点指向它的链接的颜色
	color bool
}

func NewNode(key int, value string, n int, color bool) *Node {
	return &Node{key: key, value: value, n: n, color: color}
}

type RB struct {
	head *Node
}

func (rb *RB) NodeSize(node *Node) int {
	if node == nil {
		return 0
	} else {
		return node.n
	}
}

func (rb *RB) rotateLeft(E *Node) *Node {
	S := E.right
	T2 := S.left
	E.right = T2
	S.left = E
	S.color = E.color
	E.color = RED
	S.n = E.n
	E.n = rb.NodeSize(E.left) + rb.NodeSize(E.right) + 1
	return S
}

func (rb *RB) rotateRight(S *Node) *Node {
	E := S.left
	T2 := E.right
	S.left = T2
	E.right = S
	E.color = S.color
	S.color = RED
	E.n = S.n
	S.n = rb.NodeSize(S.left) + rb.NodeSize(S.right) + 1
	return E
}

func (rb *RB) flipColors(node *Node) {
	node.color = RED
	node.left.color = BLACK
	node.right.color = BLACK
}

func (rb *RB) isRed(node *Node) bool {
	if node == nil {
		return false
	}
	return node.color == RED
}

func (rb *RB) putNode(node *Node, key int, value string) *Node {
	if node == nil {
		return NewNode(key, value, 1, RED)
	}

	if key < node.key {
		node.left = rb.putNode(node.left, key, value)
	} else if key > node.key {
		node.right = rb.putNode(node.right, key, value)
	} else {
		node.value = value
	}

	if rb.isRed(node.right) && !rb.isRed(node.left) {
		node = rb.rotateLeft(node)
	}

	if rb.isRed(node.left) && rb.isRed(node.left.left) {
		node = rb.rotateRight(node)
	}

	if rb.isRed(node.left) && rb.isRed(node.right) {
		rb.flipColors(node)
	}

	node.n = rb.NodeSize(node.left) + rb.NodeSize(node.right) + 1
	return node
}
