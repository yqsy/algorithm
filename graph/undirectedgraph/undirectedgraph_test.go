package undirectedgraph

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestSimpleString(t *testing.T) {
	f, err := os.Open("tinyG.txt")
	if err != nil {
		t.Fatal(err)
	}

	r := bufio.NewReader(f)

	g := NewGraphFromBufio(r)

	fmt.Println(g.String())
}

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

func TestDFS(t *testing.T) {
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
	r := strings.NewReader(tmpbuf)

	g := NewGraphFromBufio(r)

	d := NewDepthFirstSearch(g, 0)

	if d.Count() != 6 {
		t.Fatal("err")
	}
}

func TestDFSPaths(t *testing.T) {
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
	r := strings.NewReader(tmpbuf)

	g := NewGraphFromBufio(r)

	d := NewDepthFirstPaths(g, 0)
	paths := d.PathTo(5)

	fmt.Println(paths)
}
