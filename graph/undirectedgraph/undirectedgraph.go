package undirectedgraph

import (
	"fmt"
	"io"
	"strconv"
)

type Graph struct {
	V   int     // 顶点数量
	E   int     // 边数量
	adj [][]int // 邻接表  v:[nears:[],[],[]]
}

func NewGraph(V int) *Graph {
	g := &Graph{}
	g.V = V
	g.E = 0
	g.adj = [][]int{}
	for i := 0; i < V; i++ {
		g.adj = append(g.adj, []int{})
	}
	return g
}

func NewGraphFromBufio(r io.Reader) *Graph {
	var V, E int
	if _, err := fmt.Fscanf(r, "%d\n", &V); err != nil {
		panic(err)
	}

	if _, err := fmt.Fscanf(r, "%d\n", &E); err != nil {
		panic(err)
	}

	g := NewGraph(V)

	for i := 0; i < E; i++ {
		var v, w int
		if _, err := fmt.Fscanf(r, "%d %d\n", &v, &w); err != nil {
			panic(err)
		}
		g.AddEdge(v, w)
	}
	return g
}

func (g *Graph) AddEdge(v, w int) {
	g.adj[v] = append([]int{w}, g.adj[v]...)
	g.adj[w] = append([]int{v}, g.adj[w]...)
	g.E++
}

// 计算v的度数
func (g *Graph) Degree(v int) int {
	return len(g.adj[v])
}

// 计算所有顶点的最大度数
func (g *Graph) MaxDegree() int {
	max := 0
	for v := 0; v < g.V; v++ {
		if g.Degree(v) > max {
			max = g.Degree(v)
		}
	}
	return max
}

// 计算所有顶点的平均度数 (边意味着两个顶点都有相应的adj, adj[0] += 1 , adj[1] += 0)
func (g *Graph) AvgDegree() float64 {
	return 2.0 * float64(g.E) / float64(g.V)
}

// 计算自环的个数 (自环会在一个顶点在增加两个相应的ajd, adj[0] += 0, adj[0] += 0)
func (g *Graph) NumberOfSelfLoops() int {
	count := 0
	for v := 0; v < g.V; v++ {
		for _, w := range g.adj[v] {
			if w == v {
				count++
			}
		}
	}
	return count / 2
}

func (g *Graph) String() string {
	s := fmt.Sprintf("%v vertices, %v edges\n", g.V, g.E)

	for v := 0; v < g.V; v++ {
		s += strconv.Itoa(v) + ": "

		for _, w := range g.adj[v] {
			s += strconv.Itoa(w) + " "
		}
		s += "\n"
	}

	return s
}
