package repeat_hash

func IsArrayRepeat(array []int) bool {
	if len(array) < 2 {
		return false
	}

	d := make(map[int]struct{})

	for i := 0; i < len(array); i++ {
		if _, ok := d[array[i]]; ok {
			return true
		}
		d[array[i]] = struct{}{}
	}

	return false
}
