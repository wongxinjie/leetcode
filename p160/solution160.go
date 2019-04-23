/**
* @File: solution160.go
* @Author: wongxinjie
* @Date: 2019/4/21
*/
package main

import "fmt"

//* Definition for singly-linked list.
type ListNode struct {
      Val int
      Next *ListNode
}


func getListLength(head *ListNode) int {
	length := 0
	for cursor := head; cursor != nil; cursor = cursor.Next {
		length += 1
	}
	return length
}


func getIntersectionNode(headA, headB *ListNode) *ListNode {
	listALen := getListLength(headA)
	listBLen := getListLength(headB)

	if listALen == 0 || listBLen == 0 {
		return nil
	}

	if listALen >  listBLen {
		for n := 0; n < (listALen - listBLen); n++ {
			headA = headA.Next
		}
	} else {
		for n := 0; n < (listBLen - listALen); n++ {
			headB = headB.Next
		}
	}

	var common *ListNode
	leadZero := false
	for headA != nil && headB != nil {
		if headA.Val == headB.Val && !leadZero{
			if common == nil {
				common = headA
			}
		} else {
			common = nil
		}
		if headA.Val == 0 || headB.Val == 0 {
			leadZero = true
		} else  {
			leadZero = false
		}
		headA = headA.Next
		headB = headB.Next
	}
	return common
}

func makeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	cursor := head
	for _, n := range nums[1:] {
		node := &ListNode{n, nil}
		cursor.Next = node
		cursor = node
	}
	return head
}

func main() {
	//common := makeList([]int{2, 4})
	headA := makeList([]int{2, 4, 6})
	headB := makeList([]int{1, 5})
	//headA.Next.Next.Next = common
	//headB.Next = common

	result := getIntersectionNode(headA, headB)
	fmt.Println(result)
}