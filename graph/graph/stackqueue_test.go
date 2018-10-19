package graph

import "testing"

func TestNodeStackSimple(t *testing.T) {

	ns := NewNodeStack()

	ns.Push(Node{1})
	ns.Push(Node{2})
	ns.Push(Node{3})

	if ns.Size() != 3 || ns.IsEmpty() {
		t.Fatal("err")
	}

	if ns.Front().value != 3 {
		t.Fatal("err")
	}

	if ns.Pop().value != 3 || ns.Pop().value != 2 || ns.Pop().value != 1 {
		t.Fatal("err")
	}

	if ns.Size() != 0 || !ns.IsEmpty() {
		t.Fatal("err")
	}

}
