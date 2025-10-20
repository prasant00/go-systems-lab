package main

import "fmt"

func jump(nums []int) int {

	reach := 0
	n := len(nums)
	minJumps := n - 1

	for i := 0; i < len(nums); i++ {
		jumps := nums[i]
		k := i + 1
		for j := 0; j < jumps; j++ {
			if reach < k+nums[k] {
				reach = k + nums[k]
			}
			if reach >= n-1 {
				if k < minJumps {
					minJumps = jumps
				}
			}
			k++
		}
	}
	return minJumps
}

func main() {
	nums := []int{2, 1, 2, 1, 4}
	fmt.Printf("\t min jumps required: %d\n", jump(nums))
}
