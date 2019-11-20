//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func f(n int) int {
	if n < 1 {
		panic("n can't below 1")
	}
	return (1 + n) * n / 2
}

func sum(n int) int {
	n = n - 1
	return 3*f(n/3) + 5*f(n/5)
}

func main() {
	fmt.Println(sum(1000))
}
