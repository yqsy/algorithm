package directedgraph

type Topological struct {
	order []int
}

func NewTopological(g *Digraph) *Topological {
	t := &Topological{}
	t.order = nil

	cycleFinder := NewDirectedCycle(g)
	if !cycleFinder.HasCycle() {
		d := NewDepthFirstOrder(g)

		order := d.ReversePost()

		if len(order) > 0 {
			t.order = order
		}
	}
	return t
}
