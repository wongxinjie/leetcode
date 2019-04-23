package main

import "fmt"

func flipArrray(a []int) []int {
	s, e := 0, len(a) - 1
	for {
		a[s], a[e] = a[e], a[s]
		s++
		e--
		if s >= e {
			break
		}
	}
	return a
}

func reverseArray(a []int) []int {
	for i, n := range a {
		if n == 0 {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
	return a
}

func filpAndInvertImage(A [][]int) [][]int {
	for i, row := range A {
		row = flipArrray(row)
		row = reverseArray(row)
		A[i] = row
	}
	return A
}

func main() {
	image := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(filpAndInvertImage(image))
}

