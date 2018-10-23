package directedgraph

type DirectedDFS struct {
	marked []bool
}

func NewDirectedDFS(g *Digraph, s int) *DirectedDFS {
	dfs := &DirectedDFS{}
	dfs.marked = make([]bool, g.V)
	dfs.dfs(g, s)
	return dfs
}
func NewDirectedDFSFromSources(g *Digraph, sources []int) *DirectedDFS {
	d := &DirectedDFS{}
	d.marked = make([]bool, g.V)

	for _, s := range sources {
		if !d.marked[s] {
			d.dfs(g, s)
		}
	}

	return d
}

func (d *DirectedDFS) dfs(g *Digraph, v int) {
	d.marked[v] = true
	for _, w := range g.adj[v] {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}
