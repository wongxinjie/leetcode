package main

import "fmt"

func transpose(A [][]int) [][]int {
	row, col := len(A), len(A[0])
	result := make([][]int, 0)

	for x := 0; x < col; x ++ {
		r := make([]int, 0)
		for y := 0; y < row; y++ {
			r = append(r, A[y][x])
		}
		result = append(result, r)
	}
	return result
}

func main() {
	a := [][]int {
		{1, 2, 3}, {4, 5, 6},
	}
	fmt.Println(transpose(a))
}
