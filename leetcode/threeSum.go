package main

import (
	"fmt"
	"sort"
	"strings"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}

	n := len(nums)

	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i == 0 || nums[i] != nums[i-1] {
			low, high := i+1, n-1

			for low < high {
				sum := nums[i] + nums[low] + nums[high]

				if sum < 0 {
					low++
				} else if sum > 0 {
					high--
				} else {
					result = append(result, []int{nums[i], nums[low], nums[high]})

					low++
					high--
					for low < high && nums[low] == nums[low-1] {
						low++
					}
					for low < high && nums[high] == nums[high+1] {
						high--
					}
				}
			}
		}
	}

	return result
}

// Driver code
func main() {
	numsArrs := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{1, 2, 3, 4, 5},
		{0, 0, 0, 0},
		{-4, -1, -1, 0, 1, 2, 2},
		{-10, -7, -3, -1, 0, 3, 7, 10},
		{-3, -5, -7, -9},
	}

	for i, nums := range numsArrs {
		fmt.Printf("%d.\tnums: [", i+1)
		for j, num := range nums {
			fmt.Printf("%d", num)
			if j < len(nums)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")

		triplets := threeSum(nums)
		fmt.Print("\n\tTriplets: [")
		for j, triplet := range triplets {
			fmt.Printf("[%d, %d, %d]", triplet[0], triplet[1], triplet[2])
			if j < len(triplets)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
		fmt.Println(strings.Repeat("-", 100))
	}
}
