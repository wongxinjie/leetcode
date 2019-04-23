/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/4/23
*/
package main

import (
	"fmt"
)

func consecutiveNumbersSum(N int) int {
	count := 0
	counter := N * 2
	for n := 1; n * n < counter; n++ {
		if counter % n == 0 && ((n + N /n) & 1) == 1 {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Println(consecutiveNumbersSum(5))
	fmt.Println(consecutiveNumbersSum(9))
	fmt.Println(consecutiveNumbersSum(15))
}