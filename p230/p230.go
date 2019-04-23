package main

import (
	"fmt"
)

// * Definition for singly-linked list.
type ListNode struct {
	Val int
 	Next *ListNode
}


func getListLen(head *ListNode) int {
	length := 0
	cursor := head
	for cursor != nil  {
		cursor = cursor.Next
		length += 1
	}
	return length
}

func isPalindrome(head *ListNode) bool {
	listSize := getListLen(head)

	if listSize == 1 || listSize == 0  {
		return true
	}

	half := listSize / 2
	values := make([]int, half)
	cursor := head
	for i := 0; i < half; i++ {
		values[i] = cursor.Val
		cursor = cursor.Next
	}

	//奇数，跳过一个节点
	if listSize % 2 != 0 {
		cursor = cursor.Next
	}

	fmt.Println(values)
	n := half - 1
	for cursor != nil {
		if cursor.Val != values[n] {
			fmt.Println(cursor.Val)
			return false
		}
		n -= 1
		cursor = cursor.Next
	}

	return true
}

func main()  {
	l1 := ListNode{1, nil}
	l1.Next = &ListNode{0, nil}
	l1.Next.Next = &ListNode{1, nil}
	//l1.Next.Next.Next = &ListNode{1, nil}
	fmt.Println(isPalindrome(&l1))
}