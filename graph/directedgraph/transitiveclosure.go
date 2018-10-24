package directedgraph

type TransitiveClosure struct {
	all []*DirectedDFS
}

func NewTransitiveClosure(g *Digraph) *TransitiveClosure {
	t := &TransitiveClosure{}

	t.all = make([]*DirectedDFS, 0)

	for v := 0; v < g.V; v++ {
		t.all[v] = NewDirectedDFS(g, v)
	}

	return t
}

func (t *TransitiveClosure) Reachable(v, w int) bool {
	return t.all[v].marked[w]
}
