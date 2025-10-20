package main

import "fmt"

func merge(left, right []int) []int {
	final := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}
	if i < len(left) {
		final = append(final, left[i:]...)
		i++
	}
	if j < len(right) {
		final = append(final, right[j:]...)
		j++
	}
	return final
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	done := make(chan bool)
	var left []int
	go func() {
		left = mergeSort(nums[:mid])
		done <- true
	}()
	right := mergeSort(nums[mid:])
	<-done
	return merge(left, right)
}
func main() {
	nums := []int{38, 27, 43, 3, 9, 82, 10}
	sortedNums := mergeSort(nums)
	fmt.Printf("\t sorted numbers %+v", sortedNums)
}
