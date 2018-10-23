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
