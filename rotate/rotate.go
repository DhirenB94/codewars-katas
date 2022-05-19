package main

import "fmt"

//Write a function that rotates a slice by k elements.
//For example [1,2,3,4,5,6] rotated by 2 becomes [3,4,5,6,1,2].

//Bonus: Try solving this without creating a copy of the slice.

func main () {
	fmt.Println(rotate(2, []int{11,22,33,44,55,66} ))
}

func rotate(k int, input []int) []int  {

	for j := 0; j <k; j++{

		var i  = 0
		var temp  = input[0] //11

		for ; i < len(input) - 1; i++ {
			input[i] = input[i+1]
			//move everything 1 to the left
		}
		//[22 33 44 55 66 66]

		input[i] = temp
		//set the last index as the 0 index
		//[22 33 44 55 66 11]
	}
	return input
}
