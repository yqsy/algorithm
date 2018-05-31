package sort

import (
	"github.com/golang-collections/collections/stack"
)

func InsertSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j >= 1 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
}

func partition(array []int, first, last int) int {
	cur := array[first]
	p := first + 1
	for i := first + 1; i <= last; i++ {
		if array[i] < cur {
			array[p], array[i] = array[i], array[p]
			p += 1
		}
	}
	array[p-1], array[first] = array[first], array[p-1]
	return p - 1
}

func QuickSort(array []int, first, last int) {
	if first < last {
		s := stack.New()

		mid := partition(array, first, last)

		if mid+1 < last {
			s.Push([]int{mid + 1, last})
		}

		if first < mid-1 {
			s.Push([]int{first, mid - 1})
		}

		for ; s.Len() > 0; {

			ele := s.Pop().([]int)
			l, r := ele[0], ele[1]

			mid := partition(array, l, r)

			if mid+1 < r {
				s.Push([]int{mid + 1, r})
			}

			if l < mid-1 {
				s.Push([]int{l, mid - 1})
			}

		}
	}
}

func QuickSortRecursion(array []int, first, last int) {
	if first < last {
		mid := partition(array, first, last)
		QuickSortRecursion(array, first, mid-1)
		QuickSortRecursion(array, mid+1, last)
	}
}

// mid 是前半个空间的最后一个元素
func merge(array []int, first, mid, last int, tmp []int) {
	j := first
	k := mid + 1

	n := 0
	for ; j <= mid && k <= last; {
		if array[j] < array[k] {
			tmp[n] = array[j]
			n++
			j++
		} else {
			tmp[n] = array[k]
			n++
			k++
		}
	}

	for ; j <= mid; {
		tmp[n] = array[j]
		n++
		j++
	}

	for ; k <= last; {
		tmp[n] = array[k]
		n++
		k++
	}

	for i := 0; i <= (last - first); i++ {
		array[first+i] = tmp[i]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func MergeSort(array []int, tmp []int) {
	for n := 1; n < len(array); n *= 2 {
		for i := 0; i+n < len(array); i += 2 * n {
			merge(array, i, i+n-1, min(i+2*n-1, len(array)-1), tmp)
		}
	}
}

func MergeSortRecursion(array []int, first, last int, tmp []int) {
	if first < last {
		mid := first + (last-first)/2
		MergeSortRecursion(array, first, mid, tmp)
		MergeSortRecursion(array, mid+1, last, tmp)
		merge(array, first, mid, last, tmp)
	}
}

// * 使用index[0]
// * 位置k的结点的父节点位置为 `(k-1)/2`
// * 两个子结点的位置分别为 `2*k+1` 和 `2*k+2`

func sink(array []int, k, last int) {
	for ; 2*k+1 <= last; {
		j := 2*k + 1
		if 2*k+1 < last && array[2*k+1] < array[2*k+2] {
			j++
		}

		if !(array[k] < array[j]) {
			break
		}

		array[j], array[k] = array[k], array[j]
		k = j
	}
}

func HeapSort(array []int) {
	// 构建堆
	last := len(array) - 1
	for k := (last - 1) / 2; k >= 0; k-- {
		sink(array, k, len(array)-1)
	}

	// 最大数字往末尾放
	for j := last; j > 0; {
		array[0], array[j] = array[j], array[0]
		j--
		sink(array, 0, j)
	}
}

func ShellSort(array []int) {
	for gap := len(array) / 2; gap > 0; gap /= 2 {
		// gap是跳跃的距离
		// 从第一个能跳gap的点,遍历到终点的每个点都跳跃
		for i := gap; i < len(array); i++ {
			for j := i; j >= gap && array[j] < array[j-gap]; j -= gap {
				array[j], array[j-gap] = array[j-gap], array[j]
			}
		}
	}
}

func BubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func SelectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		minIdx := i
		for j := i; j < len(array); j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		array[i], array[minIdx] = array[minIdx], array[i]
	}
}
