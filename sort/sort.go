package sort

import (
	"github.com/golang-collections/collections/queue"
)

func insertSort(array []int) {
	if len(array) < 2 {
		return
	}

	for i := 1; i < len(array); i ++ {
		for j := i; j >= 1; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
}

func partition(array []int, begin, end int) int {
	cur := array[begin]

	leftIdx := begin + 1

	for i := begin + 1; i < end; i++ {
		if array[i] < cur {
			array[leftIdx], array[i] = array[i], array[leftIdx]
			leftIdx += 1
		}
	}
	array[leftIdx-1], array[begin] = array[begin], array[leftIdx-1]

	return leftIdx - 1
}

func quickSort(array []int, begin, end int) {
	q := queue.New()

	if begin < end {

		mid := partition(array, begin, end)

		if begin < mid {
			q.Enqueue([]int{begin, mid})
		}

		if mid+1 < end {
			q.Enqueue([]int{mid + 1, end})
		}

		for {
			if q.Len() < 1 {
				break
			}

			ele := q.Dequeue().([]int)
			begin, end = ele[0], ele[1]
			mid = partition(array, begin, end)

			if begin < mid {
				q.Enqueue([]int{begin, mid})
			}

			if mid+1 < end {
				q.Enqueue([]int{mid + 1, end})
			}
		}
	}
}

func quickSortRecursion(array []int, begin, end int) {
	if begin < end {
		mid := partition(array, begin, end)
		quickSortRecursion(array, begin, mid)
		quickSortRecursion(array, mid+1, end)
	}
}

func merge(array []int, begin, mid, end int, tmp []int) {
	j := begin
	k := mid + 1

	// tmp
	n := 0

	for {
		if !(j <= mid && k < end) {
			break
		}

		if array[j] > array[k] {
			tmp[n] = array[k]
			k++
			n++
		} else {
			tmp[n] = array[j]
			j++
			n++
		}
	}

	for {
		if j <= mid {
			tmp[n] = array[j]
			j++
			n++
		} else {
			break
		}
	}

	for {
		if k < end {
			tmp[n] = array[k]
			k++
			n++
		} else {
			break
		}
	}

	for i := begin; i < end; i++ {
		array[i] = tmp[i]
	}
}

func mergeSortRecursion(array []int, begin, end int, tmp []int) {
	if begin < end {
		mid := (begin + end) / 2
		mergeSortRecursion(array, begin, mid+1, tmp)
		mergeSortRecursion(array, mid+1, end, tmp)
		merge(array, begin, mid, end, tmp)
	}
}
