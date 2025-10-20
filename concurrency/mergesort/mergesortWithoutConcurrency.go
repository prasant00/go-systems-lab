package main

import (
	"fmt"
)

// merge merges two sorted slices into a single sorted slice
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	if i < len(left) {
		result = append(result, left[i:]...)
	}
	if j < len(right) {
		result = append(result, right[j:]...)
	}

	return result
}

// []int{38, 27, 43, 3, 9, 82, 10}
// mergeSort recursively divides and sorts the slice
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original:", arr)

	sorted := mergeSort(arr)
	fmt.Println("Sorted:  ", sorted)
}
