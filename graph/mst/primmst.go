package mst

import (
	"container/heap"
	"math"
)

type PrimMST struct {
	edgeTo []Edge              // 距离树最近的边
	distTo []float64           // distTo[w]=edgeTo[w].weight()
	marked []bool              // 如果v在树中则为true
	pq     DoublePriorityQueue // 有效的横切边
}

func NewPrimMST(g *EdgeWeightedGraph) *PrimMST {
	mst := &PrimMST{}
	mst.edgeTo = make([]Edge, g.V)
	mst.distTo = make([]float64, g.V)
	mst.marked = make([]bool, g.V)

	for v := 0; v < g.V; v++ {
		mst.distTo[v] = math.MaxFloat64
	}

	// 初始化p.pq

	mst.distTo[0] = 0.0
	heap.Push(&mst.pq, WrapFloat64{v: 0.0})

	// TODO
	//for mst.pq.Len() > 0 {
	//	mst.visit(g , mst.pq.Pop().())
	//}
}

func (mst *PrimMST) visit(g *EdgeWeightedGraph, v int) {
	mst.marked[v] = true

}
