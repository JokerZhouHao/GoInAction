package main

func fibonacci1(n int, xs []int) int {
	x := 0
	if n == 0 {
		x = 0
	} else if n == 1 {
		x = 1
	} else {
		if xs[n-1] == 0 {
			x += fibonacci1(n-1, xs)
		} else {
			x += xs[n-1]
		}
		if xs[n-2] == 0 {
			x += fibonacci1(n-2, xs)
		} else {
			x += xs[n-2]
		}
	}
	xs[n] = x
	return x
}

func Fibonacci(n int) int {
	xs := make([]int, n+1)
	return fibonacci1(n, xs)
}
