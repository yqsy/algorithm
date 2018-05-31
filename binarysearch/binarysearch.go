package binarysearch

import "errors"

func binarySearch(a []int, t int) (int, error) {
	first, last := 0, len(a)-1

	for ; first <= last; {
		mid := first + (last-first)/2

		if t > a[mid] {
			first = mid + 1
		}

		if t < a[mid] {
			last = mid - 1
		}

		if t == a[mid] {
			return mid, nil
		}
	}

	return -1, errors.New("can not find in array")
}
