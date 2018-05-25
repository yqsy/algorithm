package stackqueue

import (
	"testing"
)

func TestQueueSimple(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Dequeue() != 1 || q.Dequeue() != 2 || q.Dequeue() != 3 {
		t.Fatal("err")
	}

	if q.Dequeue() != nil {
		t.Fatal("err")
	}
}

func TestQueueSimple2(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Dequeue() != 1 {
		t.Fatal("err")
	}

	q.Enqueue(4)

	if q.Dequeue() != 2 || q.Dequeue() != 3 || q.Dequeue() != 4 {
		t.Fatal("err")
	}

	if q.Dequeue() != nil {
		t.Fatal("err")
	}
}

func TestStackSimple(t *testing.T) {
	s := NewStack()

	s.Push(1)

	if s.Pop() != 1 {
		t.Fatal("err")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	if s.Pop() != 4 {
		t.Fatal("err")
	}

	if s.Pop() != 3 {
		t.Fatal("err")
	}

	if s.Pop() != 2 {
		t.Fatal("err")
	}

	if s.Pop() != 1 {
		t.Fatal("err")
	}
}
