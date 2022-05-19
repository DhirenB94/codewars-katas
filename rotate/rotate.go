package main

import "fmt"

//Write a function that rotates a slice by k elements.
//For example [1,2,3,4,5,6] rotated by 2 becomes [3,4,5,6,1,2].

//Bonus: Try solving this without creating a copy of the slice.

func main () {
	fmt.Println(rotate(6, []int{1,2,3,4,5,6,7, 8 ,9, 10} ))
}

func rotate(k int, input []int) []int  {
	trimmedInput := input [:k]
	input = input[k:]

	for _, trim := range trimmedInput {
		input = append(input, trim)
	}
	return input
}
