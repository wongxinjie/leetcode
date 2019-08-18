/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/17
*/
package main

import "fmt"

func sortColors(nums []int)  {
	counter := [3]int{0, 0, 0}

	for _, n := range nums {
		counter[n] += 1
	}

	begin, end := 0, counter[0]
	for begin < end {
		nums[begin] = 0
		begin += 1
	}

	end += counter[1]
	for begin < end {
		nums[begin] = 1
		begin += 1
	}

	end += counter[2]
	for begin < end {
		nums[begin] = 2
		begin += 1
	}
}

func main() {
	nums := []int{2,0,2,1,1,0, 1}
	sortColors(nums)
	fmt.Printf("%v\n", nums)
}