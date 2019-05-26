/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/5/21
*/
package main

import "fmt"

func productExceptSelf(nums []int) []int {
	size := len(nums)
	left := make([]int, 0)
	right := make([]int,  0)
	for i := 0; i < size; i++ {
		left = append(left, 1)
		right = append(right, 1)
	}

	for i := 0; i < size - 1; i++ {
		left[i+1] = left[i] * nums[i]
	}
	for i := size - 1; i > 0; i-- {
		right[i - 1] = right[i] * nums[i]
	}

	result := make([]int, 0)
	for i := 0; i < size; i++ {
		result = append(result, left[i] * right[i])
	}

	return result
}


func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
