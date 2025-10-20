package main

import "fmt"

func removeDuplicatesAll(nums []int) int {
	j := 0
	for _, val := range nums {
		if j < 1 || nums[j-1] != val {
			nums[j] = val
			j++
		}
	}
	return j
}

func main() {
	tests := [][]int{
		{7, 7, 7, 6, 6, 5, 4, 4, 2},
		{1, 1, 2, 2, 2, 4, 5, 5, 5},
	}
	for _, t := range tests {
		orig := make([]int, len(t))
		copy(orig, t)
		idx := removeDuplicatesAll(t)
		fmt.Printf("\t orig %v, finalLength %d, finalArray %v\n", orig, idx, t[:idx])
	}
}
