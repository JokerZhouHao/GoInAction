package main

func JumpFloor(target int) int {
	if target == 1 {
		return 1
	} else if target == 2 {
		return 2
	} else {
		return JumpFloor(target-1) + JumpFloor(target-2)
	}
}
