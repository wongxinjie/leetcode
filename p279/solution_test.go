/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/17
 */
package p279

import "testing"

func Test_numSquares(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-1", args{12}, 3},
		{"test-2", args{13}, 2},
		{"test-3", args{16}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.args.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
