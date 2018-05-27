package heap

type Heap []int

func NewHeap() *Heap {
	heap := Heap{}
	heap = make([]int, 1)
	// no use
	heap[0] = 0
	return &heap
}

// 元素有多少个
// 最后一个元素的下标
func (heap *Heap) N() int {
	return len(*heap) - 1
}

func (heap *Heap) swim(k int) {
	for ; k > 1 && (*heap)[k/2] < (*heap)[k]; {
		(*heap)[k], (*heap)[k/2] = (*heap)[k/2], (*heap)[k]
		k = k / 2
	}
}

func (heap *Heap) sink(k int) {
	for ; 2*k <= heap.N(); {
		j := 2 * k
		if j < heap.N() && (*heap)[j] < (*heap)[j+1] {
			j++
		}
		if !((*heap)[k] < (*heap)[j]) {
			break
		}
		(*heap)[k], (*heap)[j] = (*heap)[j], (*heap)[k]
		k = j
	}
}

func (heap *Heap) append(n int) {
	*heap = append(*heap, n)
	heap.swim(heap.N())
}

func (heap *Heap) delMax() {
	if heap.N() < 1 {
		return
	}
	// remove last ele
	last := (*heap)[heap.N()]
	*heap = (*heap)[:heap.N()]

	// sink
	(*heap)[1] = last
	heap.sink(1)
}

func (heap *Heap) snap() []int {
	if heap.N() < 1 {
		return []int{}
	}

	return (*heap)[1 : heap.N()+1]
}
