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

	f, err := os.Open("tinyG3.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	g := NewDigraphFromBufio(r)

	d := NewDepthFirstOrder(g)

	fmt.Println(d.Pre())

	fmt.Println(d.Post())

	fmt.Println(d.ReversePost())

	//fmt.Println(g.String())
}

// 拓扑排序
func TestTopological(t *testing.T) {

	s := NewSymbolGraph("jobs.txt", "/")

	top := NewTopological(s.g)

	for _, v := range top.order {
		fmt.Println(s.Name(v))
	}
}

// Kosaraju 强连通算法
func TestKosarajuSCC(t *testing.T) {
	f, err := os.Open("tinyG.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	g := NewDigraphFromBufio(r)

	cc := NewKosarajuSCC(g)

	M := cc.Count()

	fmt.Printf("%v compoents\n", M)

	var compnents = make([][]int, M)

	for i := 0; i < M; i++ {
		compnents[i] = []int{}
	}

	for v := 0; v < g.V; v ++ {
		compnents[cc.Id(v)] = append([]int{v}, compnents[cc.Id(v)]...)
	}

	for i := 0; i < M; i++ {
		for _, v := range compnents[i] {
			fmt.Printf("%v ", v)
		}
		fmt.Printf("\n")
	}
}