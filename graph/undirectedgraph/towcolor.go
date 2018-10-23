package undirectedgraph

type TwoColor struct {
	marked         []bool
	color          []bool
	isTwoColorable bool
}

func NewTwoColor(g *Graph) *TwoColor {
	t := &TwoColor{}
	t.marked = make([]bool, g.V)
	t.color = make([]bool, g.V)
	t.isTwoColorable = true

	for s := 0; s < g.V; s++ {
		if !t.marked[s] {
			t.dfs(g, s)
		}
	}
	return t
}

func (t *TwoColor) dfs(g *Graph, v int) {
	t.marked[v] = true

	for _, w := range g.adj[v] {
		if !t.marked[w] {
			t.color[w] = !t.color[v]
			t.dfs(g, w)
		} else if t.color[w] == t.color[v] {
			t.isTwoColorable = false
		}
	}
}

func (t *TwoColor) IsBipartite() bool {
	return t.isTwoColorable
}
