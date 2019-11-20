package main

import "fmt"

func minNumberInRotateArray(array []int) int {
	if array == nil {
		return 0
	}

	l := len(array)
	i1, i2 := 0, 1
	for i2 < l {
		if array[i1] > array[i2] {
			return array[i2]
		}
		i1++
		i2++
	}
	return array[0]
}

func main() {
	fmt.Println(minNumberInRotateArray([]int{1, 2, 5, 3}))
}
