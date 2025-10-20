package main

import "fmt"

// {7, 1, 5, 3, 6, 8}
func maxProfit(prices []int) int {
	total := len(prices)
	fmt.Printf("\t total = %d\n", total)
	min, max := prices[0], prices[0]

	for i := 1; i <= total-1; i++ {
		fmt.Printf("\t price = %d\n", prices[i])
		if prices[i] < min {
			min = prices[i]
			max = prices[i+1]
		} else if prices[i] > max {
			max = prices[i]
		}
	}
	fmt.Printf("\t min = %d, max = %d \n", min, max)
	return max - min
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 8}
	profit := maxProfit(prices)
	fmt.Printf("\t maxProfit = %d]n", profit)
}
