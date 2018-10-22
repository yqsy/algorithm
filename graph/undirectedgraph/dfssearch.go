package undirectedgraph

type DepthFirstSearch struct {
	marked []bool
	count  int
}

func NewDepthFirstSearch(g *Graph, s int) *DepthFirstSearch {
	d := &DepthFirstSearch{}
	d.marked = make([]bool, g.V)
	d.dfs(g, s)
	return d
}

func (d *DepthFirstSearch) dfs(g *Graph, v int) {
	d.marked[v] = true
	d.count++

	for _, w := range g.adj[v] {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}

func (d *DepthFirstSearch) Count() int {
	return d.count
}
