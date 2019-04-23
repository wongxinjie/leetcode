package main

import "fmt"

func trailingZeroes(n int) int {
	count := 0
	for (n / 5) > 0 {
		n  = n / 5
		count += n
	}
	return count
}


func main() {
	n := 100
	fmt.Println(trailingZeroes(n))
}