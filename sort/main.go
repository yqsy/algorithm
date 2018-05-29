package main

import (
	"os"
	"fmt"
	"strconv"
	"math/rand"
	"time"
	hahaha "github.com/yqsy/algorithm/sort/go"
	"sort"
	"reflect"
)

var usage = `Usage:
%v num
`

func createClone(array []int) []int {
	buf := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		buf[i] = array[i]
	}
	return buf
}

func test(hint string, foo func()) {
	t1 := time.Now()
	foo()
	escaped := float64(time.Since(t1).Nanoseconds()) / 1000 / 1000 // ms

	fmt.Printf("%v cost:%.3f ms\n", hint, escaped)
}

func main() {
	arg := os.Args

	usage = fmt.Sprintf(usage, arg[0])

	if len(arg) < 2 {
		fmt.Printf(usage)
		return
	}

	num, err := strconv.Atoi(arg[1])
	if err != nil {
		panic(err)
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	buf := make([]int, num)
	for i := 0; i < len(buf); i++ {
		buf[i] = rnd.Int()
	}

	fmt.Printf("===%v===\n", num)

	array := createClone(buf)
	test("default sort", func() { sort.Ints(array) })

	sorted := createClone(array)

	array = createClone(buf)
	test("heap sort", func() { hahaha.HeapSort(array) })
	if !reflect.DeepEqual(array, sorted) {
		panic("err")
	}

	array = createClone(buf)
	test("quick sort", func() { hahaha.QuickSort(array, 0, len(array)-1) })
	if !reflect.DeepEqual(array, sorted) {
		panic("err")
	}

	array = createClone(buf)
	test("quick sort recursion", func() { hahaha.QuickSortRecursion(array, 0, len(array)-1) })
	if !reflect.DeepEqual(array, sorted) {
		panic("err")
	}

	array = createClone(buf)
	test("shell sort", func() { hahaha.ShellSort(array) })
	if !reflect.DeepEqual(array, sorted) {
		panic("err")
	}

	//array = createClone(buf)
	//test("insert sort", func() { hahaha.InsertSort(array) })
	//if !reflect.DeepEqual(array, sorted) {
	//	panic("err")
	//}

	array = createClone(buf)
	tmp := make([]int, len(array))
	test("merge sort", func() { hahaha.MergeSort(array, tmp) })
	if !reflect.DeepEqual(array, sorted) {
		panic("err")
	}

	array = createClone(buf)
	test("merge sort recursion", func() { hahaha.MergeSortRecursion(array, 0, len(array)-1, tmp) })
	if !reflect.DeepEqual(array, sorted) {
		panic("err")
	}

	//array = createClone(buf)
	//test("bubble sort", func() { hahaha.BubbleSort(array) })
	//if !reflect.DeepEqual(array, sorted) {
	//	panic("err")
	//}
	//
	//array = createClone(buf)
	//test("selection sort", func() { hahaha.SelectionSort(array) })
	//if !reflect.DeepEqual(array, sorted) {
	//	panic("err")
	//}

	fmt.Printf("===%v===\n", "end")
}
