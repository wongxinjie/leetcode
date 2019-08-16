/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/16
*/
package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	index map[int]int
	items []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	index := make(map[int]int)
	items := make([]int, 0)
	return RandomizedSet{
		index: index,
		items: items,
	}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.index[val]; ok {
		return false
	}

	this.items = append(this.items, val)
	this.index[val] = len(this.items) - 1
	return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.index[val]
	if !ok {
		return false
	}
	delete(this.index, val)

	size := len(this.items)
	// 如果不是最后一个，则需要移动
	if idx != size - 1 {
		v := this.items[size - 1]
		this.items[idx] = v
		this.index[v] = idx
	}
	this.items = this.items[:size - 1]

	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	size := len(this.items)
	if size > 0 {
		idx := rand.Intn(size)
		return this.items[idx]
	}
	return 0
}

func main()  {
	set := Constructor()
	p := set.Insert(1)
	fmt.Printf("%v\n", p)
	p = set.Insert(2)
	fmt.Printf("%v\n", p)
	p = set.Remove(2)
	p = set.Remove(1)
	v := set.GetRandom()
	fmt.Printf("%v\n", v)

}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
