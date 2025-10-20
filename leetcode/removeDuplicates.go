package main

import (
	"fmt"
)

// removeDuplicates allows each element at most twice and returns the new length.
// It modifies nums in-place; values beyond the returned length are irrelevant.
func removeDuplicates(nums []int) int {
	i := 2
	for _, val := range nums {
		if i < 2 || nums[i-2] != val {
			nums[i] = val
			i++
		}
	}

}

1st iter: i=0, val=7, nums{7}, i++
2nd iter: i=1, val=7, nums{7,7}, i++
3rd iter: i=2, val=7, nums{7,7,7}
4th iter: i=2, val=6, nums{7,7,6,6,5,5,5,4,4,3,2,1,1,1}, i = 3
5th iter: i=3, val=5, nums{7,7,}


func main() {
	tests := [][]int{
		{7, 7, 7, 6, 5, 5, 5, 4, 4, 3, 2, 1, 1, 1},
		{1, 1, 1, 2, 2, 3},
	}
	for _, t := range tests {
		orig := make([]int, len(t))
		copy(orig, t)
		result := removeDuplicates(t)
		fmt.Printf("orig=%v, k=%d, result=%v\n", orig, k, t[:k])
	}
}
