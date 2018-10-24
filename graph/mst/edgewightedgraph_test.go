package mst

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	edge3 := &Edge{weight: 3.0}
	edge2 := &Edge{weight: 2.0}
	edge1 := &Edge{weight: 1.0}

	pq := make(EdgePriorityQueue, 0)

	heap.Init(&pq)

	heap.Push(&pq, edge1)
	heap.Push(&pq, edge2)
	heap.Push(&pq, edge3)

	for pq.Len() > 0 {
		edge := heap.Pop(&pq).(*Edge)
		fmt.Printf("weight: %v\n", edge.weight)
	}
}

func TestLazyPrimMST(t *testing.T) {
	f, err := os.Open("tinyEWG.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	g := NewGraphFromBufio(r)

	msg := NewLazyPrimMST(g)

	for _, e := range msg.Edges() {
		fmt.Println(e.String())
	}
}
