package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	size := len(nums)
	maxSum := nums[0]
	sumArray := make([]int, 0)
	sumArray = append(sumArray, nums[0])

	for i := 1; i < size; i++ {
		if sumArray[i-1] > 0 {
			sumArray = append(sumArray, nums[i]+sumArray[i-1])
		} else {
			sumArray = append(sumArray, nums[i])
		}
		if sumArray[i] > maxSum {
			maxSum = sumArray[i]
		}
	}

	return maxSum
}

func main() {
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{-1}
	fmt.Printf("%v\n", maxSubArray(nums))
}
