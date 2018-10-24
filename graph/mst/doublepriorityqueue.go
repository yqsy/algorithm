package mst

type WrapFloat64 struct {
	v     float64
	index int // 在heap中的原始index
}

type DoublePriorityQueue []WrapFloat64

func (pq DoublePriorityQueue) Len() int {
	return len(pq)
}

func (pq DoublePriorityQueue) Less(i, j int) bool {
	return pq[i].v < pq[j].v
}

func (pq DoublePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *DoublePriorityQueue) Push(x interface{}) {
	n := len(*pq)
	wrapFloat64 := x.(*WrapFloat64)
	wrapFloat64.index = n
	*pq = append(*pq, *wrapFloat64)
}

func (pq *DoublePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	wrapFloat64 := old[n-1]
	wrapFloat64.index = -1
	*pq = old[0 : n-1]
	return wrapFloat64
}
