/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/4/25
*/
package main

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	n := 1
	nSqrt := 1
	for  nSqrt <= x {
		n += 1
		nSqrt = n * n
	}
	return  n - 1
}