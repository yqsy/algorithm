package list

import (
	"testing"
	"reflect"
)

func TestListSimple(t *testing.T) {

	var lst List

	if lst.length() != 0 {
		t.Fatal("len err")
	}

	lst.append(1)

	if lst.length() != 1 {
		t.Fatal("len err")
	}

	if !reflect.DeepEqual(lst.snap(), []int{1}) {
		t.Fatal("not equal")
	}

	if !reflect.DeepEqual(lst.reverseSnap(), []int{1}) {
		t.Fatal("not equal")
	}

	lst.append(1)
	lst.append(2)

	if lst.length() != 3 {
		t.Fatal("len err")
	}

	if !reflect.DeepEqual(lst.snap(), []int{1, 1, 2}) {
		t.Fatal("not equal")
	}

	if !reflect.DeepEqual(lst.reverseSnap(), []int{2, 1, 1}) {
		t.Fatal("not equal")
	}

	lst.delete(1)

	if lst.length() != 1 {
		t.Fatal("len err")
	}

	if !reflect.DeepEqual(lst.snap(), []int{2}) {
		t.Fatal("not equal")
	}

	if !reflect.DeepEqual(lst.reverseSnap(), []int{2}) {
		t.Fatal("not equal")
	}

	lst.delete(2)

	if lst.length() != 0 {
		t.Fatal("len err")
	}

	lst.append(1)
	lst.append(3)
	lst.append(0)
	lst.append(2)
	lst.append(3)
	lst.append(2)
	lst.append(5)
	lst.append(2)
	lst.append(6)
	lst.append(1)
	lst.append(7)

	if !reflect.DeepEqual(lst.reverseSnap(), []int{7, 1, 6, 2, 5, 2, 3, 2, 0, 3, 1}) {
		t.Fatal("not equal")
	}



}
