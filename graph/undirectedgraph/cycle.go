package undirectedgraph

import "fmt"

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
	// 判断: 在已经dfs过的w节点时,判断w节点是否是u节点(遍历的节点的上一个节点)

	fmt.Printf("v: %v u: %v\n", v, u)

	for _, w := range g.adj[v] {
		if !c.marked[w] {
			c.dfs(g, w, v)
		} else if w != u { // w: near中要遍历的节点 u: 当前v的来源节点, 一致时不判断cycle
			// (两个节点的互相连通不判断为cycle)
			// 在经过一个w节点,已经被dfs过了,且不是来源dfs,那么就是cycle的
			fmt.Printf("has cycle: w: %v u: %v\n", w, u)
			c.hasCycle = true
		}
	}
}

func (c *Cycle) HasCycle() bool {
	return c.hasCycle
}
