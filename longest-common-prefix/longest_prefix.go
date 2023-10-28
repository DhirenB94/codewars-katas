package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"aacc", "aa", "aa", "aa", "aaca"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 { // handle only 1 element
		return strs[0]
	}

	// sort them first, the most different ones will be in first and last
	sort.Strings(strs)

	// compare first and last
	lastPos := strs[len(strs)-1]
	first := strs[0]

	for i := 0; i < len(first); i++ {
		if first[i] != lastPos[i] {
			return first[:i]
		}
	}
	return first
}
