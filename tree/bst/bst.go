package bst

type Node struct {
	key   string
	value string

	left, right *Node

	// 以该结点为根的子树中的结点总数(包括自身结点)
	n int
}

func NewNode(key string, value string) *Node {
	return &Node{key: key, value: value}
}

type BST struct {
	head *Node
}

func (bst *BST) NodeSize(node *Node) int {
	if node == nil {
		return 0
	} else {
		return node.n
	}
}

func (bst *BST) Size() int {
	if bst.head == nil {
		return 0
	} else {
		return bst.NodeSize(bst.head)
	}
}

func (bst *BST) Get(key string) (string, bool) {
	node := bst.getNode(bst.head, key)
	if node != nil {
		return node.value, true
	} else {
		return "", false
	}
}

func (bst *BST) Put(key, value string) {
	bst.head = bst.putNode(bst.head, key, value)
}

// 最小key
func (bst *BST) Min() (string, bool) {
	minNode := bst.minNode(bst.head)
	if minNode != nil {
		return minNode.key, true
	} else {
		return "", false
	}
}

// 最大key
func (bst *BST) Max() (string, bool) {
	maxNode := bst.maxNode(bst.head)
	if maxNode != nil {
		return maxNode.key, true
	} else {
		return "", false
	}
}

// 向下取整(寻找<=key的最大key)
func (bst *BST) Floor(key string) (string, bool) {
	floorNode := bst.floorNode(bst.head, key)
	if floorNode != nil {
		return floorNode.key, true
	} else {
		return "", false
	}
}

// 向上取整(寻找>=key的最小key)
func (bst *BST) Ceiling(key string) (string, bool) {
	ceilingNode := bst.ceilingNode(bst.head, key)
	if ceilingNode != nil {
		return ceilingNode.key, true
	} else {
		return "", false
	}
}

// 返回排名为k的节点, 排名范围[0,k]
func (bst *BST) Select(k int) (string, bool) {
	selectNode := bst.selectNode(bst.head, k)
	if selectNode != nil {
		return selectNode.key, true
	} else {
		return "", false
	}
}

// 返回 < key 的键的数量
func (bst *BST) Rank(key string) (int, bool) {
	return bst.rankNode(bst.head, key)
}

// 删除最小键
func (bst *BST) DeleteMin() {
	bst.head = bst.deleteMinNode(bst.head)
}

// 删除最大键
func (bst *BST) DeleteMax() {
	bst.head = bst.deleteMaxNode(bst.head)
}

// 删除键
func (bst *BST) Delete(key string) {
	bst.head = bst.deleteNode(bst.head, key)
}

// 范围查找
func (bst *BST) Keys(lo, hi string) []string {
	var result []string
	bst.keysNode(bst.head, &result, lo, hi)
	return result
}

func (bst *BST) putNode(node *Node, key, value string) *Node {
	if node == nil {
		return NewNode(key, value)
	}

	if key < node.key {
		node.left = bst.putNode(node.left, key, value)
	} else if key > node.key {
		node.right = bst.putNode(node.right, key, value)
	} else {
		node.value = value
	}

	return node
}

func (bst *BST) getNode(node *Node, key string) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		return bst.getNode(node.left, key)
	} else if key > node.key {
		return bst.getNode(node.right, key)
	} else {
		return node
	}
}

func (bst *BST) minNode(node *Node) (*Node) {
	if node == nil {
		return nil
	}

	if node.left == nil {
		return node
	} else {
		return bst.minNode(node.left)
	}
}

func (bst *BST) maxNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.right == nil {
		return node
	} else {
		return bst.maxNode(node.right)
	}
}

func (bst *BST) floorNode(node *Node, key string) *Node {
	if node == nil {
		return nil
	}

	if node.key == key {
		return node
	} else if node.key < key {
		// is current largest?
		if moreNode := bst.floorNode(node.right, key); moreNode != nil {
			return moreNode
		} else {
			return node
		}
	} else {
		return bst.floorNode(node.left, key)
	}
}

func (bst *BST) ceilingNode(node *Node, key string) *Node {
	if node == nil {
		return nil
	}

	if node.key == key {
		return node
	} else if node.key < key {
		return bst.ceilingNode(node.right, key)
	} else {
		// is current smallest?
		if moreNode := bst.ceilingNode(node.left, key); moreNode != nil {
			return moreNode
		} else {
			return node
		}
	}
}

func (bst *BST) selectNode(node *Node, k int) *Node {
	if node == nil {
		return nil
	}

	if bst.NodeSize(node.left) == k {
		return node
	} else if bst.NodeSize(node.left) > k {
		return bst.selectNode(node.left, k)
	} else {
		return bst.selectNode(node.right, k-bst.NodeSize(node.left)-1)
	}
}

func (bst *BST) rankNode(node *Node, key string) (int, bool) {
	if node == nil {
		return 0, false
	}

	if key == node.key {
		return bst.NodeSize(node.left), true
	} else if key < node.key {
		return bst.rankNode(node.left, key)
	} else {
		moreRank, _ := bst.rankNode(node.right, key)
		return 1 + bst.NodeSize(node.left) + moreRank, true
	}
}

func (bst *BST) deleteMinNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.left == nil {
		return node.right
	} else {
		node.left = bst.deleteMinNode(node.left)
		node.n = bst.NodeSize(node.left) + bst.NodeSize(node.right) + 1
		return node
	}
}

func (bst *BST) deleteMaxNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.right == nil {
		return node.left
	} else {
		node.right = bst.deleteMaxNode(node.right)
		node.n = bst.NodeSize(node.left) + bst.NodeSize(node.right) + 1
		return node
	}
}

func (bst *BST) deleteNode(node *Node, key string) *Node {
	if node == nil {
		return nil
	}

	if key > node.key {
		node.right = bst.deleteNode(node.right, key)
	} else if key < node.key {
		node.left = bst.deleteNode(node.left, key)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		min := bst.minNode(node.right)
		min.left = node.left
		min.right = bst.deleteMinNode(node.right)
		node = min
	}

	node.n = bst.NodeSize(node.left) + bst.NodeSize(node.right) + 1
	return node
}

func (bst *BST) keysNode(node *Node, result *[]string, lo, hi string) {
	if node == nil {
		return
	}

	if lo < node.key {
		bst.keysNode(node.left, result, lo, hi)
	}

	if lo <= node.key && node.key <= hi {
		*result = append(*result, node.key)
	}

	if hi > node.key {
		bst.keysNode(node.right, result, lo, hi)
	}
}
