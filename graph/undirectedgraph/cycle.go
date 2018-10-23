package undirectedgraph

type Cycle struct {
	marked   []bool
	hasCycle bool
}

func NewCycle(g *Graph) *Cycle {
	c := &Cycle{}
	c.marked = make([]bool, g.V)

	for s := 0; s < g.V; s++ {
		if !c.marked[s] {
			c.dfs(g, s, s)
		}
	}

	return c
}

func (c *Cycle) dfs(g *Graph, v, u int) {
	c.marked[v] = true

	// v: 当前要遍历的节点
	// u: 遍历的节点的上一个节点
	// w: 当前遍历的节点的 near的元素中的一个
	// 判断: 在已经遍历过w节点时,判断w节点是否是u节点(遍历的节点的上一个节点)

	for _, w := range g.adj[v] {
		if !c.marked[w] {
			c.dfs(g, w, v)
		} else if w != u {
			c.hasCycle = true
		}
	}
}

func (c *Cycle) HasCycle() bool {
	return c.hasCycle
}
