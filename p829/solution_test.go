/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/4/23
 */
package main

import "testing"

func Test_consecutiveNumbersSum(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2", args{N: 5 }, 2},
		{"2", args{N: 9}, 3},
		{"2", args{N: 15}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := consecutiveNumbersSum(tt.args.N); got != tt.want {
				t.Errorf("consecutiveNumbersSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
