package unionfind

import (
	"fmt"
	"io"
)

type WeightedQuickUnionUF struct {
	id    []int // 父链接数组(由触点索引)
	sz    []int // (由触点索引的)各个根节点所对应的分量的大小
	count int   // 连通分量的数量
}

func NewWeightedQuickUnionUF(N int) *WeightedQuickUnionUF {
	wuf := &WeightedQuickUnionUF{}
	wuf.count = N
	wuf.id = make([]int, N)
	for i := 0; i < N; i++ {
		wuf.id[i] = i
	}
	wuf.sz = make([]int, N)
	for i := 0; i < N; i++ {
		wuf.sz[i] = 1
	}
	return wuf
}

func NewWeightedQuickUnionUFFromBufio(r io.Reader) *WeightedQuickUnionUF {
	var N int
	if _, err := fmt.Fscanf(r, "%d\n", &N); err != nil {
		panic(err)
	}

	qu := NewWeightedQuickUnionUF(N)

	for {
		var p, q int
		if _, err := fmt.Fscanf(r, "%d %d\n", &p, &q); err != nil {
			break
		}

		if qu.Connected(p, q) {
			continue
		}

		qu.Union(p, q)

		fmt.Printf("%v %v\n", p, q)
	}

	fmt.Printf("%v components\n", qu.Count())
	return qu
}

func (wuf *WeightedQuickUnionUF) Count() int {
	return wuf.count
}

func (wuf *WeightedQuickUnionUF) Connected(p, q int) bool {
	return wuf.Find(q) == wuf.Find(p)
}

func (wuf *WeightedQuickUnionUF) Find(p int) int {
	for p != wuf.id[p] {
		p = wuf.id[p]
	}
	return p
}

func (wuf *WeightedQuickUnionUF) Union(p, q int) {
	i := wuf.Find(p)
	j := wuf.Find(q)

	if i == j {
		return
	}

	// 将小树的根节点连接到大树的根节点
	if wuf.sz[i] < wuf.sz[j] {
		wuf.id[i] = j
		wuf.sz[j] += wuf.sz[i]
	} else {
		wuf.id[j] = i
		wuf.sz[i] += wuf.sz[j]
	}

	wuf.count--
}
