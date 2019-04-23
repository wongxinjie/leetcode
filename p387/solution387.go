package main

import "fmt"

func firstUniqChar(s string) int {
	counter := make(map[rune]int)
	for _, c := range s {
		if v, ok := counter[c]; ok {
			counter[c] = v + 1
		} else {
			counter[c] = 1
		}
	}

	for idx, c := range s {
		v := counter[c]
		if v == 1 {
			return idx
		}
	}
	return -1
}

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}
