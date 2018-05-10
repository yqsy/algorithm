package repeat_binary

import (
	"errors"
)

func getCount(array []int, begin, end int) int {
	var n int

	for _, val := range array {
		if val >= begin && val <= end {
			n += 1
		}
	}
	return n
}

func IsArrayRepeat(array []int) (bool, error) {
	if len(array) < 2 {
		return false, nil
	}

	for i := 0; i < len(array); i ++ {
		if array[i] < 0 || array[i] > len(array)-1 {
			return false, errors.New("array in not legal")
		}
	}

	var low, high = 0, len(array)-1

	for low <= high {
		mid := low + (high-low)/2

		n := getCount(array, low, mid)

		// 这里不找到重复的数字
		// 直接跳出
		if n > mid-low+1 {
			high = mid - 1
			return true, nil
		} else {
			low = mid + 1
		}
	}

	return false, nil
}
