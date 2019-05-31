/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/6/1
*/
package main

import "fmt"

func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	n := len(nums)
	bitLimit := 1 << uint(n)
	fmt.Println(bitLimit)

	for i := 0; i < bitLimit; i++ {
		subset := make([]int, 0)

		for j := 0; j < n ; j++ {
			if i & ( 1 << uint(j)) != 0 {
				subset = append(subset, nums[j])
			}
		}

		result = append(result, subset)
	}
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
