package main

import "fmt"

func reOrderArray(array []int) {
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 1 {
			for j := i - 1; j >= 0; j-- {
				if array[j]%2 == 0 {
					t := array[j]
					array[j] = array[j+1]
					array[j+1] = t
				} else {
					break
				}
			}
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 31, 89, 6}
	reOrderArray(arr)
	fmt.Println(arr)
}
