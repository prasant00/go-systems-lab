package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	minLen := n + 1
	sum := 0
	left := 0

	for right := 0; right < n; right++ {
		sum += nums[right]

		// shrink window from left while sum is enough
		for sum >= target {
			currLen := right - left + 1
			if currLen < minLen {
				minLen = currLen
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == n+1 {
		return 0
	}
	return minLen
}

func main() {
	target := 11
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(minSubArrayLen(target, nums)) // Output: 2
}
