package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	row, col := len(matrix), len(matrix[0])

	for i := 0; i < row - 1; i++ {
		for j := 0; j < col - 1; j++ {
			if(matrix[i][j] != matrix[i+1][j+1]) {
				return false
			}
		}
	}
	return true
}


func main() {
	matrix := [][]int {
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}
	fmt.Println(isToeplitzMatrix(matrix))
}