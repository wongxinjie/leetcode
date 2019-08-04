/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/4
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
}


func makeList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{
		Val: values[0],
		Next: nil,
	}

	current := head
	values = values[1:]
	for _, v := range values {
		n := &ListNode{
			Val: v,
			Next: nil,
		}
		current.Next = n
		current = n
	}
	return head
}

func (head *ListNode) String() string{
	values := make([]string, 0)
	current := head
	for current != nil {
		values = append(values, strconv.Itoa(current.Val))
		current = current.Next
	}
	return strings.Join(values,"->")
}


func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = even.Next
		even.Next = odd.Next
		even = odd.Next
	}
	odd.Next = evenHead
	return head
}

func main()  {
	head := makeList([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Printf("%v\n", head)
	node := oddEvenList(head)
	fmt.Printf("%v\n", node)
}