/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/5/26
*/
package main

import (
	"fmt"
	"sort"
)

func (k KV) String() string {
	return fmt.Sprintf("key: %v, value: %v", k.key, k.value)
}

type KV struct {
	key int
	value int
}

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)

	for _, n := range nums {
		_, ok := counter[n]
		if ok {
			counter[n] += 1
		} else {
			counter[n] = 1
		}
	}

	frequent := make([]KV, 0)
	for k, v := range counter {
		frequent = append(frequent, KV{k, v})
	}

	sort.Slice(frequent, func(i, j int) bool {
		return frequent[i].value > frequent[j].value
	})

	result := make([]int, 0)
	for _, k := range frequent[:k] {
		result = append(result, k.key)
	}
	return result
}

func main() {
	r := topKFrequent([]int{1}, 1)
	fmt.Println(r)
}
