package main

/*
在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/

func find(target int, array [][]int) bool {
	if array == nil || array[0] == nil {
		return false
	}
	xlen := len(array)
	ylen := len(array[0])
	i, j := 0, ylen-1
	for i < xlen && j >= 0 {
		if array[i][j] > target {
			j--
		} else if array[i][j] < target {
			i++
		} else {
			return true
		}
	}
	return false
}
