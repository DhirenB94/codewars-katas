package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestConsecutive([]int{0, 1}))
}

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)

	// Create a set of unique numbers for efficient lookups
	for _, num := range nums {
		numSet[num] = true
	}
	fmt.Println(numSet)

	longestStreak := 0

	for num := range numSet {
		fmt.Println("range", num)
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			// Check for consecutive elements greater than the current number
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			longestStreak = max(longestStreak, currentStreak)
		}
	}

	return longestStreak
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
