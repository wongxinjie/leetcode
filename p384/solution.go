/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/7/31
*/
package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	original []int
	shuffled []int
}


func Constructor(nums []int) Solution {
	shuffled := make([]int, 0)
	for _, n := range nums {
		shuffled = append(shuffled, n)
	}

	s := Solution{
		original: nums,
		shuffled: shuffled,
	}
	return s
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.original
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	size := len(this.original)
	for i := size -1; i > 0; i-- {
		pos := rand.Int() % (i + 1)
		fmt.Println(pos)
		this.shuffled[pos], this.shuffled[i] = this.shuffled[i], this.shuffled[pos]
	}
	return this.shuffled
}

func main() {
	o := Constructor([]int{1, 2, 3, 4})
	fmt.Println(o.Shuffle())
	fmt.Println(o.Reset())

}