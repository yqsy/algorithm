package graph

import (
	"fmt"
)

type Node struct {
	value int
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

type Graph struct {
	nodes []*Node          // 顶点集
	edges map[Node][]*Node // 边集
}

func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func (g *Graph) String() string {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "

		near := g.edges[*g.nodes[i]]

		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	return s
}

// 宽度优先搜索 Breadth First Search (队列)
func (g *Graph) BFS(f func(n *Node), nodeIdx int) {
	if nodeIdx >= len(g.nodes) {
		return
	}

	q := NewNodeQueue()
	firstNode := g.nodes[nodeIdx]
	q.Enqueue(*firstNode)

	visited := make(map[*Node]bool)
	visited[firstNode] = true

	for {
		if q.IsEmpty() {
			break
		}

		node := q.Dequeue()
		visited[node] = true
		nexts := g.edges[*node]

		for _, next := range nexts {
			if !visited[next] {
				q.Enqueue(*next)
				visited[next] = true
			}
		}

		if f != nil {
			f(node)
		}
	}
}

// 深度优先搜索 Depth First Search (栈)
func (g *Graph) DFS(f func(n *Node), nodeIdx int) {
	if nodeIdx >= len(g.nodes) {
		return
	}

	s := NewNodeStack()
	firstNode := g.nodes[nodeIdx]
	s.Push(*firstNode)

	visited := make(map[*Node]bool)
	visited[firstNode] = true

	for {
		if s.IsEmpty() {
			break
		}

		node := s.Pop()
		visited[node] = true
		nexts := g.edges[*node]

		for _, next := range nexts {
			if !visited[next] {
				s.Push(*next)
				visited[next] = true
			}
		}

		if f != nil {
			f(node)
		}
	}
}
