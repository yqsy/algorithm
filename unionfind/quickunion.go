package unionfind

import (
	"fmt"
	"io"
)

type QuickUnion struct {
	id    []int // 分量id (以触点作为索引)
	count int   // 分量数量
}

func NewQuickUnion(N int) *QuickUnion {
	qu := &QuickUnion{}
	qu.count = N
	qu.id = make([]int, N)

	for i := 0; i < N; i++ {
		qu.id[i] = i
	}
	return qu
}

func NewQuickUnionFromBufio(r io.Reader) *QuickUnion {
	var N int
	if _, err := fmt.Fscanf(r, "%d\n", &N); err != nil {
		panic(err)
	}

	qu := NewQuickUnion(N)

	for {
		var p, q int
		if _, err := fmt.Fscanf(r, "%d %d\n", &p, &q); err != nil {
			break
		}

		if qu.Connected(p, q) {
			continue
		}

		qu.union(p, q)

		fmt.Printf("%v %v\n", p, q)
	}

	fmt.Printf("%v components\n", qu.Count())
	return qu
}

func (qu *QuickUnion) Count() int {
	return qu.count
}

func (qu *QuickUnion) Connected(p, q int) bool {
	return qu.find(p) == qu.find(q)
}

func (qu *QuickUnion) find(p int) int {
	for p != qu.id[p] {
		p = qu.id[p]
	}
	return p
}

func (qu *QuickUnion) union(p, q int) {
	pRoot := qu.find(p)
	qRoot := qu.find(q)

	if pRoot == qRoot {
		return
	}

	qu.id[pRoot] = qRoot

	qu.count--
}
