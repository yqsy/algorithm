package undirectedgraph

import (
	"bufio"
	"fmt"
	"strconv"
)

type Graph struct {
	V   int           // 顶点数量
	E   int           // 边数量
	adj map[int][]int // 邻接表
}

func NewGraph(V int) *Graph {
	graph := &Graph{}
	graph.V = V
	graph.E = 0
	graph.adj = make(map[int][]int)
	for i := 0; i < V; i++ {
		graph.adj[i] = []int{}
	}
	return graph
}

func NewGraphFromBufio(r *bufio.Reader) *Graph {
	var V, E int
	if _, err := fmt.Fscanf(r, "%d\n", &V); err != nil {
		panic(err)
	}

	if _, err := fmt.Fscanf(r, "%d\n", &E); err != nil {
		panic(err)
	}

	graph := NewGraph(V)

	for i := 0; i < E; i++ {
		var v, w int
		if _, err := fmt.Fscanf(r, "%d %d\n", &v, &w); err != nil {
			panic(err)
		}
		graph.AddEdge(v, w)
	}
	return graph
}

func (g *Graph) AddEdge(v, w int) {
	g.adj[v] = append([]int{w}, g.adj[v]...)
	g.adj[w] = append([]int{v}, g.adj[w]...)
	g.E++
}

// 计算度数
func (g *Graph) Degree(v int) int {
	return len(g.adj[v])
}

// 计算最大度数
func (g *Graph) MaxDegree() int {
	max := 0
	for k := range g.adj {
		if g.Degree(k) > max {
			max = g.Degree(k)
		}
	}
	return max
}



func (g *Graph) String() string {
	s := fmt.Sprintf("%v vertices, %v edges\n", g.V, g.E)

	for v := 0; v < g.V; v++ {
		s += strconv.Itoa(v) + ": "

		near := g.adj[v]

		for w := 0; w < len(near); w++ {
			s += strconv.Itoa(near[w]) + " "
		}
		s += "\n"
	}

	return s
}
