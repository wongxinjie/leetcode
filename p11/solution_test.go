/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/16
 */
package p11

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"testA", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"testA", args{[]int{3, 2, 1, 3}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
