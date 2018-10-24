package mst

import (
	"fmt"
	"io"
)

type EdgeWeightedGraph struct {
	V   int       // 顶点总数
	E   int       // 边的总数
	adj [][]*Edge // 邻接表
}

func NewEdgeWeightedGraph(V int) *EdgeWeightedGraph {
	g := &EdgeWeightedGraph{}
	g.V = V
	g.E = 0
	g.adj = make([][]*Edge, V)
	for v := 0; v < V; v++ {
		g.adj = append(g.adj, []*Edge{})
	}
	return g
}

func NewGraphFromBufio(r io.Reader) *EdgeWeightedGraph {
	var V, E int

	if _, err := fmt.Fscanf(r, "%d\n", &V); err != nil {
		panic(err)
	}

	if _, err := fmt.Fscanf(r, "%d\n", &E); err != nil {
		panic(err)
	}

	g := NewEdgeWeightedGraph(V)

	for i := 0; i < E; i++ {
		var v, w int
		var weight float64
		if _, err := fmt.Fscanf(r, "%d %d %f\n", &v, &w, &weight); err != nil {
			panic(err)
		}
		e := NewEdge(v, w, weight)
		g.AddEdge(e)
	}

	return g
}

func (g *EdgeWeightedGraph) AddEdge(e *Edge) {
	v := e.Either()
	w := e.Other(v)
	g.adj[v] = append([]*Edge{e}, g.adj[v]...)
	g.adj[w] = append([]*Edge{e}, g.adj[w]...)
	g.E++
}
