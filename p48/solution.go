/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/7
*/
package main

import "fmt"

func printMatrix(matrix [][]int) {
	for _, row := range matrix  {
		for _, n := range row {
			fmt.Printf("%v ", n)
		}
		fmt.Println()
	}
	fmt.Println()
}

func rotate(matrix [][]int)  {
	n := len(matrix)

	for i := 0; i  < n / 2; i++ {
		for j :=i; j < n - 1 -i; j++ {
			top := matrix[i][j]
			matrix[i][j] = matrix[n - 1 - j][i]
			matrix[n - 1 - j][i] = matrix[n - 1 - i][n - 1 - j]
			matrix[n - 1 - i][n - 1 - j] = matrix[j][n - 1 - i]
			matrix[j][n - 1 - i] = top
		}
	}
	printMatrix(matrix)

}


func main() {
	matrix := [][]int {
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
}
