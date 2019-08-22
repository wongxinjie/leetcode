/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/21
*/
package main

import "fmt"

func numIslands(grid [][]byte) int {
	count := 0
	if len(grid) == 0 {
		return count
	}

	row, col := len(grid), len(grid[0])
	visited := make([][]bool, 0)
	for i := 0; i < row; i++ {
		visited = append(visited, make([]bool, col))
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				count += 1
				dfs(grid, &visited, i, j)
			}
		}
	}


	return count
}


func dfs(grid [][]byte, visited *[][]bool, x, y int) {
	row, col := len(grid), len(grid[0])

	if  x < 0 || y < 0 || x >= row || y >= col {
		return
	}
	if grid[x][y] == '0' || (*visited)[x][y]{
		return
	}

	if grid[x][y] == '1' {
		(*visited)[x][y] = true
	}

	dfs(grid, visited, x - 1, y)
	dfs(grid, visited, x + 1, y)
	dfs(grid, visited, x, y - 1)
	dfs(grid, visited, x, y + 1)
}


func main() {
	/*
	 */
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	c := numIslands(grid)
	fmt.Printf("%v\n", c)
}
