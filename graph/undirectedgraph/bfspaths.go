package undirectedgraph

import "github.com/golang-collections/collections/queue"

type BreadthFirstPaths struct {
	marked []bool // 达到该顶点的最短路径已知吗?
	edgeTo []int  // 到达该顶点的已知路径上的最后一个顶点
	s      int    // 起点
}

func NewBreadthFirstPaths(g *Graph, s int) *BreadthFirstPaths {
	b := &BreadthFirstPaths{}
	b.marked = make([]bool, g.V)
	b.edgeTo = make([]int, g.V)
	b.s = s
	b.bfs(g, s)
	return b
}

func (b *BreadthFirstPaths) bfs(g *Graph, s int) {
	queue := queue.New()

	b.marked[s] = true

	queue.Enqueue(s)

	for queue.Len() != 0 {
		v := queue.Dequeue().(int)

		for _, w := range g.adj[v] {
			if !b.marked[w] {
				b.edgeTo[w] = v
				b.marked[w] = true
				queue.Enqueue(w)
			}
		}
	}
}

func (b *BreadthFirstPaths) HasPathTo(v int) bool {
	return b.marked[v]
}

func (b *BreadthFirstPaths) PathTo(v int) []int {
	if !b.HasPathTo(v) {
		return []int{}
	}

	var paths []int
	for x := v; x != b.s; x = b.edgeTo[x] {
		paths = append([]int{x}, paths...)
	}
	paths = append([]int{b.s}, paths...)
	return paths
}
