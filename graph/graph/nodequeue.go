package graph

type NodeQueue struct {
	items []Node
}

func NewNodeQueue() *NodeQueue {
	q := &NodeQueue{}
	q.items = []Node{}
	return q
}

func (q *NodeQueue) Enqueue(n Node) {
	q.items = append(q.items, n)
}

func (q *NodeQueue) Dequeue() *Node {
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	return &item
}

func (q *NodeQueue) Front() *Node {
	item := q.items[0]
	return &item
}

func (q *NodeQueue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *NodeQueue) Size() int {
	return len(q.items)
}
