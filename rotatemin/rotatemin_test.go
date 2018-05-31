package rotatemin

import (
	"testing"
	"reflect"
)

func TestRotate1(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(array, 3)
	if !reflect.DeepEqual(array, []int{5, 6, 7, 1, 2, 3, 4}) {
		t.Fatalf("err %v", array)
	}

	array = []int{1, 2}
	rotate(array, 0)
	if !reflect.DeepEqual(array, []int{1, 2}) {
		t.Fatalf("err %v", array)
	}

	array = []int{1}
	rotate(array, 1)
}

func TestRotate2(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(array, 3)
	if !reflect.DeepEqual(array, []int{5, 6, 7, 1, 2, 3, 4}) {
		t.Fatalf("err %v", array)
	}

	array = []int{1, 2}
	rotate2(array, 0)
	if !reflect.DeepEqual(array, []int{1, 2}) {
		t.Fatalf("err %v", array)
	}

	array = []int{1}
	rotate2(array, 1)
}

func TestRotate3(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	rotate3(array, 3)
	if !reflect.DeepEqual(array, []int{5, 6, 7, 1, 2, 3, 4}) {
		t.Fatalf("err %v", array)
	}

	array = []int{1, 2}
	rotate3(array, 0)
	if !reflect.DeepEqual(array, []int{1, 2}) {
		t.Fatalf("err %v", array)
	}

	array = []int{1}
	rotate3(array, 1)
}

func TestFindMin(t *testing.T) {
	array := []int{5, 6, 7, 1, 2, 3, 4}

	if findMin2(array) != 1 {
		t.Fatal("err")
	}
}
