package main

import (
	"fmt"
	"math"
)

//Write a method, that will get an integer array as parameter and will process every number from this array.
//Return a new array with processing every number of the input-array like this:
//If the number has an integer process root, take this, otherwise process the number.
//The input array will always contain only positive numbers, and will never be empty or null

//Example
//[4,3,9,7,2,1] -> [2,9,3,49,4,1]
func main() {
	fmt.Println(process([]int{4,3,9,7,2,1}))

}


func process(input []int) []int{
	output := []int{}
	for _, number := range input {
		checkSquare := squareOrSquareRoot(number)
		output = append(output, checkSquare)
	}
	return output
}

func squareOrSquareRoot(num int) int {
	float :=  float64(num)
	intRoot := int(math.Sqrt(float))

	if intRoot * intRoot == num {
		return intRoot
	}
	return num * num
}



