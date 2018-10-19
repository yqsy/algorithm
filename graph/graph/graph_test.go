package graph

import (
	"fmt"
	"testing"
)

func GenerateGraph() *Graph {
	g := Graph{}

	n1, n2, n3, n4, n5 := Node{1}, Node{2}, Node{3}, Node{4}, Node{5}
	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddNode(&n4)
	g.AddNode(&n5)

	g.AddEdge(&n1, &n2)
	g.AddEdge(&n1, &n5)
	g.AddEdge(&n2, &n3)
	g.AddEdge(&n2, &n4)
	g.AddEdge(&n2, &n5)
	g.AddEdge(&n3, &n4)
	g.AddEdge(&n4, &n5)

	return &g
}

func TestAdd(t *testing.T) {
	g := GenerateGraph()
	fmt.Printf("%v", g.String()) // TODO ?
}

func TestNode(t *testing.T) {

	n1 := Node{1}

	n2 := Node{1}

	a := make(map[Node]string)

	a[n1] = "A"
	a[n2] = "B"

	fmt.Println(a[n1])
	fmt.Println(a[n2])
}

func TestBFS(t *testing.T) {
	g := GenerateGraph()

	g.BFS(func(n *Node) {
		fmt.Printf("%v\n", n.value)
	}, 0)
}

func TestDFS(t *testing.T) {
	g := GenerateGraph()

	g.DFS(func(n *Node) {
		fmt.Printf("%v\n", n.value)
	}, 0)
}
