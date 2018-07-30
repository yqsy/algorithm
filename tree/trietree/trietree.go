package trie

type Node struct {
	children map[byte]*Node
	value    string
}

// 作为结果列表
type out struct {
	o []string
}

func NewNode() *Node {
	return &Node{children: make(map[byte]*Node)}
}

type Trie struct {
	head *Node
}

func NewTrie() *Trie {
	return &Trie{head: NewNode()}
}

func (trie *Trie) Find(key string) string {
	return trie.findNode(trie.head, key)
}

func (trie *Trie) Insert(key string) {
	trie.insertNode(trie.head, key)
}

func (trie *Trie) FindPrefix(prefix string) []string {
	return trie.findPrefixNode(trie.head, prefix)
}

// dfs 遍历节点,提取所有有value的
func (trie *Trie) recursionFindNode(node *Node, out *out) {

	// 不可能的
	if node == nil || out == nil {
		return
	}

	if node.value != "" {
		out.o = append(out.o, node.value)
	}

	for k, v := range node.children {

		byte_ := k
		_ = byte_
		node_ := v

		if node_ != nil {
			trie.recursionFindNode(node_, out)
		}
	}
}

// 公共前缀所有value
func (trie *Trie) findPrefixNode(node *Node, prefix string) []string {
	// 到达公共前缀的顶部
	for i := 0; i < len(prefix); i++ {
		if _, ok := node.children[prefix[i]]; ok {
			node = node.children[prefix[i]]
		} else {
			return []string{}
		}
	}

	out := &out{}
	trie.recursionFindNode(node, out)
	return out.o
}

// 指定key查找
func (trie *Trie) findNode(node *Node, key string) string {
	for i := 0; i < len(key); i++ {
		if _, ok := node.children[key[i]]; ok {
			node = node.children[key[i]]
		} else {
			return ""
		}
	}
	return node.value
}

func (trie *Trie) insertNode(node *Node, key string) {
	value := key

	indexLastChar := -1

	for i := 0; i < len(key); i++ {
		if _, ok := node.children[key[i]]; ok {
			node = node.children[key[i]]
		} else {
			indexLastChar = i
			break
		}
	}

	remain := key[indexLastChar:]
	if indexLastChar != -1 {
		for i := 0; i < len(remain); i++ {
			node.children[remain[i]] = NewNode()
			node = node.children[remain[i]]
		}
	}

	node.value = value
}
