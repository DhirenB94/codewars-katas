package main

import "fmt"

func main () {
	rotate(2, []int{11,22,33,44,55,66} )
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
	fmt.Println(input)
	return input
}
