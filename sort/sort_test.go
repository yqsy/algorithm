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
	insertSort(array1)
	if !checkArray1(array1) {
		t.Fatal("err")
	}

	array2 := createArray2()
	insertSort(array2)
	if !checkArray2(array2) {
		t.Fatal("err")
	}

	array3 := createArray3()
	insertSort(array3)
	if !checkArray3(array3) {
		t.Fatal("err")
	}
}

func TestQuickSortRecursion(t *testing.T) {

	array1 := createArray1()
	quickSortRecursion(array1, 0, len(array1))
	if !checkArray1(array1) {
		t.Fatalf("err %v", array1)
	}

	array2 := createArray2()
	quickSortRecursion(array2, 0, len(array2))
	if !checkArray2(array2) {
		t.Fatal("err")
	}

	array3 := createArray3()
	quickSortRecursion(array3, 0, len(array3))
	if !checkArray3(array3) {
		t.Fatal("err")
	}
}

func TestQuickSort(t *testing.T) {

	array1 := createArray1()
	quickSort(array1, 0, len(array1))
	if !checkArray1(array1) {
		t.Fatalf("err %v", array1)
	}

	array2 := createArray2()
	quickSort(array2, 0, len(array2))
	if !checkArray2(array2) {
		t.Fatal("err")
	}

	array3 := createArray3()
	quickSort(array3, 0, len(array3))
	if !checkArray3(array3) {
		t.Fatal("err")
	}
}

func TestMergeSort(t *testing.T) {

	array1 := createArray1()
	tmp := make([]int, len(array1))
	mergeSortRecursion(array1, 0, len(array1), tmp)
	if !checkArray1(array1) {
		t.Fatalf("err %v", array1)
	}

	//array2 := createArray2()
	//tmp = make([]int, len(array2))
	//mergeSortRecursion(array2, 0, len(array2), tmp)
	//if !checkArray2(array2) {
	//	t.Fatal("err")
	//}
	//
	//array3 := createArray3()
	//tmp = make([]int, len(array3))
	//mergeSortRecursion(array3, 0, len(array3), tmp)
	//if !checkArray3(array3) {
	//	t.Fatal("err")
	//}
}
