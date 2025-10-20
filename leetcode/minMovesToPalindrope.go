package main

import (
	"fmt"
	"strings"
)

// "mamad"
// maamd
// maadm
// madam

func minMovesToMakePalindrome(s string) int {
	runes := []rune(s)

	moves := 0

	for i, j := 0, len(runes)-1; i < j; i++ {
		k := j
		for k > i {
			if runes[i] == runes[k] {
				for k < j {
					runes[k], runes[k+1] = runes[k+1], runes[k]
					moves++
					k++
				}
				j--
				break
			}
			k--
		}
		if k == i {
			moves += len(runes)/2 - i
		}
	}
	return moves
}

// Driver code
func main() {
	strs := []string{"mamad"}

	for i, s := range strs {
		fmt.Printf("%d.\ts: %s\n", i+1, s)
		fmt.Printf("\tMoves: %d\n", minMovesToMakePalindrome(s))
		fmt.Println(strings.Repeat("-", 100))
	}
}
