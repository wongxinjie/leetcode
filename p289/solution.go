/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/14
*/
package main

import "fmt"

func updateBoard(i, j, row, col int, board [][]int){
	liveCount := 0

	state := board[i][j]

	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, 1}, {1, -1},
	}

	for _, d := range directions {
		rowOffset, colOffset := d[0], d[1]
		x := i + rowOffset
		y := j + colOffset
		if x >= 0 && x < row && y >= 0 && y < col {
			if board[x][y] == 1 || board[x][y] == -1 {
				liveCount += 1
			}
		}
	}

	if state == 1 {
		if liveCount < 2 {
			state = -1
		}
		if liveCount > 3 {
			state = -1
		}
	}
	if state == 0 && liveCount == 3 {
		state = 2
	}
	board[i][j] = state
}

func gameOfLife(board [][]int)  {
	row := len(board)
	if row == 0 {
		return
	}
	col := len(board[0])

	for i := 0; i < row; i++  {
		for j := 0; j < col; j++ {
			updateBoard(i, j, row, col, board)
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			v := board[i][j]
			if v == -1 {
				board[i][j] = 0
			}
			if v == 2 {
				board[i][j] = 1
			}
		}
	}
}

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	fmt.Printf("%v\n", board)
	gameOfLife(board)
	fmt.Printf("%v\n", board)
}