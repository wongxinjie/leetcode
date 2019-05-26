/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/5/4
*/
package main

func rotate(nums []int, k int)  {
	size := len(nums)
	rotate := make([]int, 0, size)

	for n := 0; n < size; n++  {
		rotate = append(rotate, 0)
	}

	for n := 0; n < size; n++ {
		idx := (n + k) % size
		rotate[idx] = nums[n]
	}

	for n := 0; n < size; n++ {
		nums[n] = rotate[n]
	}
}

func main()  {
	rotate([]int{-1, -100, 3, 99}, 3)
}
