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
