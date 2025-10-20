package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] > target {
			j--
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			break
		}
	}
	return []int{i + 1, j + 1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	indices := twoSum(nums, target)
	fmt.Printf("\t indices: %v\n", indices)
}
