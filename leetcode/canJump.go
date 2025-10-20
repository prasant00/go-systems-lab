package main

import "fmt"

func canJump(nums []int) bool {
	n := len(nums)
	reach := 0

	for i := 0; i < n; i++ {
		if i > reach {
			return false
		}
		if i+nums[i] > reach {
			reach = i + nums[i]
		}
		if reach >= n-1 {
			return true
		}
	}
	return true
}

func main() {
	fmt.Printf("canJump: %t\n", canJump([]int{3, 1, 2, 1, 0, 5}))
}
