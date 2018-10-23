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

		length := d.reversePost.Len()

		order := make([]int, 0)
		for i := 0; i < length; i++ {
			ele := d.reversePost.Pop().(int)
			order = append(order, ele)
		}

		if len(order) > 0 {
			t.order = order
		}
	}
	return t
}
