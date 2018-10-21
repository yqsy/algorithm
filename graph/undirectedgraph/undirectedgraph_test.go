package undirectedgraph

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestSimple(t *testing.T) {
	f, err := os.Open("tinyG.txt")
	if err != nil {
		t.Fatal(err)
	}

	r := bufio.NewReader(f)

	g := NewGraphFromBufio(r)

	fmt.Println(g.String())
}
