package list

type Node struct {
	data int
	next *Node
}

func NewNode(data int) *Node {
	node := &Node{data: data, next: nil}
	return node
}

type List struct {
	node *Node
	len  int
}

func (list *List) length() int {
	return list.len
}

func (list *List) append(data int) {
	if list.length() == 0 {
		list.node = NewNode(data)
		list.len += 1
	} else {
		it := list.node
		for ; it.next != nil; it = it.next {
		}
		it.next = NewNode(data)
		list.len += 1
	}
}

func (list *List) delete(data int) {
	if list.length() == 0 {
		return
	} else {
		// 打桩
		pile := NewNode(0)
		pile.next = list.node

		it := pile
		for ; it.next != nil; {
			if it.next.data == data {
				it.next = it.next.next
				list.len -= 1
			} else {
				it = it.next
			}
		}
		list.node = pile.next
	}
}

func (list *List) snap() []int {
	var snap []int

	if list.len < 1 {
		return snap
	}

	it := list.node
	for ; it != nil; it = it.next {
		snap = append(snap, it.data)
	}

	return snap
}

func (list *List) reverseSnap() []int {
	var snap []int

	if list.len < 1 {
		return snap
	}

	it := list.node
	for ; it != nil; it = it.next {
		snap = append(snap, it.data)
	}

	// reverse
	half := len(snap) / 2

	for i := 0; i < half; i++ {
		snap[i], snap[len(snap)-1-i] = snap[len(snap)-1-i], snap[i]
	}

	return snap
}
