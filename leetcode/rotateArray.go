package main

import "fmt"

// 1,2,3,4,5,6,7
// 7,6,5,4,3,2,1
// 5,6,7,4,3,2,1
// 5,6,7,1,2,3,4

func reverse(nums []int, left, right int) {

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	k = k % len(nums)
	if k == 0 {
		return
	}

	reverse(nums, 0, len(nums)-1)
	reverse(nums, k, len(nums)-1)
	reverse(nums, 0, k-1)
}

func main() {
	tests := [][]int{
		//{1, 2, 3, 4, 5, 6, 7},
		{1, 2},
	}
	for _, t := range tests {
		orig := make([]int, len(t))
		copy(orig, t)
		rotate(t, 7)
		fmt.Printf("\t rotated array = %v", t)
	}
}
