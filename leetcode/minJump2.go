package main

import "fmt"

func jump2(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps := 0
	end := 0
	farthest := 0

	for i := 0; i < n-1; i++ {
		// update farthest reachable from positions up to i
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		// when we reach the end of the current range, make a jump
		if i == end {
			jumps++
			end = farthest
			// early exit if we can already reach or pass last index
			if end >= n-1 {
				break
			}
		}
	}
	return jumps
}

func main() {
	nums := []int{1, 2, 5, 2, 3, 1, 4, 6}
	fmt.Printf("\t min Jumps %d\n", jump2(nums))
}
