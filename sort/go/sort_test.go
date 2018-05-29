package sort

import (
	"testing"
	"reflect"
)

func createArray1() []int {
	return []int{5, 7, 6, 4, 8, 9, 2, 1, 3}
}

func checkArray1(array []int) bool {
	return reflect.DeepEqual(array, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func createArray2() []int {
	return []int{3, 2}
}

func checkArray2(array []int) bool {
	return reflect.DeepEqual(array, []int{2, 3})
}

func createArray3() []int {
	return []int{3}
}

func checkArray3(array []int) bool {
	return reflect.DeepEqual(array, []int{3})
}

func TestInsertSort(t *testing.T) {

	array1 := createArray1()
	InsertSort(array1)
	if !checkArray1(array1) {
		t.Fatal("err")
	}

	array2 := createArray2()
	InsertSort(array2)
	if !checkArray2(array2) {
		t.Fatal("err")
	}

	array3 := createArray3()
	InsertSort(array3)
	if !checkArray3(array3) {
		t.Fatal("err")
	}
}

func TestQuickSortRecursion(t *testing.T) {
	array1 := createArray1()
	QuickSortRecursion(array1, 0, len(array1)-1)
	if !checkArray1(array1) {
		t.Fatalf("err %v", array1)
	}

	array2 := createArray2()
	QuickSortRecursion(array2, 0, len(array2)-1)
	if !checkArray2(array2) {
		t.Fatal("err")
	}

	array3 := createArray3()
	QuickSortRecursion(array3, 0, len(array3)-1)
	if !checkArray3(array3) {
		t.Fatal("err")
	}
}

func TestQuickSort(t *testing.T) {

	array1 := createArray1()
	QuickSort(array1, 0, len(array1)-1)
	if !checkArray1(array1) {
		t.Fatalf("err %v", array1)
	}

	array2 := createArray2()
	QuickSort(array2, 0, len(array2)-1)
	if !checkArray2(array2) {
		t.Fatal("err")
	}

	array3 := createArray3()
	QuickSort(array3, 0, len(array3)-1)
	if !checkArray3(array3) {
		t.Fatal("err")
	}
}

func TestMergeSortRecursion(t *testing.T) {
	array1 := createArray1()
	tmp := make([]int, len(array1))
	MergeSortRecursion(array1, 0, len(array1)-1, tmp)
	if !checkArray1(array1) {
		t.Fatalf("err %v", array1)
	}

	array2 := createArray2()
	tmp = make([]int, len(array2))
	MergeSortRecursion(array2, 0, len(array2)-1, tmp)
	if !checkArray2(array2) {
		t.Fatal("err")
	}

	array3 := createArray3()
	tmp = make([]int, len(array3))
	MergeSortRecursion(array3, 0, len(array3)-1, tmp)
	if !checkArray3(array3) {
		t.Fatal("err")
	}
}

func TestMergeSort(t *testing.T) {
	array1 := createArray1()
	tmp := make([]int, len(array1))
	MergeSort(array1, tmp)
	if !checkArray1(array1) {
		t.Fatalf("err %v", array1)
	}

	array2 := createArray2()
	tmp = make([]int, len(array2))
	MergeSort(array2, tmp)
	if !checkArray2(array2) {
		t.Fatal("err")
	}

	array3 := createArray3()
	tmp = make([]int, len(array3))
	MergeSort(array3, tmp)
	if !checkArray3(array3) {
		t.Fatal("err")
	}
}

func TestHeapSort(t *testing.T) {
	array1 := createArray1()
	HeapSort(array1)
	if !checkArray1(array1) {
		t.Fatalf("err %v", array1)
	}

	array2 := createArray2()
	HeapSort(array2)
	if !checkArray2(array2) {
		t.Fatal("err")
	}

	array3 := createArray3()
	HeapSort(array3)
	if !checkArray3(array3) {
		t.Fatal("err")
	}
}

func TestShellSort(t *testing.T) {
	array1 := createArray1()
	ShellSort(array1)
	if !checkArray1(array1) {
		t.Fatalf("err %v", array1)
	}

	array2 := createArray2()
	ShellSort(array2)
	if !checkArray2(array2) {
		t.Fatal("err")
	}

	array3 := createArray3()
	ShellSort(array3)
	if !checkArray3(array3) {
		t.Fatal("err")
	}
}

func TestBubbleSort(t *testing.T) {
	array1 := createArray1()
	BubbleSort(array1)
	if !checkArray1(array1) {
		t.Fatalf("err %v", array1)
	}

	array2 := createArray2()
	BubbleSort(array2)
	if !checkArray2(array2) {
		t.Fatal("err")
	}

	array3 := createArray3()
	BubbleSort(array3)
	if !checkArray3(array3) {
		t.Fatal("err")
	}
}

func TestSelectionSort(t *testing.T) {
	array1 := createArray1()
	SelectionSort(array1)
	if !checkArray1(array1) {
		t.Fatalf("err %v", array1)
	}

	array2 := createArray2()
	SelectionSort(array2)
	if !checkArray2(array2) {
		t.Fatal("err")
	}

	array3 := createArray3()
	SelectionSort(array3)
	if !checkArray3(array3) {
		t.Fatal("err")
	}
}
