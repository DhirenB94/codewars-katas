package main

import (
	"fmt"
	"math"
)

func main() {
	process([]int{4,3,9,7,2,1})
}


func process(input []int) []int{
	output := []int{}
	for _, number := range input {
		checkSquare := squareOrSquareRoot(number)
		output = append(output, checkSquare)
	}
	fmt.Println(output)
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



