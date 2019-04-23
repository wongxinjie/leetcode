package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	row, col := len(nums), len(nums[0])

	if r * c != row * col {
		return nums
	}

	result := make([][]int, 0)
	x, y := 0, 0
	each := make([]int, 0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			each = append(each, nums[i][j])
			y += 1
			if y == c {
				x += 1
				y = 0
				result = append(result, each)
				each = make([]int, 0)
			}
		}
	}
	return result
}


func main() {
	nums := [][]int {
		{1, 2},
		{3, 4},
	}
	fmt.Println(matrixReshape(nums, 2 , 4))
	// fmt.Println(matrixReshape(nums, 2, 4))
}