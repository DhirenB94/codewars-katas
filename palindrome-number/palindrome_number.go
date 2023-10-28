package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(-1213))
}

func isPalindrome(x int) bool {
	number := strconv.Itoa(x)

	for i := 0; i < len(number)/2; i++ {
		for j := len(number) - 1; j >= len(number)/2; j-- {
			if number[i] != number[j] {
				return false
			}
			i++
		}
	}
	return true
}
