package main

// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number.
//Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 < numbers.length.

// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
// Your solution must use only constant extra space.

func main() {

}

func twoSum(numbers []int, target int) []int {

	sumNumbers := []int{}

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				sumNumbers = append(sumNumbers, i+1, j+1)
				return sumNumbers
			}
		}
	}

	return sumNumbers
}

//sudocode
//1. find 2 numbers in the array that add to the target number (2 loops)
//2. return the index of these 2 numbers plus 1
//3. -1, 0 case
//array cannot have repeat numbers, they always inrease. Array will only hvae one solution
