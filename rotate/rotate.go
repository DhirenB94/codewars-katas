package main

import "fmt"

func main() {
	// rotateLeft(2, []int{11, 22, 33, 44, 55, 66})
	rotateRight([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

// 1. Write a function that rotates a slice to the left by k elements. For example [1,2,3,4,5,6] rotated by 2 becomes [3,4,5,6,1,2].
func rotateLeft(k int, input []int) []int {

	for j := 0; j < k; j++ {

		var i = 0
		var temp = input[0] //11

		for ; i < len(input)-1; i++ {
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

//2. Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

//without using a new slice
func rotateRight(input []int, k int) []int {

	for j := 0; j <k; j++{

	var i = len(input) - 1
	var temp = input[len(input)-1] // 7

	for ; i > 0; i-- {
		input[i] = input[i-1]
		//move everything 1 to the right
		//[1 1 2 3 4 5 6]
	}
	input[i] = temp
	//set the firts index as last position before the movinng to the right
	//[22 33 44 55 66 11]
	}
	fmt.Println(input)
	return input
}

//using a new slice
// func rotate(nums []int, k int) []int {
// 	numsLength := len(nums) //6
// 	lastPos := numsLength -1
// 	newArray := make([]int, numsLength)

// 	for index, num := range nums {
// 		if index+k >= numsLength {
// 			pos := index + k
// 			newPos := pos - lastPos -1
// 			newArray[newPos] = num
// 		} else {
// 			newArray[index + k] = num
// 		}
// 	}
// 	return newArray
// }
