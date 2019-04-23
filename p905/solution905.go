package main

import (
	"fmt"
	"sort"
)

func sortArrayByParity(A []int) []int {
	evenArray := make([]int, 0)
	oddArray := make([]int, 0)

	for _, n := range A {
		if n%2 == 0 {
			evenArray = append(evenArray, n)
		} else {
			oddArray = append(oddArray, n)
		}
	}

	sort.Ints(evenArray)
	sort.Ints(oddArray)
	evenArray = append(evenArray, oddArray...)
	return evenArray
}

func main() {
	a := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(a))
}
