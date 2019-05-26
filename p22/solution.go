/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/5/26
*/
package main

import "fmt"

type Result []string

func (r *Result) generateDfs(left, right int, str string) {
	if left < right {
		return
	}
	if left == 0 && right == 0 {
		*r = append(*r, str)
	} else {
		if left > 0 {
			r.generateDfs(left - 1, right, str + ")")
		}
		if right > 0 {
			r.generateDfs(left, right - 1, str + "(")
		}
	}

}

func generateParenthesis(n int) []string {
	var result Result
	result.generateDfs(n, n, "")
	return []string(result)
}

func main() {
	r := generateParenthesis(3)
	fmt.Println(r)
}
