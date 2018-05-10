package repeat_sort

import "sort"

func IsArrayRepeat(array []int) bool {
	sort.Ints(array)

	if len(array) < 2 {
		return false
	}

	for i := 0; i < len(array)-1; i++ {
		if array[i] == array[i+1] {
			return true
		}
	}

	return false
}
