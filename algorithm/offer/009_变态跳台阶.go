package main

func jumpFloorII(target int, nums []int) int {
	num := 0
	if target == 0 {
		num = 1
	} else {
		for i := 0; i < target; i++ {
			if nums[i] != 0 {
				num += nums[i]
			} else {
				num += jumpFloorII(i, nums)
			}
		}
	}
	nums[target] = num
	return num
}

func JumpFloorII(target int) int {
	return jumpFloorII(target, make([]int, target+1))
}
