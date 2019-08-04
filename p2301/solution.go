/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/6/2
*/
/*
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Note:
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

Example 1:

Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
Output: 1
Example 2:

Input: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
Output: 3
Follow up:
What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently?
How would you optimize the kthSmallest routine?
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Result struct {
	counter int
	val int
	nth int
}

func walk(root *TreeNode, r *Result){
	if root != nil {
		walk(root.Left, r)

		r.counter += 1
		if r.counter == r.nth {
			r.val = root.Val
		}
		walk(root.Right, r)
	}

}

func kthSmallest(root *TreeNode, k int) int {
	result := Result{0, 0, k}
	walk(root, &result)
	return result.val
}


func makeTree() *TreeNode{
	root := TreeNode{3 , nil, nil }
	root.Left = &TreeNode{1, nil, nil}
	root.Right = &TreeNode{4, nil, nil}
	root.Left.Right = &TreeNode{3, nil, nil}
	return &root
}

func main() {
	tree := makeTree()
	fmt.Println(kthSmallest(tree, 1))
}