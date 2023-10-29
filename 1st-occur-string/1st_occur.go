package main

import "fmt"


func main () {
	fmt.Println(strStr("leetcode", "leeto"))
}


func strStr(haystack string, needle string) int {

	needleLength := len(needle) //5
	haystackLenghth := len(haystack) //8

	maxIterations := haystackLenghth - needleLength //3

	if maxIterations == 0 && haystack != needle {
		return -1
	}

	for i := 0; i <= maxIterations; i++ {
		if haystack[i:needleLength] == needle {
			return i
		} else {
			needleLength ++
		}
	}

	return -1
}