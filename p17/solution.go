/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/19
*/
package main

import (
	"strings"
)

type Tracking struct {
	result []string
	digitMap map[string][]string
}

func backTracking(combination string, nextDigits []string, track *Tracking) {
	if len(nextDigits) == 0 {
		track.result = append(track.result, combination)
	} else {
		for _, letter := range track.digitMap[nextDigits[0]] {
			backTracking(combination + letter, nextDigits[1:], track)
		}
	}
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}

	digitCharMap := map[string][]string {
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	track := &Tracking{
		result: result,
		digitMap: digitCharMap,
	}

	combination := ""
	digitArray := strings.Split(digits, "")
	backTracking(combination, digitArray, track)
	return track.result
}

//func main() {
//	s := ""
//	r := letterCombinations(s)
//	fmt.Printf("%v\n", r)
//}