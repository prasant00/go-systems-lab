package main

import (
	"fmt"
)

func lengthOfLongestSubstringASCII(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// last index seen for each byte; initialize to -1
	var last [256]int
	for i := range last {
		last[i] = -1
	}

	maxLen := 0
	left := 0

	for right := 0; right < n; right++ {
		ch := s[right]
		fmt.Printf("\n ch = %d\n", ch)
		fmt.Printf("1->last[ch] = %d \n", last[ch])
		fmt.Printf("1->left = %d \n", left)
		if last[ch] >= left {
			// duplicate inside window, move left one past previous index
			left = last[ch] + 1
			fmt.Printf("2->left = %d \n", left)
		}
		fmt.Printf("3->left = %d \n", left)
		last[ch] = right
		fmt.Printf("2->last[ch] = %d \n", last[ch])
		if cur := right - left + 1; cur > maxLen {
			maxLen = cur
		}
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstringASCII("abcabcbb")) // 3
}
