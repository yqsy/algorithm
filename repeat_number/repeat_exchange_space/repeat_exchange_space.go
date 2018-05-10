package repeat_exchange_space

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

	buf := make([]int, len(array))
	copy(buf, array)

	for i := 0; i < len(buf); i++ {
		for buf[i] != i {

			// 将当前下标内元素换到到它应该在的位置
			k := buf[i]
			if k == buf[k] {
				return true, nil
			}

			buf[i], buf[k] = buf[k], buf[i]
		}
	}

	return false, nil
}
