package mst

import (
	"fmt"
)

type Edge struct {
	v      int     // 顶点之一
	w      int     // 另一顶点
	weight float64 // 边的权重

	index int // 在heap中的原始index
}

func NewEdge(v, w int, weight float64) *Edge {
	edge := &Edge{}
	edge.v = v
	edge.w = w
	edge.weight = weight
	return edge
}

func (e *Edge) Weight() float64 {
	return e.weight
}

func (e *Edge) Either() int {
	return e.v
}

func (e *Edge) Other(vertex int) int {
	if vertex == e.v {
		return e.w
	} else if vertex == e.w {
		return e.v
	} else {
		panic("err")
	}
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d-%d %.2f", e.v, e.w, e.weight)
}

//func Compare(a, b *Edge) int {
//	if a.Weight() < b.Weight() {
//		return -1
//	} else if a.Weight() > b.Weight() {
//		return +1
//	} else {
//		return 0
//	}
//}

// PQ interface

type EdgePriorityQueue []*Edge

func (pq EdgePriorityQueue) Len() int {
	return len(pq)
}

func (pq EdgePriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq EdgePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *EdgePriorityQueue) Push(x interface{}) {
	n := len(*pq)
	edge := x.(*Edge)
	edge.index = n
	*pq = append(*pq, edge)
}

func (pq *EdgePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	edge := old[n-1]
	edge.index = -1
	*pq = old[0 : n-1]
	return edge
}

