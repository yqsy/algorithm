package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	test := []int{3, 5, 10, 0, 1, 2, 20, 30}

	for _, n := range test {
		if fibonacciRecursion(n) != fibonacci(n) {
			t.Fatal("err")
		}
	}
}
