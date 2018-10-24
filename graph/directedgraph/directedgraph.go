package directedgraph

import (
	"fmt"
	"io"
	"strconv"
)

type Digraph struct {
	V   int     // 顶点数量
	E   int     // 边数量
	adj [][]int // 邻接表
}

func NewDigraph(V int) *Digraph {
	g := &Digraph{}
	g.V = V
	g.E = 0
	g.adj = [][]int{}
	for v := 0; v < V; v++ {
		g.adj = append(g.adj, []int{})
	}
	return g
}

func NewDigraphFromBufio(r io.Reader) *Digraph {
	var V, E int
	if _, err := fmt.Fscanf(r, "%d\n", &V); err != nil {
		panic(err)
	}

	if _, err := fmt.Fscanf(r, "%d\n", &E); err != nil {
		panic(err)
	}

	g := NewDigraph(V)

	for i := 0; i < E; i++ {
		var v, w int
		if _, err := fmt.Fscanf(r, "%d %d\n", &v, &w); err != nil {
			panic(err)
		}
		g.AddEdge(v, w)
	}
	return g
}

func (g *Digraph) AddEdge(v, w int) {
	g.adj[v] = append([]int{w}, g.adj[v]...)
	g.E++
}

func (g *Digraph) Reverse() *Digraph {
	R := NewDigraph(g.V)

	for v := 0; v < g.V; v++ {
		for _, w := range g.adj[v] {
			R.AddEdge(w, v)
		}
	}
	return R
}

func (g *Digraph) String() string {
	s := fmt.Sprintf("%v vertices, %v edges\n", g.V, g.E)

	for v := 0; v < g.V; v++ {
		s += strconv.Itoa(v) + ": "

		for _, w := range g.adj[v] {
			s += strconv.Itoa(w)
			s += " "
		}
		s += "\n"
	}

	return s
}
