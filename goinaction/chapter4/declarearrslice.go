package main

import "fmt"

func main() {
	/* ------------- 数组 ---------------- */
	// 声明一个包含5个元素的整型数组
	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [5]int{1: 10, 4: 40}
	// 多维数组
	arr4 := [4][2]int{1: {20, 21}, 3: {1: 41}}
	fmt.Printf("%v\n%v\n%v\n%v\n", arr1, arr2, arr3, arr4)

	/* ------------- 切片 ---------------- */
	slice1 := make([]string, 5)
	slice2 := make([]string, 5, 10)
	slice3 := []int{10, 20, 30}
	// 声明空切片
	var slice4 []int
	slice5 := make([]int, 0)
	slice6 := []int{}
	// 多维切片
	slice7 := [][]int{{10}, {100, 200}}
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n", slice1, slice2, slice3, slice4, slice5, slice6, slice7)
}
