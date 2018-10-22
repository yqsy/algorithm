package undirectedgraph

type DFS struct {
	marked []bool
	count  int
}

func NewDFS(g *Graph, s int) *DFS {
	d := &DFS{}
	d.marked = make([]bool, g.V)
	d.dfs(g, s)
	return d
}

func (d *DFS) dfs(g *Graph, v int) {
	d.marked[v] = true
	d.count++

	for _, w := range g.adj[v] {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}

func (d *DFS) Count() int {
	return d.count
}
