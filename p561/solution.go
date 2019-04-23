package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	total := 0
	for i := 0; i < len(nums); i+= 2 {
		total += nums[i]
	}
	return total
}

func main() {
	a := []int{1, 3, 2, 4}
	fmt.Println(arrayPairSum(a))
}