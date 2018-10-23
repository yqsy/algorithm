package undirectedgraph

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	g := NewGraphFromBufio(r)

	fmt.Println(g.String())
}

// 度数
func TestSimpleDegree(t *testing.T) {

	// 1. 两个顶点 2. 一条边
	tmpbuf := `2
1
0 1
`
	r := strings.NewReader(tmpbuf)

	g := NewGraphFromBufio(r)

	if g.Degree(0) != 1 || g.Degree(1) != 1 {
		t.Fatal("err")
	}

	if g.MaxDegree() != 1 {
		t.Fatal("err")
	}

	if g.AvgDegree() != 1 {
		t.Fatal("err")
	}

	if g.NumberOfSelfLoops() != 0 {
		t.Fatal("err")
	}
}

// 自环
func TestSimpleSelfLoop(t *testing.T) {
	// 1. 一个顶点 2. 一条边
	tmpbuf := `1
1
0 0`
	r := strings.NewReader(tmpbuf)

	g := NewGraphFromBufio(r)

	if g.Degree(0) != 2 {
		t.Fatal("err")
	}

	if g.MaxDegree() != 2 {
		t.Fatal("err")
	}

	if g.AvgDegree() != 2 {
		t.Fatal("err")
	}

	if g.NumberOfSelfLoops() != 1 {
		t.Fatal("err")
	}
}

func getTmpBuf() string {
	tmpbuf := `6
8
0 5
2 4
2 3
1 2
0 1
3 4
3 5
0 2
`

	return tmpbuf
}

// 简单DFS
func TestDFS(t *testing.T) {
	tmpbuf := getTmpBuf()

	r := strings.NewReader(tmpbuf)

	g := NewGraphFromBufio(r)

	d := NewDepthFirstSearch(g, 0)

	if d.Count() != 6 {
		t.Fatal("err")
	}
}

// DFS 路径
func TestDFSPaths(t *testing.T) {
	tmpbuf := getTmpBuf()

	r := strings.NewReader(tmpbuf)

	g := NewGraphFromBufio(r)

	d := NewDepthFirstPaths(g, 0)

	for v := 0; v < g.V; v++ {
		fmt.Printf("%v to %v : ", d.s, v)

		if d.HasPathTo(v) {
			paths := d.PathTo(v)
			for _, x := range paths {
				if x == d.s {
					fmt.Printf("%v", x)
				} else {
					fmt.Printf("-%v", x)
				}
			}
			fmt.Printf("\n")
		}
	}
}

// BFS 路径
func TestBFSPaths(t *testing.T) {
	tmpbuf := getTmpBuf()

	r := strings.NewReader(tmpbuf)

	g := NewGraphFromBufio(r)

	b := NewBreadthFirstPaths(g, 0)
	for v := 0; v < g.V; v++ {
		fmt.Printf("%v to %v : ", b.s, v)

		if b.HasPathTo(v) {
			paths := b.PathTo(v)
			for _, x := range paths {
				if x == b.s {
					fmt.Printf("%v", x)
				} else {
					fmt.Printf("-%v", x)
				}
			}
			fmt.Printf("\n")
		}
	}
}

// 连通性检查
func TestCC(t *testing.T) {
	f, err := os.Open("tinyG.txt")
	if err != nil {
		t.Fatal(err)
	}

	r := bufio.NewReader(f)

	g := NewGraphFromBufio(r)

	cc := NewCC(g)
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

// 两点 连通
func TestSimpleCycle(t *testing.T) {
	tmpbuf := `2
1
0 1
`
	r := strings.NewReader(tmpbuf)
	g := NewGraphFromBufio(r)
	c := NewCycle(g)

	fmt.Println("has cycle: ", c.HasCycle())
}

// 三点 非连通
func TestSimpleCycle2(t *testing.T) {
	tmpbuf := `3
2
0 1
1 2
`
	r := strings.NewReader(tmpbuf)
	g := NewGraphFromBufio(r)
	c := NewCycle(g)

	fmt.Println("has cycle: ", c.HasCycle())
}

// 三点,连通
func TestSimpleCycle3(t *testing.T) {
	tmpbuf := `3
3
0 1
1 2
2 0
`
	r := strings.NewReader(tmpbuf)
	g := NewGraphFromBufio(r)
	c := NewCycle(g)

	fmt.Println("has cycle: ", c.HasCycle())
}

// 四点,两两分
func TestSimpleCycle4(t *testing.T) {
	tmpbuf := `4
2
0 1
2 3
`
	r := strings.NewReader(tmpbuf)
	g := NewGraphFromBufio(r)
	c := NewCycle(g)

	fmt.Println("has cycle: ", c.HasCycle())
}

// 两点判读是否为二分图
func TestSimpleBipartite1(t *testing.T) {
	tmpbuf := `2
1
0 1
`
	r := strings.NewReader(tmpbuf)
	g := NewGraphFromBufio(r)
	c := NewTwoColor(g)

	fmt.Println(c.IsBipartite())
}

// 三点判读是否为二分图
func TestSimpleBipartite2(t *testing.T) {
	tmpbuf := `3
3
0 1
1 2
2 0
`
	r := strings.NewReader(tmpbuf)
	g := NewGraphFromBufio(r)
	c := NewTwoColor(g)

	fmt.Println(c.IsBipartite())
}

// 四点判读是否为二分图
func TestSimpleBipartite3(t *testing.T) {
	tmpbuf := `4
4
0 1
1 2
2 3
3 0
`
	r := strings.NewReader(tmpbuf)
	g := NewGraphFromBufio(r)
	c := NewTwoColor(g)

	fmt.Println(c.IsBipartite())
}

// 测试symbol转为图
func TestSymbolGraph(t *testing.T) {
	s := NewSymbolGraph("routes.txt", " ")

	if !s.Contains("MCO") {
		t.Fatal("err")
	}

	idx := s.Index("MCO")

	fmt.Println(s.Name(idx))
}

