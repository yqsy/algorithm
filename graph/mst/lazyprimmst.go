package mst

import (
	"container/heap"
	"github.com/golang-collections/collections/queue"
)

type LazyPrimMST struct {
	marked []bool            // 最小生成树的顶点
	mst    *queue.Queue      // 最小生成树的边
	pq     EdgePriorityQueue // 横切边(包括失效的边)
}

func NewLazyPrimMST(g *EdgeWeightedGraph) *LazyPrimMST {
	mst := &LazyPrimMST{}
	mst.marked = make([]bool, g.V)
	mst.mst = queue.New()

	mst.visit(g, 0)

	for mst.pq.Len() > 0 {
		// 最小的边
		e := heap.Pop(&mst.pq).(*Edge)

		v := e.Either()
		w := e.Other(v)

		// 跳过失效的边
		if mst.marked[v] && mst.marked[w] {
			continue
		}

		// 边添加到树中
		mst.mst.Enqueue(e)

		// 顶点添加到树中
		if !mst.marked[v] {
			mst.visit(g, v)
		}

		if !mst.marked[w] {
			mst.visit(g, w)
		}
	}

	return mst
}

func (mst *LazyPrimMST) visit(g *EdgeWeightedGraph, v int) {
	mst.marked[v] = true

	for _, e := range g.adj[v] {
		if !mst.marked[e.Other(v)] {
			// 邻居节点全部放到优先级队列里
			heap.Push(&mst.pq, e)
		}
	}
}

func (mst *LazyPrimMST) Edges() []*Edge {
	eles := make([]*Edge, 0)
	for mst.mst.Len() > 0 {
		eles = append(eles, mst.mst.Dequeue().(*Edge))
	}
	return eles
}
