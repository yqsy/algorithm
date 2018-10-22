package undirectedgraph

type CC struct {
	marked []bool // 这个顶点上调用过dfs()了吗?
	id     []int  // 指定顶点是属于哪一个连通分量的
	count  int    // 当前的连通分量计数
}

func NewCC(g *Graph) *CC {
	cc := &CC{}
	cc.marked = make([]bool, g.V)
	cc.id = make([]int, g.V)

	for s := 0; s < g.V; s++ {
		if !cc.marked[s] {
			cc.dfs(g, s)
			cc.count++
		}
	}
	return cc
}

func (cc *CC) dfs(g *Graph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	for _, w := range g.adj[v] {
		if !cc.marked[w] {
			cc.dfs(g, w)
		}
	}
}

func (cc *CC) Connected(v, w int) bool {
	return cc.id[v] == cc.id[w]
}

func (cc *CC) Id(v int) int {
	return cc.id[v]
}

func (cc *CC) Count() int {
	return cc.count
}
