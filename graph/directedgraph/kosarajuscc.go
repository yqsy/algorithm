package directedgraph

type KosarajuSCC struct {
	marked []bool // 已访问过的顶点
	id     []int  // 强连通分量的标识符
	count  int    // 强连通分量的数量
}

func NewKosarajuSCC(g *Digraph) *KosarajuSCC {
	k := &KosarajuSCC{}
	k.marked = make([]bool, g.V)
	k.id = make([]int, g.V)

	d := NewDepthFirstOrder(g.Reverse())
	length := d.reversePost.Len()
	order := make([]int, 0)
	for i := 0; i < length; i++ {
		ele := d.reversePost.Pop().(int)
		order = append(order, ele)
	}

	for _, s := range order {
		if !k.marked[s] {
			k.dfs(g, s)
			k.count++
		}
	}

}

func (k *KosarajuSCC) dfs(g *Digraph, v int) {
	k.marked[v] = true
	k.id[v] = k.count

	for _, w := range g.adj[v] {
		if !k.marked[w] {
			k.dfs(g, w)
		}
	}
}

func (k *KosarajuSCC) StronglyConnected(v, w int) bool {
	return k.id[v] == k.id[w]
}

func (k *KosarajuSCC) Id(v int) int {
	return k.id[v]
}

func (k *KosarajuSCC) Count() int {
	return k.count
}
