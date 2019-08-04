/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/6/20
*/
package main

import "fmt"

func calSumMap(A []int, B []int) map[int]int {
	sumMap := make(map[int]int, 0)
	for _, x := range A {
		for _, y := range B {
			s := x + y
			if count, ok := sumMap[s]; ok {
				sumMap[s] = count + 1
			} else {
				sumMap[s] = 1
			}
		}
	}
	return sumMap
}


func fourSumCount(A []int, B []int, C []int, D []int) int {
	totalCount :=  0

	abSumMap := calSumMap(A, B)
	cdSumMap := calSumMap(C, D)

	for k, c := range abSumMap {
		if count, ok := cdSumMap[-k]; ok {
			totalCount += c * count
		}
	}
	return totalCount
}


func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}