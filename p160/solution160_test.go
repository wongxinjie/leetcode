/**
* @File: solution160.go
* @Author: wongxinjie
* @Date: 2019/4/21
 */
package main

import (
	"reflect"
	"testing"
)

func makeList(nums []int) *ListNode{
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

func Test_getIntersectionNode(t *testing.T) {
	common := makeList([]int{8, 4, 5})
	headA := makeList([]int{4, 1})
	headB := makeList([]int{5, 0, 1})
	headA.Next.Next = common
	headB.Next.Next.Next = common
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test1", args{headA, headB}, common},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
