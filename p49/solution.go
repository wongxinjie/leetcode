/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/13
*/
package main

import (
	"sort"
)

type sortedRunes []rune

func(s sortedRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortedRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortedRunes) Len() int{
	return len(s)
}

func SortString(s string) string {
	runes := []rune(s)
	sort.Sort(sortedRunes(runes))
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	stringMap := make(map[string][]string)

	for _, s := range strs {
		key := SortString(s)
		if _, ok := stringMap[key]; ok {
			stringMap[key] = append(stringMap[key], s)
		} else {
			values := make([]string, 0)
			values = append(values, s)
			stringMap[key] = values
		}
	}

	result := make([][]string, 0)
	for _, v := range stringMap {
		result = append(result, v)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(strs)
}
