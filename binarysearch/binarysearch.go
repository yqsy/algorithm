package binarysearch

import "errors"

func binarySearch(a []int, t int) (int, error) {
	low, high := 0, len(a)-1

	for low <= high {
		mid := low + (high-low)/2

		if t > a[mid] {
			low = mid + 1
		}

		if t < a[mid] {
			high = mid - 1
		}

		if t == a[mid] {
			return mid, nil
		}
	}

	return -1, errors.New("can not find in array")
}
