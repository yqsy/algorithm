package binarysearch

import "testing"

func TestSimple(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if idx, _ := binarySearch(a, 5); idx != 4 {
		t.Fatal("err")
	}

	a = []int{1}

	if idx, _ := binarySearch(a, 1); idx != 0 {
		t.Fatal("err")
	}

	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if idx, err := binarySearch(a, 10); idx != -1 || err == nil {
		t.Fatal("err")
	}
}
