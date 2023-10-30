package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(1212))
}

func isPalindrome(x int) bool {
	number := strconv.Itoa(x)

	j := len(number) -1
	for i := 0; i < len(number)/2; i++ {
			if number[i] != number[j] {
				return false
			}
			j--
		}
	return true
}
