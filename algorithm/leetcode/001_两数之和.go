package main

import "fmt"

func twoSum1(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		j, ok := mp[target-v]
		if ok {
			return []int{j, i}
		}
		mp[v] = i
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		s := nums[i] + nums[j]
		if s > target {
			j--
		} else if s < target {
			i++
		} else {
			return []int{i, j}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum3(nums, 13))

}
