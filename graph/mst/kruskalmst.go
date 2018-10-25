package mst

import (
	"container/heap"
	"github.com/golang-collections/collections/queue"
	"github.com/yqsy/algorithm/unionfind"
)

type KruskalMST struct {
	mst *queue.Queue
}

func NewKruskalMST(g *EdgeWeightedGraph) *KruskalMST {
	mst := &KruskalMST{}
	mst.mst = queue.New()

	pq := EdgePriorityQueue{}
	for _, e := range g.Edges() {
		heap.Push(&pq, e)
	}

	uf := unionfind.NewWeightedQuickUnionUF(g.V)

	for pq.Len() > 0 && mst.mst.Len() < g.V-1 {
		// 权重最小的边和它的顶点
		e := heap.Pop(&pq).(*Edge)
		v := e.Either()
		w := e.Other(v)

		// 忽略失效的边
		if uf.Connected(v, w) {
			continue
		}

		// 合并分量
		uf.Union(v, w)

		// 将边添加到最小生成树中
		mst.mst.Enqueue(e)
	}
	return mst
}
