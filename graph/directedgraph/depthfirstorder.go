package directedgraph

import (
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

type DepthFirstOrder struct {
	marked      []bool
	pre         *queue.Queue // 所有顶点的前序排列
	post        *queue.Queue // 所有顶点的后续排列
	reversePost *stack.Stack // 所有顶点的逆后续排序
}

func NewDepthFirstOrder(g *Digraph) *DepthFirstOrder {
	d := &DepthFirstOrder{}
	d.pre = queue.New()
	d.post = queue.New()
	d.reversePost = stack.New()
	d.marked = make([]bool, g.V)
	for v := 0; v < g.V; v++ {
		if !d.marked[v] {
			d.dfs(g, v)
		}
	}
	return d
}

func (d *DepthFirstOrder) dfs(g *Digraph, v int) {
	d.pre.Enqueue(v)

	d.marked[v] = true
	for _, w := range g.adj[v] {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}

	d.post.Enqueue(v)
	d.reversePost.Push(v)
}

func (d *DepthFirstOrder) Pre() []int {
	eles := make([]int, 0)
	length := d.pre.Len()
	for i := 0; i < length; i++ {
		eles = append(eles, d.pre.Dequeue().(int))
	}
	return eles
}

func (d *DepthFirstOrder) Post() []int {
	eles := make([]int, 0)
	length := d.post.Len()
	for i := 0; i < length; i++ {
		eles = append(eles, d.post.Dequeue().(int))
	}
	return eles
}

func (d *DepthFirstOrder) ReversePost() []int {
	eles := make([]int, 0)
	length := d.reversePost.Len()
	for i := 0; i < length; i++ {
		eles = append(eles, d.reversePost.Pop().(int))
	}
	return eles
}
