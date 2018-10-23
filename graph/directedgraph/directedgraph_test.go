package directedgraph

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// 从文件读入图
func TestSimpleString(t *testing.T) {
	f, err := os.Open("tinyG.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	g := NewDigraphFromBufio(r)

	fmt.Println(g.String())
}

// DFS
func TestSimleDfs(t *testing.T) {
	f, err := os.Open("tinyG.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	g := NewDigraphFromBufio(r)

	arr := []int{1, 2, 6}

	dfs := NewDirectedDFSFromSources(g, arr)

	for v := 0; v < g.V; v ++ {
		if dfs.marked[v] {
			fmt.Printf("%v ", v)
		}
	}
	fmt.Println()
}

// 有向图是否是有环
func TestDirectedCycle(t *testing.T) {
	f, err := os.Open("tinyG2.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	g := NewDigraphFromBufio(r)

	d := NewDirectedCycle(g)

	fmt.Println(d.HasCycle())
}

// 基于深度优先搜索的顶点排序
func TestDepthFirstOrder(t *testing.T) {

	f, err := os.Open("tinyG.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	g := NewDigraphFromBufio(r)

	d := NewDepthFirstOrder(g)

	length := d.pre.Len()
	for i := 0; i < length; i++ {
		ele := d.pre.Dequeue().(int)
		fmt.Printf("%v ", ele)
	}
	fmt.Println()

	length = d.post.Len()
	for i := 0; i < length; i++ {
		ele := d.post.Dequeue().(int)
		fmt.Printf("%v ", ele)
	}
	fmt.Println()


	length = d.reversePost.Len()
	for i := 0; i < length; i++ {
		ele := d.reversePost.Pop().(int)
		fmt.Printf("%v ", ele)
	}
	fmt.Println()


}
