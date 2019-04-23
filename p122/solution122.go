package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	size := len(prices)
	for i := 0; i < size-1; i++ {
		if prices[i+1] > prices[i] {
			maxProfit += prices[i+1] - prices[i]
		}
	}
	return maxProfit
}

func main() {
	testCase1 := []int{7, 1, 5, 3, 6, 4}
	testCase2 := []int{1, 2, 3, 4, 5}
	testCase3 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(testCase1))
	fmt.Println(maxProfit(testCase2))
	fmt.Println(maxProfit(testCase3))
}
