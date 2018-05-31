package rotatemin

// 借O(n)的辅助空间腾挪
func rotate(nums []int, k int) {
	tmp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp[i] = nums[i]
	}

	for i := 0; i < len(tmp); i ++ {
		nums[(i+k)%len(tmp)] = tmp[i]
	}
}

// 借助O(1)的辅助空间腾挪 (前面部分借助空间)
func rotate2(nums []int, k int) {
	k %= len(nums)

	s := len(nums) - k
	tmp := make([]int, s)
	for i := 0; i < s; i++ {
		tmp[i] = nums[i]
	}

	n := 0
	for i := s; i < len(nums); i++ {
		nums[n] = nums[i]
		n++
	}

	for i := 0; i < len(tmp); i++ {
		nums[n] = tmp[i]
		n++
	}
}

// 借助O(1)的辅助空间腾挪 (后面部分借助空间) (leetcode最快速解)
func rotate3(nums []int, k int) {
	k %= len(nums)

	s := len(nums) - k
	tmp := make([]int, k)

	n := 0
	for i := s; i < len(nums); i++ {
		tmp[n] = nums[i]
		n++
	}

	for i := s - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}

	for i := 0; i < len(tmp); i++ {
		nums[i] = tmp[i]
	}
}

func findMinSimple(a []int) int {
	if len(a) < 1 {
		return 0
	}

	minest := 0

	for i := 0; i < len(a); i++ {
		if a[i] < a[minest] {
			minest = i
		}
	}

	return a[minest]
}

// 前面部分找到最大的数字,再加1
func findMin(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}

	// 中分找到最大数字,最大数字下标加一
	begin, end := 0, len(a)

	maxIdx := -1
	for {
		pivot := begin + (end-begin)/2

		if begin == pivot {
			maxIdx = begin
			break
		}

		if a[pivot] < a[begin] {
			end = pivot
		}

		if a[pivot] > a[begin] {
			begin = pivot
		}
	}

	if maxIdx >= len(a)-1 {
		return findMinSimple(a)
	} else {
		return a[maxIdx+1]
	}
}

// 直接找后半部分最小数字
func findMin2(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}

	first, last := 0, len(a)-1
	for ; first < last; {
		mid := first + (last-first)/2

		if a[mid] < a[last] {
			last = mid
		} else if a[mid] > a[last] {
			first = mid + 1
		} else {
			last--
		}
	}

	return a[first]
}
