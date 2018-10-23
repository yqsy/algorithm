package directedgraph

import "github.com/golang-collections/collections/stack"

type DirectedCycle struct {
	marked  []bool
	edgeTo  []int
	cycle   *stack.Stack // 有向环中的所有顶点
	onStack []bool       // 递归调用的栈上的所有顶点
}

func NewDirectedCycle(g *Digraph) *DirectedCycle {
	d := &DirectedCycle{}
	d.marked = make([]bool, g.V)
	d.edgeTo = make([]int, g.V)
	d.onStack = make([]bool, g.V)
	d.cycle = nil

	for v := 0; v < g.V; v++ {
		if !d.marked[v] {
			d.dfs(g, v)
		}
	}
	return d
}

func (d *DirectedCycle) dfs(g *Digraph, v int) {
	d.onStack[v] = true
	d.marked[v] = true

	for _, w := range g.adj[v] {
		if d.HasCycle() {
			return
		} else if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		} else if d.onStack[w] {
			d.cycle = stack.New()
			for x := v; x != w; x = d.edgeTo[x] {
				d.cycle.Push(x)
			}

			d.cycle.Push(w)
			d.cycle.Push(v)
		}
	}
	d.onStack[v] = false
}

func (d *DirectedCycle) HasCycle() bool {
	return d.cycle != nil
}
