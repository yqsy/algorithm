package fibonacci

func fibonacciRecursion(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	a1, a2 := 1, 1

	for i := 0; i < n-2; i++ {
		a1, a2 = a2, a1+a2
	}

	return a2
}
