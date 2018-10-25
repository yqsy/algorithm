package unionfind

import (
	"fmt"
	"io"
)

type QuickFind struct {
	id    []int // 分量id (以触点作为索引)
	count int   // 分量数量
}

func NewQuickFind(N int) *QuickFind {
	qf := &QuickFind{}
	qf.count = N
	qf.id = make([]int, N)

	for i := 0; i < N; i++ {
		qf.id[i] = i
	}
	return qf
}

func NewQuickFindFromBufio(r io.Reader) *QuickFind {
	var N int
	if _, err := fmt.Fscanf(r, "%d\n", &N); err != nil {
		panic(err)
	}

	qf := NewQuickFind(N)

	for {
		var p, q int
		if _, err := fmt.Fscanf(r, "%d %d\n", &p, &q); err != nil {
			break
		}

		if qf.Connected(p, q) {
			continue
		}

		qf.Union(p, q)

		fmt.Printf("%v %v\n", p, q)
	}

	fmt.Printf("%v components\n", qf.Count())
	return qf
}

func (qf *QuickFind) Count() int {
	return qf.count
}

func (qf *QuickFind) Connected(p, q int) bool {
	return qf.Find(p) == qf.Find(q)
}

func (qf *QuickFind) Find(p int) int {
	return qf.id[p]
}

func (qf *QuickFind) Union(p, q int) {
	// 将p和q归并到相同的分量中
	pID := qf.Find(p)
	qID := qf.Find(q)

	// 如果p和q已经在相同的分量中则不需要采取任何行动
	if pID == qID {
		return
	}

	// 将p的分量重命名为q的名称
	for i := 0; i < len(qf.id); i++ {
		if qf.id[i] == pID {
			qf.id[i] = qID
		}
	}

	qf.count--
}
