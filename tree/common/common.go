package common


func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}


func AddSpaces(n int, s *string) {
	for i := 0; i < n; i++ {
		*s += " "
	}
}


