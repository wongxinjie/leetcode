/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/4/26
 */
package main

import "fmt"

func reverseBits(num uint32) uint32 {
	bits := make([]uint32, 0)
	for i := 0; i < 32; i++ {
		bits = append(bits, (num>>uint(i))&1)
	}

	result := uint32(0)
	for i := 32; i > 0; i-- {
		result += bits[i-1] << uint32(32-i)
	}
	return result
}

func main() {
	n := reverseBits(4294967293)
	fmt.Println(n)
	fmt.Println()
}
