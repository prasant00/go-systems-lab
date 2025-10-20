package main

import (
	"fmt"
	"sort"
)

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	count := 0
	low, high := 0, len(nums)-1

	for low < high {
		if nums[low]+nums[high] < target {
			count += high - low
			low++
		} else {
			high--
		}
	}

	return count
}

func main() {

	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{10, 1, 6, 2, 3, 8}, 9},
	}

	for i, testCase := range testCases {
		fmt.Printf("%d\tnums: %v\n", i+1, testCase.nums)
		fmt.Printf("%d\ttarget: %d\n", i+1, testCase.target)
		result := countPairs(testCase.nums, testCase.target)
		fmt.Printf("\n\tno of pairs: %d\n", result)
	}
}
