package sort

import "github.com/golang-collections/collections/queue"

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

func quickSort(array []int, first, last int) {
	if first < last {
		q := queue.New()

		mid := partition(array, first, last)

		if first < mid-1 {
			q.Enqueue([]int{first, mid - 1})
		}

		if mid+1 < last {
			q.Enqueue([]int{mid + 1, last})
		}

		for ;q.Len()>0;{

			ele := q.Dequeue().([]int)
			l, r := ele[0], ele[1]

			mid := partition(array, l, r)

			if l < mid-1 {
				q.Enqueue([]int{l, mid - 1})
			}

			if mid+1 < r {
				q.Enqueue([]int{mid + 1, r})
			}
		}
	}
}

func quickSortRecursion(array []int, first, last int) {
	if first < last {
		mid := partition(array, first, last)
		quickSortRecursion(array, first, mid-1)
		quickSortRecursion(array, mid+1, last)
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

func mergeSortRecursion(array []int, first, last int, tmp []int) {
	if first < last {
		mid := (first + last) / 2
		mergeSortRecursion(array, first, mid, tmp)
		mergeSortRecursion(array, mid+1, last, tmp)
		merge(array, first, mid, last, tmp)
	}
}
