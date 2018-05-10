package repeat_exchange

import "errors"

func IsArrayRepeat(array []int) (bool, error) {
	if len(array) < 2 {
		return false, nil
	}

	for i := 0; i < len(array); i ++ {
		if array[i] < 0 || array[i] > len(array)-1 {
			return false, errors.New("array in not legal")
		}
	}

	for i := 0; i < len(array); i++ {
		for array[i] != i {

			// 将当前下标内元素换到到它应该在的位置
			k := array[i]
			if k == array[k] {
				return true, nil
			}

			array[i], array[k] = array[k], array[i]
		}
	}

	return false, nil
}
