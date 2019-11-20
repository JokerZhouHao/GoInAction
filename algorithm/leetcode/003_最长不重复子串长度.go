package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	seqLen := 0
	len := len(s)
	nums := make([]int, len)
	for i := 0; i < len; i++ {
		nums[i] = 1
		j := i - 1
		for ; j >= 0; j-- {
			if s[i] != s[j] && nums[j] != -1 {
				nums[j]++
			} else {
				nums[j] = -1
				break
			}
		}
		j++
		if seqLen < nums[j] {
			seqLen = nums[j]
		}
	}
	return seqLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
