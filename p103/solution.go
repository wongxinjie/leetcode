/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/19
*/
package main

import "fmt"

type TreeNode struct {
	Val int
 	Left *TreeNode
 	Right *TreeNode
}

func reverse(nums []int) {
	length := len(nums)
	for i := 0; i < length / 2; i++ {
		nums[i], nums[length - 1 - i] = nums[length - 1 - i], nums[i]
	}
}


func zigzagLevelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	values := make([][]int, 0)

	zigzag := 0
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		level := make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			currentNode := queue[0]
			queue = queue[1:]
			level = append(level, currentNode.Val)

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
		if zigzag % 2 != 0 {
			reverse(level)
		}
		zigzag += 1
		values = append(values, level)
	}
	return values
}


func main() {
	left := &TreeNode{2, nil, nil }
	right := &TreeNode{3, nil, nil }
	left.Left = &TreeNode{4, nil, nil }
	right.Right = &TreeNode{5, nil, nil}
	root := &TreeNode{1, left, right }

	r := zigzagLevelOrder(root)
	fmt.Printf("%v\n", r)
}