package graph

import "testing"

func TestNodeQueueSimple(t *testing.T) {

	nq := NewNodeQueue()

	nq.Enqueue(Node{1})
	nq.Enqueue(Node{2})
	nq.Enqueue(Node{3})

	if nq.Size() != 3 || nq.IsEmpty() {
		t.Fatal("err")
	}

	if nq.Front().value != 1 {
		t.Fatal("err")
	}

	if nq.Dequeue().value != 1 || nq.Dequeue().value != 2 || nq.Dequeue().value != 3{
		t.Fatal("err")
	}

	if nq.Size() != 0 || !nq.IsEmpty() {
		t.Fatal("err")
	}
}
