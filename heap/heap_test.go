package heap

import (
	"testing"
	"reflect"
)

func TestHeapSimle(t *testing.T) {
	heap := NewHeap()

	if heap.N() != 0 {
		t.Fatalf("err %v", heap.N())
	}

	heap.append(1)
	heap.append(2)
	heap.append(3)

	if heap.N() != 3 {
		t.Fatal("err")
	}

	heap.delMax()
	heap.delMax()

	if heap.N() != 1 {
		t.Fatal("err")
	}

	if !reflect.DeepEqual(heap.snap(), []int{1}) {
		t.Fatal("err")
	}
}
