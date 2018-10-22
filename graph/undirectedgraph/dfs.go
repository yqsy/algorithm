package undirectedgraph

type DFS struct {
	marked []bool
	count  int
}

func NewDFS(g *Graph, s int) *DFS{
	d := &DFS{}
	d.marked = make([]bool, g.V)
	d.dfs(g, s)
	return d
}

func (d *DFS) Count() int {
	return d.count
}

func (d *DFS) dfs(g *Graph, v int) {
	d.marked[v] = true
	d.count++
	near := g.adj[v]
	for w := 0; w < len(near); w++ {
		if !d.marked[near[w]] {
			d.dfs(g, near[w])
		}
	}
}
