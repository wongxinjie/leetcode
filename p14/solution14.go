package main

import (
	"fmt"
)

func commonPrefix(a, b string) string {
	if a == "" || b == "" {
		return ""
	}

	aLen, bLen := len(a), len(b)
	idx := 0
	for idx < aLen && idx < bLen {
		if a[idx] == b[idx] {
			idx += 1
		} else {
			break
		}
	}
	fmt.Println(idx)
	if idx > 0 {
		return a[:idx]
	} else {
		return ""
	}
}

func longestCommonPrefix(strs []string) string{
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0]
	strs = strs[1:]
	for _, word := range strs {
		fmt.Println(word, prefix)
		prefix = commonPrefix(prefix, word)
	}
	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}