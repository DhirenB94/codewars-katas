package main

import "fmt"

func main() {
	_ = trailingZeroes(25)

}

func trailingZeroes(n int) int {
	count := 0
	for n >= 5 {
		n /= 5
		count += n

	}
	fmt.Println(count)
	return count
}
