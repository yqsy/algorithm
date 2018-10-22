package undirectedgraph

type DepthFirstPaths struct {
	marked []bool // 这个顶点上调用过dfs()了吗?
	edgeTo []int  // 从起点到一个顶点的已知路径上的最后一个顶点
	s      int    // 起点
}

func NewDepthFirstPaths(g *Graph, s int) *DepthFirstPaths {
	d := &DepthFirstPaths{}
	d.marked = make([]bool, g.V)
	d.edgeTo = make([]int, g.V)
	d.s = s
	d.dfs(g, s)
	return d
}

func (d *DepthFirstPaths) dfs(g *Graph, v int) {
	d.marked[v] = true

	for _, w := range g.adj[v] {
		if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		}
	}
}

func (d *DepthFirstPaths) HasPathTo(v int) bool {
	return d.marked[v]
}

func (d *DepthFirstPaths) PathTo(v int) []int {
	if !d.HasPathTo(v) {
		return []int{}
	}

	var paths []int
	for x := v; x != d.s; x = d.edgeTo[x] {
		paths = append(paths, x)
	}
	paths = append(paths, d.s)
	return paths
}
