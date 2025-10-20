package main

import "fmt"

func hIndex(citations []int) int {
	n := len(citations)
	total := 0
	bucket := make([]int, len(citations)+1)
	for _, c := range citations {
		if c >= n {
			bucket[n]++
		} else {
			bucket[c]++
		}
	}
	for i := n; i >= 0; i-- {
		total += bucket[i]
		if total >= i {
			return i
		}
	}
	return 0
}

func main() {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Printf("\t h-index : %d\n", hIndex(citations))

}
