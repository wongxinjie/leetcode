/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/11
*/
package p62


func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	matrix := make([][]int, 0)
	for i := 0; i < m; i++ {
		row := make([]int, n)
		matrix = append(matrix, row)
	}
	matrix[0][0] = 1

	for i := 1; i < m; i++ {
		matrix[i][0] = 1
	}
	for i := 1; i < n; i++ {
		matrix[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}
	return matrix[m-1][n-1]
}