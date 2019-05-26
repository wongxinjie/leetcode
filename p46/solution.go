/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/5/12
*/
package main

import "fmt"

var result [][]int


func permute(nums []int) [][]int {
	result = make([][]int, 0)
	rec_permute(nums, 0)
	return result
}

func rec_permute(nums []int, n int) {
	size := len(nums)
	if n == size - 1 {
		t := make([]int, len(nums))
		copy(t, nums)
		result = append(result, t)
		return
	}

	for i := n; i < size; i++ {
		nums[i], nums[n] = nums[n], nums[i]
		rec_permute(nums, n + 1)
		nums[i], nums[n] = nums[n], nums[i]
	}
}


func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}