package graph

type NodeStack struct {
	items []Node
}

func NewNodeStack() *NodeStack {
	s := &NodeStack{}
	s.items = []Node{}
	return s
}

func (s *NodeStack) Push(n Node) {
	s.items = append(s.items, n)
}

func (s *NodeStack) Pop() *Node {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return &item
}

func (s *NodeStack) Front() *Node {
	item := s.items[len(s.items)-1]
	return &item
}

func (s *NodeStack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *NodeStack) Size() int {
	return len(s.items)
}


